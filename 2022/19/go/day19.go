// Advent of code 2022, day 19
// https://adventofcode.com/2022/day/19
package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Costs type. Map of robot types pointing to array of cost: [ore, clay, obsidian]
type Costs map[int][3]int // (0 = ore, 1 = clay, 2 = obsidian, 3 = geode)

// Keeps track of blueprints, with costs for robots and eventually max geodes cracked
type Blueprint struct {
	id int
	// Robot types pointing to array of cost: [ore, clay, obsidian]
	costs     Costs // (0 = ore, 1 = clay, 2 = obsidian, 3 = geode)
	maxGeodes int
	quality   int
}

func NewBlueprint(id int, costs Costs) *Blueprint {
	blueprint := Blueprint{id: id, costs: costs}
	return &blueprint
}

// Parses puzzle input from txt file.
// Returns slice of blueprint instances
func readInput(path string) []*Blueprint {
	inputBytes, err := os.ReadFile(path)
	check(err)
	lines := strings.Split(strings.ReplaceAll(string(inputBytes), "\r\n", "\n"), "\n")
	blueprints := []*Blueprint{}
	for i, line := range lines {
		costs := Costs{}
		split := strings.Split(line, ".")
		oreRobotOreCost := byteToInt(split[0][len(split[0])-2])
		costs[0] = [3]int{oreRobotOreCost, 0, 0}
		clayRobotOreCost := byteToInt(split[1][len(split[1])-2])
		costs[1] = [3]int{clayRobotOreCost, 0, 0}
		blueprint := NewBlueprint(i+1, costs)
		blueprints = append(blueprints, blueprint)
	}
	return blueprints
}

func byteToInt(s byte) int {
	value, err := strconv.Atoi(string(s))
	check(err)
	return value
}

// Part 1: What do you get if you add up the quality level of all of the blueprints in your list?
func solvePart1(blueprints []*Blueprint) int {
	sumQuality := 0
	return sumQuality
}

// Part 2: ?
func solvePart2(blueprints []*Blueprint) int {
	return 0
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
	log.Printf("Quality level of all blueprints: %v\n", answer1)
	log.Printf("?: %v\n", answer2)
}
