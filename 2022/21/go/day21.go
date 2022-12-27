// Advent of code 2022, day 21
// https://adventofcode.com/2022/day/21
package main

import (
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

const inputPath = "../input.txt"

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Keeps track of all monkeys, with methods for solving puzzle
type Troop struct {
	monkeys      map[string]*Monkey
	dependencies map[string][]string
}

func NewTroop(monkeys map[string]*Monkey, dependencies map[string][]string) *Troop {
	troop := Troop{monkeys: monkeys, dependencies: dependencies}
	return &troop
}

// Simulates monkey business, resolving all monkey operations
func (troop *Troop) Go() {
	for {
		if len(troop.dependencies) == 0 {
			break
		}
		for name, dependents := range troop.dependencies {
			if troop.monkeys[name].value == 0 {
				continue
			}
			for i, dependent := range dependents {
				troop.monkeys[dependent].Resolve(troop.monkeys)
				if troop.monkeys[dependent].value != 0 {
					dependents = remove(dependents, i)
				}
			}
			if len(dependents) == 0 {
				delete(troop.dependencies, name)
			}
		}
	}
}

// Recursively resolve monkey values
func (troop *Troop) resolve(name string) {
	monkey := troop.monkeys[name]
	if monkey.value != 0 {
		return
	}
	for _, dependency := range monkey.dependencies {
		troop.resolve(dependency)
	}
	monkey.Resolve(troop.monkeys)
}

// Removes item from slice
func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// Individual monkey stuff
type Monkey struct {
	name         string
	value        int
	job          string
	operator     string
	dependencies []string
}

func NewMonkey(name string, job string) *Monkey {
	monkey := Monkey{name: name, job: job, dependencies: []string{}}
	value, err := strconv.Atoi(job)
	if err == nil {
		monkey.value = value
	} else {
		split := strings.Fields(job)
		monkey.operator = split[1]
		monkey.dependencies = append(monkey.dependencies, split[0])
		monkey.dependencies = append(monkey.dependencies, split[2])
	}
	return &monkey
}

func (monkey *Monkey) Resolve(monkeys map[string]*Monkey) {
	if monkey.value != 0 {
		return
	}
	v0 := monkeys[monkey.dependencies[0]].value
	v1 := monkeys[monkey.dependencies[1]].value
	if v0 != 0 && v1 != 0 {
		if monkey.operator == "+" {
			monkey.value = v0 + v1
		} else if monkey.operator == "-" {
			monkey.value = v0 - v1
		} else if monkey.operator == "*" {
			monkey.value = v0 * v1
		} else if monkey.operator == "/" {
			monkey.value = v0 / v1
		} else {
			panic("unknown operator: " + monkey.operator)
		}
	}
}

// Parses puzzle input from txt file.
// Returns troop of monkeys
func readInput(path string) *Troop {
	inputBytes, err := os.ReadFile(path)
	check(err)
	lines := strings.Split(strings.ReplaceAll(string(inputBytes), "\r\n", "\n"), "\n")
	monkeys := map[string]*Monkey{}
	dependencies := map[string][]string{}
	for _, line := range lines {
		split := strings.Split(line, ": ")
		name := split[0]
		operation := split[1]
		monkey := NewMonkey(name, operation)
		if monkey.value == 0 {
			for _, dependency := range monkey.dependencies {
				_, created := dependencies[dependency]
				if !created {
					dependencies[dependency] = []string{name}
				} else {
					dependencies[dependency] = append(dependencies[dependency], name)
				}
			}
		}
		monkeys[name] = monkey
	}
	troop := NewTroop(monkeys, dependencies)
	return troop
}

// Part 1: What number will the monkey named root yell?
func solvePart1(troop *Troop) int {
	troop.resolve("root")
	return troop.monkeys["root"].value
}

// Part 2: What number do you yell to pass root's equality test?
// 7010269744524
func solvePart2(troop *Troop) int {
	// Figure out which of root's dependency trees has humn
	human := "humn"
	humanTree := getHumanTree(troop, human)
	reverse(humanTree)
	// Resolve other path and get first true value, then resolve other side with wanted value
	root := troop.monkeys["root"]
	troop.resolve(root.name)
	wanted := 0
	if root.dependencies[0] == humanTree[0] {
		wanted = troop.monkeys[root.dependencies[1]].value
	} else {
		wanted = troop.monkeys[root.dependencies[0]].value
	}
	for {
		current := humanTree[0]
		if current == human {
			return wanted
		}
		humanTree = humanTree[1:]
		monkey := troop.monkeys[current]
		input := 0
		if monkey.dependencies[0] == humanTree[0] {
			input = troop.monkeys[monkey.dependencies[1]].value
		} else {
			input = troop.monkeys[monkey.dependencies[0]].value
		}
		if monkey.operator == "+" {
			wanted = wanted - input
		} else if monkey.operator == "-" {
			// Non-commutative bullshit
			if monkey.dependencies[0] == humanTree[0] {
				wanted = wanted + input
			} else {
				wanted = -(wanted - input)
			}
		} else if monkey.operator == "*" {
			wanted = wanted / input
		} else if monkey.operator == "/" {
			// Non-commutative bullshit
			if monkey.dependencies[0] == humanTree[0] {
				wanted = wanted * input
			} else {
				wanted = input / wanted
			}
		}
	}
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func getHumanTree(troop *Troop, human string) []string {
	queue := [][]string{}
	queue = append(queue, []string{human})
	for {
		if len(queue) == 0 {
			break
		}
		humanTree := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		dependents := troop.dependencies[humanTree[len(humanTree)-1]]
		for _, dependent := range dependents {
			if dependent == "root" {
				return humanTree
			}
			branch := make([]string, len(humanTree))
			copy(branch, humanTree)
			branch = append(branch, dependent)
			queue = append(queue, branch)
		}
	}
	panic("no human tree found")
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
	log.Printf("root yells: %v\n", answer1)
	log.Printf("humn yells: %v\n", answer2)
}
