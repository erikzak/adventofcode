// Advent of Code, day 1.
// https://adventofcode.com/2022/day/1
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Parses puzzle input from txt file.
// Returns a list of sum calories per elf.
func readInput(inputPath string) (elfCalories []int) {
	inputBytes, err := os.ReadFile(inputPath)
	check(err)
	input := string(inputBytes)

	elfStrings := strings.Split(input, "\n\n")
	for _, elfString := range elfStrings {
		elfString = strings.TrimSpace(elfString)
		if elfString == "" {
			continue
		}
		foodItems := strings.Split(elfString, "\n")
		sumCalories := 0
		for _, foodItem := range foodItems {
			calories, err := strconv.Atoi(foodItem)
			check(err)
			sumCalories += calories
		}
		elfCalories = append(elfCalories, sumCalories)
	}
	return elfCalories
}

// Parses input as txt into list then sorts to find max calories
func main() {
	elfCalories := readInput(inputPath)
	sort.Ints(elfCalories)

	// Part 1
	fmt.Println(elfCalories[len(elfCalories)-1])

	// Part 2
	topThreeElves := elfCalories[len(elfCalories)-3:]
	sum := 0
	for _, calories := range topThreeElves {
		sum += calories
	}
	fmt.Println(sum)
}
