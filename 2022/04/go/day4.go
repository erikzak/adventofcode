// Advent of code 2022, day 4
// https://adventofcode.com/2022/day/4
package main

import (
	"fmt"
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

// A small often mischievous fairy
type Elf struct {
	assignment string
	minSection int
	maxSection int
}

// Checks if an elf's section(s) fully contain another elf's section(s)
func (elf Elf) fullyContainsSections(otherElf Elf) bool {
	if elf.minSection <= otherElf.minSection &&
		elf.maxSection >= otherElf.maxSection {
		return true
	}
	return false
}

// Checks if an elf's section(s) overlap another elf's section(s)
func (elf Elf) overlapsSections(otherElf Elf) bool {
	if elf.minSection <= otherElf.maxSection &&
		elf.maxSection >= otherElf.minSection {
		return true
	}
	return false
}

// Elf constructor. Gets min/max assignment section
func newElf(assignment string) Elf {
	elf := Elf{assignment: assignment}
	split := strings.Split(assignment, "-")
	var err error
	elf.minSection, err = strconv.Atoi(split[0])
	check(err)
	elf.maxSection, err = strconv.Atoi(split[1])
	check(err)
	return elf
}

// Parses puzzle input from txt file.
// Returns a slice of elf pair slices
func readInput(inputPath string) (elfPairs [][]Elf) {
	inputBytes, err := os.ReadFile(inputPath)
	check(err)
	input := string(inputBytes)

	for _, line := range strings.Split(input, "\n") {
		assignments := strings.Split(strings.TrimSpace(line), ",")
		elf1 := newElf(assignments[0])
		elf2 := newElf(assignments[1])
		elfPair := []Elf{elf1, elf2}
		elfPairs = append(elfPairs, elfPair)
	}
	return elfPairs
}

// Reads input and solves puzzle parts
func main() {
	elfPairs := readInput(inputPath)

	// Part 1: in how many assignment pairs does one range fully contain the other?
	nContains := 0
	for _, elfPair := range elfPairs {
		if elfPair[0].fullyContainsSections(elfPair[1]) ||
			elfPair[1].fullyContainsSections(elfPair[0]) {
			nContains += 1
		}
	}
	fmt.Printf("Assignment pairs where one fully contains the other: %d\n", nContains)

	// Part 2: in how many assignment pairs do the ranges overlap?
	nOverlaps := 0
	for _, elfPair := range elfPairs {
		if elfPair[0].overlapsSections(elfPair[1]) {
			nOverlaps += 1
		}
	}
	fmt.Printf("Assignment pairs where sections overlap: %d\n", nOverlaps)
}
