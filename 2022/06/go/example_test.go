// Tests puzzle example data
package main

import (
	"testing"
)

const testPath = "../test.txt"

// Tests part 1 example data
func TestPart1Example(t *testing.T) {
	want := 11
	inputBytes := readInput(testPath)
	radio := NewRadio(inputBytes)
	charactersProcessed := radio.FindMarker(4)
	if *charactersProcessed != want {
		t.Fatalf(`radio.FindMarker() = %v, want %v`, *charactersProcessed, want)
	}
}

// Tests part 2 example data
func TestPart2Example(t *testing.T) {
	want := 26
	inputBytes := readInput(testPath)
	radio := NewRadio(inputBytes)
	charactersProcessed := radio.FindMarker(14)
	if *charactersProcessed != want {
		t.Fatalf(`radio.FindMarker() = %v, want %v`, *charactersProcessed, want)
	}
}
