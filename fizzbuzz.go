package main

import (
	"fmt"
	"sync"
)

// FizzBuzz to manipulatr fizzbuzz struct
type FizzBuzz interface {
	Convert(i int) string
	Limit() int
}

type fizzBuzz struct {
	motifs map[int]string
	limit  int
}

func (f *fizzBuzz) Limit() int { return f.limit }

// Convert take an int as parameters and must return
// fizz field of fizzBuzz when i is only a multiple of fizzStep field of fizzBuzz
// buzz field of fizzBuzz when i is only a multiple of buzzStep field of fizzBuzz
// concatenation of fizz and buzz fields of fizzBuzz when i is a multiple of fizzStep and buzzStep field of fizzBuzz
// in other case it will return i as a string
func (f *fizzBuzz) Convert(i int) string {
	var str string
	var found bool
	for k, v := range f.motifs {
		if i%k == 0 {
			str += v
			found = true
		}
	}
	if !found {
		return fmt.Sprint(i)
	}
	return str
}

func render(f FizzBuzz) []string {
	out := make([]string, f.Limit())
	for i, j := 0, f.Limit()-1; i <= f.Limit()/2; i, j = i+1, j-1 {
		out[i], out[j] = f.Convert(i+1), f.Convert(j+1)
	}
	return out
}

func result(f FizzBuzz) []string {
	out := make([]string, f.Limit())

	if f.Limit() <= 100 {
		return render(f)
	}

	nbgoroutine := f.Limit() / 100
	if nbgoroutine > 100 {
		nbgoroutine = 100
	}

	wg := &sync.WaitGroup{}

	apply := func(f FizzBuzz, min, max int) {
		length := max - min
		for i, j := min, max-1; i <= min+(length/2); i, j = i+1, j-1 {
			out[i], out[j] = f.Convert(i+1), f.Convert(j+1)
		}
		wg.Done()
	}

	for i := 0; i < nbgoroutine; i++ {
		wg.Add(1)
		go apply(f, f.Limit()*i/nbgoroutine, f.Limit()*(i+1)/nbgoroutine)
	}

	wg.Wait()
	return out
}
