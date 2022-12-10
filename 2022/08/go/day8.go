// Advent of code 2022, day 8
// https://adventofcode.com/2022/day/8
//
// My OOP approach is to generate rows and columns of Tree objects, add them to
// a Forest container and calculate visibility per tree in the grid.
package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/erikzak/adventofcode/2022/8/foresting"
)

const inputPath = "../input.txt"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Parses puzzle input from txt file.
// Returns Forest object initialized with input tree grid of rows and columns
func readInput(path string) (forest *foresting.Forest) {
	inputBytes, err := os.ReadFile(path)
	check(err)
	lines := strings.Split(string(inputBytes), "\n")

	// Rows from input
	rows := make([][]*foresting.Tree, len(lines))
	for i, line := range lines {
		line = strings.Trim(line, "\n")
		row := make([]*foresting.Tree, len(line))
		for j, value := range line {
			height, err := strconv.Atoi(string(value))
			check(err)
			row[j] = foresting.NewTree(height)
		}
		rows[i] = row
	}
	// Columns from rows
	nCols := len(rows[0])
	columns := make([][]*foresting.Tree, nCols)
	for colIdx := 0; colIdx < nCols; colIdx++ {
		column := make([]*foresting.Tree, nCols)
		for i, row := range rows {
			column[i] = row[colIdx]
		}
		columns[colIdx] = column
	}

	forest = foresting.NewForest(rows, columns)
	return forest
}

// Part 1: how many trees are visible from outside the grid?
func solvePart1(forest *foresting.Forest) (visibleTrees int) {
	visibleTrees = forest.CalculateTreeVisibility()
	return visibleTrees
}

// Part 2: ?
func solvePart2(forest *foresting.Forest) int {
	highestScenicScore := forest.CalculateTreeScenicScores()
	return highestScenicScore
}

// Solves puzzle parts. Split up for benchmarking
func solvePuzzle() (int, int) {
	root := readInput(inputPath)
	answer1 := solvePart1(root)
	answer2 := solvePart2(root)
	return answer1, answer2
}

// Reads input, solves puzzle parts and logs answers
func main() {
	answer1, answer2 := solvePuzzle()
	log.Printf("Trees visible from outside the grid: %v\n", answer1)
	log.Printf("Highest scenic score possible: %v\n", answer2)
}
