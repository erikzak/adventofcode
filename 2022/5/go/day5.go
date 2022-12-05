// Advent of code 2022, day 5
// https://adventofcode.com/2022/day/5
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/erikzak/adventofcode/2022/5/crane"
)

const inputPath = "../input.txt"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Parses puzzle input from txt file.
// Returns a crane with stack setup and planned moves
func readInput(path string, crateLines int) crane.Crane {
	inputBytes, err := os.ReadFile(path)
	check(err)
	lines := strings.Split(string(inputBytes), "\n")

	// Stack config in first x lines
	crates := [][]rune{}
	for _, line := range lines[:crateLines] {
		for i := 1; i < len(line); i += 4 {
			stackIdx := (i - 1) / 4
			crate := rune(line[i])
			if !unicode.IsSpace(crate) {
				for len(crates) <= stackIdx {
					crates = append(crates, []rune{})
				}
				crates[stackIdx] = append(crates[stackIdx], crate)
			}
		}
	}
	// Get stack ids
	stacks := map[rune]crane.Stack{}
	for i, field := range strings.Fields(lines[crateLines]) {
		id := rune(field[0])
		// First item in crate slice is top crate
		stacks[id] = crane.NewStack(id, crates[i])
	}
	// Crate moves in the remaining lines (after 1 blank)
	crane := crane.NewCrane(stacks, lines[crateLines+2:])
	return crane
}

// Reads input and solves puzzle parts
func main() {
	// Part 1: CrateMover 9000 - what crate ends up on top of each stack?
	crane := readInput(inputPath, 8)
	crane.ExecuteMoves(true)
	fmt.Printf("CrateMover 9000 - Crates on top of each stack: %v\n", crane.GetTopCrates())

	// Part 2: CrateMover 9001 - what crate ends up on top of each stack?
	crane = readInput(inputPath, 8)
	crane.ExecuteMoves(false)
	fmt.Printf("CrateMover 9001 - Crates on top of each stack: %v\n", crane.GetTopCrates())
}
