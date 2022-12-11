// Advent of code 2022, day 11
// https://adventofcode.com/2022/day/11
package main

import (
	"log"
	"os"
	"sort"
	"strings"

	"github.com/erikzak/adventofcode/2022/11/simians"
)

const inputPath = "../input.txt"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Parses puzzle input from txt file.
// Returns a troop of monkeys
func readInput(path string) (troop *simians.Troop) {
	inputBytes, err := os.ReadFile(path)
	check(err)
	inputString := string(inputBytes)
	monkeyParts := strings.Split(strings.ReplaceAll(inputString, "\r\n", "\n"), "\n\n")

	monkeys := []*simians.Monkey{}
	for _, monkeyPart := range monkeyParts {
		lines := strings.Split(monkeyPart, "\n")
		items := strings.Split(lines[1], ": ")[1]
		operation := strings.Split(lines[2], ": ")[1]
		test := strings.Split(lines[3], ": ")[1]
		onTrue := strings.Split(lines[4], ": ")[1]
		onFalse := strings.Split(lines[5], ": ")[1]
		monkeys = append(monkeys, simians.NewMonkey(items, operation, test, onTrue, onFalse))
	}
	return simians.NewTroop(monkeys)
}

func sortMonkeysByInspectionCount(troop *simians.Troop) {
	sort.Slice(troop.Monkeys, func(i, j int) bool {
		return troop.Monkeys[i].InspectionCount > troop.Monkeys[j].InspectionCount
	})
}

// Part 1: Level of monkey business after 20 rounds of stuff-slinging simian shenanigans?
func solvePart1(troop *simians.Troop) (monkeyBusinessLevel int) {
	troop.Go(20)
	sortMonkeysByInspectionCount(troop)
	monkeyBusinessLevel = troop.Monkeys[0].InspectionCount * troop.Monkeys[1].InspectionCount
	return monkeyBusinessLevel
}

// Part 2: What is the level of monkey business after 10000 rounds with no reduction?
func solvePart2(troop *simians.Troop) (monkeyBusinessLevel int) {
	troop.WorryReduces = false
	troop.Go(10000)
	sortMonkeysByInspectionCount(troop)
	monkeyBusinessLevel = troop.Monkeys[0].InspectionCount * troop.Monkeys[1].InspectionCount
	return monkeyBusinessLevel
}

// Solves puzzle parts. Split up for benchmarking
func solvePuzzle() (int, int) {
	input := readInput(inputPath)
	answer1 := solvePart1(input)
	input = readInput(inputPath)
	answer2 := solvePart2(input)
	return answer1, answer2
}

// Reads input, solves puzzle parts and logs answers
func main() {
	answer1, answer2 := solvePuzzle()
	log.Printf("Monkey business after 20 rounds: %v\n", answer1)
	log.Printf("Monkey business after 10000 rounds: %v\n", answer2)
}
