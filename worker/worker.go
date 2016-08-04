package worker

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Takes a list of strings containing integers and words and returns a
// sorted version of that list while preserving the type of each nth element
func JumbleSort(inputs []string, output []string) []string {

	var ints []string
	var words []string

	// Initialize output array with blank strings
	for i := range output {
		output[i] = ""
	}

	// Partition ints and words into individual arrays for sorting
	// store the type of the nth element for reference when merging
	for i, input := range inputs {
		if isInt(input) {
			ints = append(ints, input)
			output[i] = "int"
		} else if isWord(input) {
			words = append(words, strings.ToLower(input))
			output[i] = "string"
		} else {
			panic(fmt.Sprintf("WTF: Invalid input=%s", input))
		}
	}

	// Sort arrays into LIFO stacks
	sort.Sort(sort.Reverse(sort.StringSlice(ints)))
	sort.Sort(sort.Reverse(sort.StringSlice(words)))

	// Merge the arrays
	for i, value := range output {
		if value == "string" {
			output[i], words = pop(words)
		} else if value == "int" {
			output[i], ints = pop(ints)
		} else {
			panic(fmt.Sprintf("Unexpected value=%s", value))
		}
	}

	return output
}

// Converts to integers from a string input with the range -999999 to 999999 inclusive
// Returns true if conversion succeeds and false if conversion fails
func isInt(input string) bool {
	if v, err := strconv.Atoi(input); err == nil {
		if math.Abs(float64(v)) > 999999 {
			panic("Integer out of range [-999999, 999999]")
		}
		return true
	} else {
		return false
	}
}

// Uses regex to match a string containing the lower-case letters a-z
func isWord(input string) bool {
	match, _ := regexp.MatchString("^[a-z]+$", input)
	return match
}

// Pops the last element out of a slice
// Returns the popped element and the truncated slice
func pop(s []string) (string, []string) {
	return s[len(s)-1], s[:len(s)-1]
}
