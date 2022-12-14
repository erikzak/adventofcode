// Advent of code 2022, day 13
// https://adventofcode.com/2022/day/13
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/erikzak/adventofcode/2022/13/signal"
)

const inputPath = "../input.txt"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Parses puzzle input from txt file.
// Returns signal with packet pairs.
func readInput(path string) (data signal.Signal) {
	inputBytes, err := os.ReadFile(path)
	check(err)
	packetPairInputs := strings.Split(strings.ReplaceAll(string(inputBytes), "\r\n", "\n"), "\n\n")
	packetPairs := [][]signal.Packet{}
	for _, pair := range packetPairInputs {
		packetStrings := strings.Split(pair, "\n")
		packetPair := []signal.Packet{
			signal.NewPacket(packetStrings[0]),
			signal.NewPacket(packetStrings[1]),
		}
		packetPairs = append(packetPairs, packetPair)
	}
	data = signal.NewSignal(packetPairs)
	return data
}

// Part 1: What is the sum of the indices of pairs in the right order?
func solvePart1(data signal.Signal) int {
	return data.CheckOrdering()
}

// Part 2: What is the decoder key for the distress signal?
func solvePart2(data signal.Signal) int {
	packets := []signal.Packet{
		signal.NewPacket("[[2]]"),
		signal.NewPacket("[[6]]"),
	}
	for _, packetPair := range data.PacketPairs {
		packets = append(packets, packetPair[0])
		packets = append(packets, packetPair[1])
	}
	sort.Slice(packets, func(i, j int) bool {
		return packets[i].IsOrdered(packets[j])
	})
	decoderKey := (getDividerIndex(packets, "[[2]]") + 1) * (getDividerIndex(packets, "[[6]]") + 1)
	return decoderKey
}

func getDividerIndex(packets []signal.Packet, divider string) int {
	for i, packet := range packets {
		if fmt.Sprintf("%v", packet.Values) == divider {
			return i
		}
	}
	panic("packet not found")
}

// Solves puzzle parts. Split up for benchmarking
func solvePuzzle() (int, int) {
	input := readInput(inputPath)
	answer1 := solvePart1(input)
	answer2 := solvePart2(input)
	return answer1, answer2
}

// Reads input, solves puzzle parts and logs answers
func main() {
	answer1, answer2 := solvePuzzle()
	log.Printf("Sum of indices of right order pairs: %v\n", answer1)
	log.Printf("Decoder key: %v\n", answer2)
}
