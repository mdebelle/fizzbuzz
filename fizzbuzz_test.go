package main

import "testing"

func newfb(limit int) FizzBuzz {
	return &fizzBuzz{fizz: "fizz", buzz: "buzz", fizzStep: 3, buzzStep: 5, limit: limit}
}

func benchmarkAlgoOne(limit int, b *testing.B) {
	f := newfb(limit)
	for n := 0; n < b.N; n++ {
		algoOne(f)
	}
}

func benchmarkAlgoTwo(limit int, b *testing.B) {
	f := newfb(limit)
	for n := 0; n < b.N; n++ {
		result(f)
	}
}

func BenchmarkAlgoOne100(b *testing.B) {
	benchmarkAlgoOne(100, b)
}

func BenchmarkAlgoOne1000(b *testing.B) {
	benchmarkAlgoOne(1000, b)
}

func BenchmarkAlgoOne10000(b *testing.B) {
	benchmarkAlgoOne(10000, b)
}

func BenchmarkAlgoOne100000(b *testing.B) {
	benchmarkAlgoOne(100000, b)
}

func BenchmarkAlgoOne1000000(b *testing.B) {
	benchmarkAlgoOne(1000000, b)
}

func BenchmarkAlgoOne10000000(b *testing.B) {
	benchmarkAlgoOne(10000000, b)
}

func BenchmarkAlgoOne100000000(b *testing.B) {
	benchmarkAlgoOne(100000000, b)
}

func BenchmarkAlgoTwo100(b *testing.B) {
	benchmarkAlgoTwo(100, b)
}

func BenchmarkAlgoTwo1000(b *testing.B) {
	benchmarkAlgoTwo(1000, b)
}

func BenchmarkAlgoTwo10000(b *testing.B) {
	benchmarkAlgoTwo(10000, b)
}

func BenchmarkAlgoTwo100000(b *testing.B) {
	benchmarkAlgoTwo(100000, b)
}

func BenchmarkAlgoTwo1000000(b *testing.B) {
	benchmarkAlgoTwo(1000000, b)
}

func BenchmarkAlgoTwo10000000(b *testing.B) {
	benchmarkAlgoTwo(10000000, b)
}

func BenchmarkAlgoTwo100000000(b *testing.B) {
	benchmarkAlgoTwo(100000000, b)
}
