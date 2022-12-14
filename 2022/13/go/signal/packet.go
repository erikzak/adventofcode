package signal

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Checks if []any is int
func isNum(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Float64
}

// Packet type with any value
type Packet struct {
	Values []any
}

func NewPacket(input string) Packet {
	// Parse packet string through JSON into either list or int
	var values []any
	err := json.Unmarshal([]byte(input), &values)
	if err != nil {
		panic("unable to parse packet input:" + fmt.Sprintf("%v", values))
	}
	return Packet{Values: values}
}

// Recursively compares packet values
func (left Packet) IsOrdered(right Packet) bool {
	result := compare(left.Values, right.Values)
	return result >= 0
}

// Compares left packet value to right
func compare(left any, right any) float64 {
	leftIsNum, rightIsNum := isNum(left), isNum(right)
	if leftIsNum && rightIsNum {
		// Compare numbers directly. Parsed JSON numeric output is float64
		return right.(float64) - left.(float64)
	}

	if !leftIsNum && !rightIsNum {
		// Assign correct slice type to values
		left := left.([]any)
		right := right.([]any)
		for i := range left {
			if len(right) < i+1 {
				return -1
			}
			result := compare(left[i], right[i])
			if result != 0 {
				return result
			}
		}
		if len(left) < len(right) {
			return 1
		}
		return 0
	}

	if leftIsNum {
		return compare([]any{left}, right)
	} else if rightIsNum {
		return compare(left, []any{right})
	}
	panic("noping out")
}
