package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func checkString(parameters url.Values, name string) (string, error) {
	v, ok := parameters[name]
	if !ok {
		return "", fmt.Errorf("missing '%s' parameter", name)
	}
	if len(v) == 0 {
		return "", fmt.Errorf("missing '%s' parameter", name)
	}
	if len(v) > 1 {
		return "", fmt.Errorf("parameters '%s' confilct: found %d time(s): %v", name, len(v), v)
	}
	return v[0], nil
}

func checkInt(parameters url.Values, name string) (int, error) {
	v, err := checkString(parameters, name)
	if err != nil {
		return 0, err
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, fmt.Errorf("could not convert '%s' parameter: %s", name, err.Error())
	}

	if i < 0 {
		return 0, fmt.Errorf("parameter '%s' could not be negative value: %d", name, i)
	}
	return i, nil

}

func checkFizzBuzz(f *fizzBuzz, w http.ResponseWriter, r *http.Request) error {

	if r.Method != http.MethodGet {
		return fmt.Errorf("method '%s' not recognized", r.Method)
	}

	var err error
	if f.fizz, err = checkString(r.URL.Query(), "string1"); err != nil {
		return err
	}
	if f.buzz, err = checkString(r.URL.Query(), "string2"); err != nil {
		return err
	}
	if f.fizzStep, err = checkInt(r.URL.Query(), "int1"); err != nil {
		return err
	}
	if f.buzzStep, err = checkInt(r.URL.Query(), "int2"); err != nil {
		return err
	}
	if f.limit, err = checkInt(r.URL.Query(), "limit"); err != nil {
		return err
	}
	if f.limit == 0 {
		return fmt.Errorf("parameter 'limit' could not be zero")
	}
	return nil
}

func writeJson(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		panic(err)
	}
}

func writeError(w http.ResponseWriter, err error) {
	http.Error(w, fmt.Sprintf("could not fizz-buzz: %v", err), http.StatusBadRequest)
}

func isError(err error) string {
	if err != nil {
		return "\033[22;31mERROR \033[0m"
	}
	return ""
}

func main() {

	http.HandleFunc("/fizzbuzz", func(w http.ResponseWriter, r *http.Request) {
		received := time.Now().UTC()
		var err error
		defer func(t time.Time) {
			log.Printf("%s[%s] %s %v", isError(err), r.Method, r.URL.String(), time.Since(t))
		}(received)
		f := fizzBuzz{}
		if err = checkFizzBuzz(&f, w, r); err != nil {
			writeError(w, err)
			return
		}
		writeJson(w, result(&f), http.StatusOK)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
