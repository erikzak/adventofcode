// Advent of code 2022, day 6
// https://adventofcode.com/2022/day/6
package main

import (
	"log"
	"os"
	"time"
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

// Reads input and solves puzzle parts
func main() {
	start := time.Now()
	inputBytes := readInput(inputPath)

	// Part 1: How many characters need to be processed before the first
	// 4-character start-of-packet marker is detected?
	radio := NewRadio(inputBytes)
	charactersProcessed := radio.FindMarker(4)
	log.Printf("Characters processed before 4-length marker: %v\n", *charactersProcessed)

	// Part 2: How many characters need to be processed before the first
	// 14-character start-of-packet marker is detected?
	radio = NewRadio(inputBytes)
	charactersProcessed = radio.FindMarker(14)
	log.Printf("Characters processed before 14-length marker: %v\n", *charactersProcessed)

	// Execution time
	elapsed := time.Since(start)
	log.Printf("Execution time: %s\n", elapsed)
}
