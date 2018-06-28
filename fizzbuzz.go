package main

import (
	"fmt"
	"sync"
)

const (
	algoNumberOne = 1
	algoNumberTwo = 2
)

type FizzBuzz interface {
	Convert(i int) string
	Limit() int
}

type fizzBuzz struct {
	fizz     string
	buzz     string
	fizzStep int
	buzzStep int
	limit    int
}

func (f *fizzBuzz) Limit() int { return f.limit }

// Convert take an int as parameters and must return
// fizz field of fizzBuzz when i is only a multiple of fizzStep field of fizzBuzz
// buzz field of fizzBuzz when i is only a multiple of buzzStep field of fizzBuzz
// concatenation of fizz and buzz fields of fizzBuzz when i is a multiple of fizzStep and buzzStep field of fizzBuzz
// in other case it will return i as a string
func (f *fizzBuzz) Convert(i int) string {
	if i%f.fizzStep == 0 {
		str := f.fizz
		if i%f.buzzStep == 0 {
			str += f.buzz
		}
		return str
	} else if i%f.buzzStep == 0 {
		return f.buzz
	}
	return fmt.Sprint(i)
}

func algoOne(f FizzBuzz) []string {
	out := make([]string, f.Limit())
	for i, j := 0, f.Limit()-1; i <= f.Limit()/2; i, j = i+1, j-1 {
		out[i], out[j] = f.Convert(i+1), f.Convert(j+1)
	}
	return out
}

func result(f FizzBuzz) []string {
	out := make([]string, f.Limit())
	var wg sync.WaitGroup

	if f.Limit() <= 100 {
		return algoOne(f)
	}

	nbgoroutine := f.Limit() / 100

	wg.Add(nbgoroutine)
	apply := func(f FizzBuzz, min, max int) {
		for i, j := min, max-1; i <= max/2; i, j = i+1, j-1 {
			out[i], out[j] = f.Convert(i+1), f.Convert(j+1)
		}
		wg.Done()
	}

	for i := 0; i < nbgoroutine; i++ {
		go apply(f, f.Limit()*i/nbgoroutine, f.Limit()*(i+1)/nbgoroutine)
	}

	wg.Wait()
	return out
}
