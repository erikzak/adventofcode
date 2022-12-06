package main

import "testing"

// Benchmark main function
func BenchmarkSolvePuzzle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePuzzle()
	}
}
