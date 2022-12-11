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
	operation       func(int) int
	testValue       int
	onTrue          int
	onFalse         int
	InspectionCount int
}

// Inits new monkey based on input text
func NewMonkey(
	itemStr string, operationStr string,
	testStr string, onTrueStr string, onFalseStr string) *Monkey {
	// Items
	items := []int{}
	for _, itemValue := range strings.Split(itemStr, ", ") {
		value := stringToInt(itemValue)
		items = append(items, value)
	}
	// On true/false test values
	onTrue := stringToInt((strings.Split(onTrueStr, " ")[3]))
	onFalse := stringToInt((strings.Split(onFalseStr, " ")[3]))
	// Test value
	testValue := stringToInt(strings.Split(testStr, " ")[2])

	// Init monkey
	monkey := Monkey{items: items, testValue: testValue, onTrue: onTrue, onFalse: onFalse}

	// Operation method
	operationParts := strings.Split(operationStr, " ")
	operator := operationParts[3]
	operationValue, err := strconv.Atoi(operationParts[4])
	if operator == "+" {
		if err == nil {
			monkey.operation = func(item int) int {
				return item + operationValue
			}
		} else {
			monkey.operation = func(item int) int {
				return item + item
			}
		}
	} else if operator == "*" {
		if err == nil {
			monkey.operation = func(item int) int {
				return item * operationValue
			}
		} else {
			monkey.operation = func(item int) int {
				return item * item
			}
		}
	} else {
		panic("undefined operator " + string(operator))
	}
	return &monkey
}

// Tests if item worry level is divisible by monkey test value
func (monkey *Monkey) test(item int) bool {
	return (item % monkey.testValue) == 0
}
