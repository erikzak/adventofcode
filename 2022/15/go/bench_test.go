package main

import "testing"

// Benchmark full solve
func BenchmarkSolvePuzzle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solvePuzzle()
	}
}

// Benchmark input parsing
func BenchmarkReadInput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		readInput(inputPath)
	}
}

// Benchmark part 1
func BenchmarkSolvePart1(b *testing.B) {
	input := readInput(inputPath)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solvePart1(input, false)
	}
}

// Benchmark part 2
func BenchmarkSolvePart2(b *testing.B) {
	input := readInput(inputPath)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solvePart2(input, false)
	}
}
