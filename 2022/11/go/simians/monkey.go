package simians

import (
	"log"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// Returns int of string
func stringToInt(str string) int {
	value, err := strconv.Atoi(str)
	check(err)
	return value
}

// Keeps track of monkey items and round logic
type Monkey struct {
	items           []int
	operator        string
	operationValue  string
	testValue       int
	onTrueMonkey    int
	onFalseMonkey   int
	InspectionCount int
}

// Inits new monkey based on input text
func NewMonkey(
	itemStr string, operation string,
	testStr string, onTrueStr string, onFalseStr string) *Monkey {
	// Items
	items := []int{}
	for _, itemValue := range strings.Split(itemStr, ", ") {
		value := stringToInt(itemValue)
		items = append(items, value)
	}
	// Operation method
	operationParts := strings.Split(operation, " ")
	operator := operationParts[3]
	operationValue := operationParts[4]
	// Test value
	testValue := stringToInt(strings.Split(testStr, " ")[2])
	// On true/false test values
	onTrueMonkey := stringToInt((strings.Split(onTrueStr, " ")[3]))
	onFalseMonkey := stringToInt((strings.Split(onFalseStr, " ")[3]))
	monkey := Monkey{items: items, operator: operator, operationValue: operationValue,
		testValue: testValue, onTrueMonkey: onTrueMonkey, onFalseMonkey: onFalseMonkey}
	return &monkey
}

// Returns updated item worry level based on monkey inspection
func (monkey *Monkey) operation(item int) int {
	if monkey.operator == "+" {
		if monkey.operationValue == "old" {
			return item + item
		} else {
			return item + stringToInt(monkey.operationValue)
		}
	} else if monkey.operator == "*" {
		if monkey.operationValue == "old" {
			return item * item
		} else {
			return item * stringToInt(monkey.operationValue)
		}
	} else {
		panic("undefined operator " + string(monkey.operator))
	}
}

// Tests if item worry level is divisible by monkey test value
func (monkey *Monkey) test(item int) bool {
	return (item % monkey.testValue) == 0
}
