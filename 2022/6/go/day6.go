// Advent of code 2022, day 6
// https://adventofcode.com/2022/day/6
package main

import (
	"log"
	"os"
)

const inputPath = "../input.txt"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Keeps track of radio properties. Has method for finding marker in buffer
type Radio struct {
	buffer []byte
}

func NewRadio(buffer *[]byte) *Radio {
	radio := Radio{buffer: *buffer}
	return &radio
}

// Returns buffer length before finding marker of given length
func (radio *Radio) FindMarker(length int) *int {
	for i := 0; i < len(radio.buffer)-length; i++ {
		bufferPart := make(map[byte]struct{})
		for j := 0; j < length; j++ {
			bufferPart[radio.buffer[i+j]] = struct{}{}
		}
		if len(bufferPart) == length {
			processedChars := i + length
			return &processedChars
		}
	}
	panic("no marker found in buffer")
}

// Parses puzzle input from txt file.
// Returns []byte.
func readInput(path string) *[]byte {
	inputBytes, err := os.ReadFile(path)
	check(err)
	return &inputBytes
}

// Part 1: How many characters need to be processed before the first
// 4-character start-of-packet marker is detected?
func solvePart1(inputBytes *[]byte) (*Radio, *int) {
	radio := NewRadio(inputBytes)
	answer1 := radio.FindMarker(4)
	return radio, answer1
}

// Part 2: How many characters need to be processed before the first
// 14-character start-of-packet marker is detected?
func solvePart2(radio *Radio) *int {
	answer2 := radio.FindMarker(14)
	return answer2
}

// Solves puzzle parts, split out for benchmarking
func solvePuzzle() (*int, *int) {
	inputBytes := readInput(inputPath)
	radio, answer1 := solvePart1(inputBytes)
	answer2 := solvePart2(radio)
	return answer1, answer2
}

// Reads input, solves puzzle parts and logs answers
func main() {
	answer1, answer2 := solvePuzzle()
	log.Printf("Characters processed before 4-length marker: %v\n", *answer1)
	log.Printf("Characters processed before 14-length marker: %v\n", *answer2)
}
