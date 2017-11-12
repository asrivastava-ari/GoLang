
/*
Define a simple "spelling correction" function correct() that 
takes a string and sees to it that 1) two or more occurrences 
of the space character is compressed into one, and 2) inserts 
an extra space after a period if the period is directly followed 
by a letter. E.g. correct("Learning     Go.It is interesting!") 
should return " Learning Go. It is interesting!".
*/

package main

import (
    "fmt"
	"strings"
	)

func standardizeSpaces(s string) string {
    return strings.Join(strings.Fields(s), " ")
}

func main() {
    testStrings := []string{" Learning     Go.It is interesting!"}
    for _, testString := range testStrings {
		fmt.Println(standardizeSpaces(testString))
		fmt.Println(strings.Replace(standardizeSpaces(testString),".",". ",-1))
    }
}