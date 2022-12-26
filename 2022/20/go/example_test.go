// Tests puzzle example data
package main

import (
	"testing"
)

// Tests part 1 against example data
func TestPart1Example(t *testing.T) {
	want := 3
	input := readInput("../test.txt")
	answer1 := solvePart1(input, true)
	if answer1 != want {
		t.Fatalf(`solvePart1() = %v, want %v`, answer1, want)
	}
}

// Tests part 2 against example data
func TestPart2Example(t *testing.T) {
	want := 1623178306
	input := readInput("../test.txt")
	answer2 := solvePart2(input, true)
	if answer2 != want {
		t.Fatalf(`solvePart2() = %v, want %v`, answer2, want)
	}
}
