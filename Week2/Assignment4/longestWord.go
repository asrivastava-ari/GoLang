
/*
Write a function find_longest_word() that takes a list of words
and returns the length of the longest one.
*/
package main

import (
	"fmt"
	"strings"
	)

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func find_longest_word(str string) int{
	var longest int
	longest = 0
	testStrings:= strings.SplitAfter(str, " ")
	for _, testString := range testStrings {
		if len(standardizeSpaces(testString))> longest{
		 longest = len(standardizeSpaces(testString))
		}
	}
	return longest
}

func main() {
	
	var testString string
	testString = " Learning     Go.It is interesting!"
	fmt.Printf("\nTest result for longest word: %d\n",find_longest_word(testString))
}