package main

import "testing"

func newfb(limit int) FizzBuzz {
	return &fizzBuzz{fizz: "fizz", buzz: "buzz", fizzStep: 3, buzzStep: 5, limit: limit}
}

func benchmarkRender(limit int, b *testing.B) {
	f := newfb(limit)
	for n := 0; n < b.N; n++ {
		render(f)
	}
}

func benchmarkResult(limit int, b *testing.B) {
	f := newfb(limit)
	for n := 0; n < b.N; n++ {
		result(f)
	}
}

func BenchmarkRender100(b *testing.B) {
	benchmarkRender(100, b)
}

func BenchmarkRender1000(b *testing.B) {
	benchmarkRender(1000, b)
}

func BenchmarkRender10000(b *testing.B) {
	benchmarkRender(10000, b)
}

func BenchmarkRender100000(b *testing.B) {
	benchmarkRender(100000, b)
}

func BenchmarkRender1000000(b *testing.B) {
	benchmarkRender(1000000, b)
}

func BenchmarkRender10000000(b *testing.B) {
	benchmarkRender(10000000, b)
}

func BenchmarkResult100(b *testing.B) {
	benchmarkResult(100, b)
}

func BenchmarkResult1000(b *testing.B) {
	benchmarkResult(1000, b)
}

func BenchmarkResult10000(b *testing.B) {
	benchmarkResult(10000, b)
}

func BenchmarkResult100000(b *testing.B) {
	benchmarkResult(100000, b)
}

func BenchmarkResult1000000(b *testing.B) {
	benchmarkResult(1000000, b)
}

func BenchmarkResult10000000(b *testing.B) {
	benchmarkResult(10000000, b)
}

