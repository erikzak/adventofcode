// Advent of Code, day 3
// https://adventofcode.com/2022/day/3
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const inputPath = "../input.txt"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Rucksack structure for keeping track of content, compartments, duplicates
// and calculating priority score.
type Rucksack struct {
	items        string
	itemMap      map[rune]struct{}
	compartments []string
	duplicates   map[rune]struct{}
	sumPriority  int
}

// Constructor for new rucksacks. Organizes items, finds duplicates and
// calculates sum priority
func newRucksack(items string) Rucksack {
	sack := Rucksack{items: items}
	sack.itemMap = make(map[rune]struct{})
	sack.duplicates = make(map[rune]struct{})

	// Generate item rune map
	for _, r := range items {
		sack.itemMap[r] = struct{}{}
	}

	// Process inventory
	sack.splitItemsIntoCompartments()
	sack.findDuplicates()
	sack.calculatePriority()
	return sack
}

// Splits items into two rucksack compartments
func (sack *Rucksack) splitItemsIntoCompartments() {
	midIdx := len(sack.items) / 2
	sack.compartments = append(sack.compartments, sack.items[:midIdx])
	sack.compartments = append(sack.compartments, sack.items[midIdx:])
}

// Searches compartments for duplicate values
func (sack *Rucksack) findDuplicates() {
	// Create rune map of second compartment for performance
	secondCompartment := make(map[rune]struct{})
	for _, r := range sack.compartments[1] {
		secondCompartment[r] = struct{}{}
	}
	// Look for duplicates
	for _, r := range sack.compartments[0] {
		if _, ok := secondCompartment[r]; ok {
			sack.duplicates[r] = struct{}{}
		}
	}
}

// Calculates priority of rucksack based on duplicates
func (sack *Rucksack) calculatePriority() {
	for r := range sack.duplicates {
		sack.sumPriority += getItemPriority(r)
	}
}

// Parses puzzle input from txt file.
// Returns a slice of rucksacks.
func readInput(inputPath string) (sacks []Rucksack) {
	inputBytes, err := os.ReadFile(inputPath)
	check(err)
	input := string(inputBytes)

	for _, line := range strings.Split(input, "\n") {
		items := strings.TrimSpace(line)
		sack := newRucksack(items)
		sacks = append(sacks, sack)
	}
	return sacks
}

// Returns the priority value of a given item
func getItemPriority(item rune) int {
	// Use unicode code points to designate priority
	if unicode.IsLower(item) {
		return int(item) - 96
	}
	return int(item) - 38 // -64 + 26
}

// Finds the item present in all rucksacks
func findBadge(sacks []Rucksack) rune {
	for r := range sacks[0].itemMap {
		for i := 1; i < len(sacks); i++ {
			if _, ok := sacks[i].itemMap[r]; !ok {
				break
			}
			if i == len(sacks)-1 {
				return r
			}
		}
	}
	panic("no badge found")
}

// Parses input as txt into slice of rucksacks with calculated priority scores
func main() {
	sacks := readInput(inputPath)

	// Part 1
	sumPriority := 0
	for _, sack := range sacks {
		sumPriority += sack.sumPriority
	}
	fmt.Printf("Sum of rucksack priorities: %d\n", sumPriority)

	// Part 2
	sumBadgePriority := 0
	// Process 3 and 3 rucksacks, find badge and add priority to sum
	for i := 0; i < len(sacks); i += 3 {
		badge := findBadge(sacks[i : i+3])
		sumBadgePriority += getItemPriority(badge)
	}
	fmt.Printf("Sum of badge priorities: %d\n", sumBadgePriority)
}
