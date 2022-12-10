// Tests puzzle example data
package main

import (
	"testing"
)

const testPath = "../test.txt"

// Tests part 1 example data
func TestPart1Example(t *testing.T) {
	want := "CMZ"
	crane := readInput(testPath, 3)
	crane.ExecuteMoves(true)
	topCrates := crane.GetTopCrates()
	if topCrates != want {
		t.Fatalf(`crane.GetTopCrates() = "%v", want "%v"`, topCrates, want)
	}
}

// Tests part 2 example data
func TestPart2Example(t *testing.T) {
	want := "MCD"
	crane := readInput(testPath, 3)
	crane.ExecuteMoves(false)
	topCrates := crane.GetTopCrates()
	if topCrates != want {
		t.Fatalf(`crane.GetTopCrates() = "%v", want "%v"`, topCrates, want)
	}
}
