// Advent of code 2022, day 10
// https://adventofcode.com/2022/day/10
package main

import (
	"log"
	"os"
	"strings"

	"github.com/erikzak/adventofcode/2022/10/handheld"
)

const inputPath = "../input.txt"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Parses puzzle input from txt file.
// Returns device instance with instructions executed
func readInput(path string) (device *handheld.Device) {
	inputBytes, err := os.ReadFile(path)
	check(err)
	instructions := strings.Split(strings.ReplaceAll(string(inputBytes), "\r\n", "\n"), "\n")
	device = handheld.NewDevice(instructions, 6, 40)
	interestingCycles := []int{20, 60, 100, 140, 180, 220}
	device.ExecuteInstructions(interestingCycles)
	return device
}

// Part 1: What is the sum of these six signal strengths?
func solvePart1(device *handheld.Device) int {
	return device.GetSumSignalStrengths()
}

// Part 2: What eight capital letters appear on your CRT?
func solvePart2(device *handheld.Device) []string {
	return device.GetImage()
}

// Solves puzzle parts. Split up for benchmarking
func solvePuzzle() (int, []string) {
	input := readInput(inputPath)
	answer1 := solvePart1(input)
	answer2 := solvePart2(input)
	return answer1, answer2
}

// Reads input, solves puzzle parts and logs answers
func main() {
	answer1, answer2 := solvePuzzle()
	log.Printf("Sum of six signal strengths: %v\n", answer1)
	log.Print("Eight capital letters appear:\n")
	for _, line := range answer2 {
		log.Printf("\t%v\n", line)
	}
}
