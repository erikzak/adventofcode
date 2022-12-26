// Advent of code 2022, day 20
// https://adventofcode.com/2022/day/20
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

// File class, keeps track of numbers and mixing methods
type File struct {
	values  map[int]*Number
	orders  map[int]*Number
	indexes map[int]*Number
	count   int
}

func NewFile(numbers []*Number) File {
	values := map[int]*Number{}
	orders := map[int]*Number{}
	indexes := map[int]*Number{}
	for _, number := range numbers {
		values[number.value] = number
		orders[number.order] = number
		indexes[number.index] = number
	}
	file := File{values: values, orders: orders, indexes: indexes, count: len(numbers)}
	return file
}

// Mixes file to decrypt. Modifier is applied to value before index shuffling
func (file File) Mix(modifier int, debug bool) {
	if debug {
		fmt.Print("Initial arrangement:\n")
		fmt.Print(file.print() + "\n\n")
	}
	for i := 0; i < file.count; i++ {
		number := file.orders[i]
		if number.value == 0 {
			if debug {
				fmt.Printf("0 does not move:\n%v\n\n", file.print())
			}
			continue
		}
		oldIdx := int(number.index)
		number.index = (number.index + number.value*modifier) % (file.count - 1)
		for number.index < 0 {
			number.index += file.count - 1
		}
		if number.index == 0 {
			number.index = file.count - 1
		} else if number.index == file.count-1 {
			number.index = 0
		}
		if debug {
			fmt.Printf("(%v + %v = %v)\n", oldIdx, number.value*modifier%file.count, number.index)
			if number.index == file.count-1 {
				fmt.Printf(
					"%v moves between %v and %v:\n",
					number.value*modifier,
					file.indexes[number.index].value*modifier,
					file.indexes[0].value*modifier,
				)
			} else {
				fmt.Printf("%v moves between %v and %v:\n",
					number.value*modifier,
					file.indexes[number.index].value*modifier,
					file.indexes[number.index+1].value*modifier,
				)
			}
		}
		// Update affected indexes
		affected := []*Number{}
		if oldIdx < number.index {
			for idx := oldIdx + 1; idx <= number.index; idx++ {
				num := file.indexes[idx]
				num.index--
				affected = append(affected, num)
			}
		} else if number.index < oldIdx {
			for idx := number.index; idx < oldIdx; idx++ {
				num := file.indexes[idx]
				num.index++
				affected = append(affected, num)
			}
		}
		for _, num := range affected {
			file.indexes[num.index] = num
		}
		file.indexes[number.index] = number
		if debug {
			fmt.Print(file.print() + "\n\n")
		}
	}
}

func (file File) GetValue(idx int) int {
	return file.indexes[idx%file.count].value
}

func (file File) print() string {
	numbers := []string{}
	for i := 0; i < file.count; i++ {
		numbers = append(numbers, fmt.Sprint(file.indexes[i].value))
	}
	return strings.Join(numbers, ", ")
}

// Number class, keeps track of value, mixing order and current index
type Number struct {
	value int
	order int
	index int
}

func NewNumber(value int, order int, index int) *Number {
	return &Number{value: value, order: order, index: index}
}

// Parses puzzle input from txt file.
// Returns maps of Number instances referenced by order, value and (mixed) index
func readInput(path string) File {
	inputBytes, err := os.ReadFile(path)
	check(err)
	lines := strings.Split(strings.ReplaceAll(string(inputBytes), "\r\n", "\n"), "\n")

	numbers := []*Number{}
	for i, line := range lines {
		value, err := strconv.Atoi(line)
		check(err)
		number := NewNumber(value, i, i)
		numbers = append(numbers, number)
	}
	file := NewFile(numbers)
	return file
}

// Part 1: What is the sum of the three numbers that form the grove coordinates?
func solvePart1(file File, debug bool) int {
	file.Mix(1, debug)
	zeroIdx := file.values[0].index
	answer := file.GetValue(zeroIdx+1000) +
		file.GetValue(zeroIdx+2000) +
		file.GetValue(zeroIdx+3000)
	if debug {
		fmt.Printf(
			"\nAnswer: %v + %v + %v = %v\n",
			file.GetValue(zeroIdx+1000),
			file.GetValue(zeroIdx+2000),
			file.GetValue(zeroIdx+3000),
			answer,
		)
	}
	return answer
}

// Part 2: What is the sum of the three keyed numbers that form the grove coordinates?
func solvePart2(file File, debug bool) int {
	key := 811589153
	if debug {
		fmt.Print("Initial arrangement:\n")
		fmt.Printf("%v\n", file.print())
	}
	for i := 0; i < 10; i++ {
		file.Mix(key, false)
		if debug {
			fmt.Printf("After %v round(s) of mixing:\n", i+1)
			fmt.Printf("%v\n\n", file.print())
		}
	}
	zeroIdx := file.values[0].index
	answer := file.GetValue(zeroIdx+1000)*key +
		file.GetValue(zeroIdx+2000)*key +
		file.GetValue(zeroIdx+3000)*key
	if debug {
		fmt.Printf(
			"\nAnswer: %v + %v + %v = %v\n\n",
			file.GetValue(zeroIdx+1000)*key,
			file.GetValue(zeroIdx+2000)*key,
			file.GetValue(zeroIdx+3000)*key,
			answer,
		)
	}
	return answer
}

// Solves puzzle parts. Split up for benchmarking
func solvePuzzle() (int, int) {
	input := readInput(inputPath)
	answer1 := solvePart1(input, false)
	input = readInput(inputPath)
	answer2 := solvePart2(input, false)
	return answer1, answer2
}

// Reads input, solves puzzle parts and logs answers
func main() {
	answer1, answer2 := solvePuzzle()
	log.Printf("Sum of the three numbers: %v\n", answer1)
	log.Printf("Sum of the three keyed numbers: %v\n", answer2)
}
