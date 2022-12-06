// Advent of code 2022, day 6
// https://adventofcode.com/2022/day/6
package main

import (
	"fmt"
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
// Returns a radio instance with methods to find markers.
func readInput(path string) *Radio {
	inputBytes, err := os.ReadFile(path)
	check(err)
	radio := NewRadio(&inputBytes)
	return radio
}

// Reads input and solves puzzle parts
func main() {
	// Part 1: How many characters need to be processed before the first
	// 4-character start-of-packet marker is detected?
	radio := readInput(inputPath)
	charactersProcessed := radio.FindMarker(4)
	fmt.Printf("Characters processed before 4-length marker: %v\n", *charactersProcessed)

	// Part 2: How many characters need to be processed before the first
	// 14-character start-of-packet marker is detected?
	radio = readInput(inputPath)
	charactersProcessed = radio.FindMarker(14)
	fmt.Printf("Characters processed before 14-length marker: %v\n", *charactersProcessed)
}
