// Tests puzzle example data
package main

import (
	"testing"
)

const testPath = "../test.txt"

// Tests part 1 against example data
func TestPart1Example(t *testing.T) {
	want := 95437
	root := readInput(testPath)
	answer1 := solvePart1(root)
	if answer1 != want {
		t.Fatalf(`solvePart1() = %v, want %v`, answer1, want)
	}
}

// Tests part 2 against example data
func TestPart2Example(t *testing.T) {
	want := 24933642
	root := readInput(testPath)
	answer2 := solvePart2(root)
	if answer2 != want {
		t.Fatalf(`solvePart2() = %v, want %v`, answer2, want)
	}
}
