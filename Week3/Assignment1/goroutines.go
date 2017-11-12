
/*
Implement the following through Goroutines and Channels
	a.	The main routine sends resources to be processed by different Go Routines
		and receives back the output.
	b.	Consider resources to be strings and the go routines processes the strings say one routine finds the longest token in the string (Here string is a set of words separated by whitespace(s) and another routine finds the palindrome among the words. It sends back the palindrome or “None” to the main routine.
	c.	The main routine passes the strings as function arguments to the go routines and receives the output through channels.
	d.	Here create channel in main and can be passed as argument to the go routines
*/
package main

import (
	"fmt"
	"strings"
)

func isPalindrome (str string) bool {
	
		midString:=len(str) / 2
		lastString:= len(str) - 1
		for i:= 0; i < midString; i++ {
	
			if str [i] != str[ lastString - i ]{
				return false
			}
		}
		return true
	}

func palindrome(c chan string, s string){
		
	str :=strings.Fields(s)
		var palin string
		for i:=0; i<len(str); i++{
			if isPalindrome(str[i]){
				palin=str[i]
			} 
		}
	c<-palin
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
	
	func find_longest_word(c chan string,str string) {
		var longest int
		var finalSting string
		longest = 0
		testStrings:= strings.SplitAfter(str, " ")
		for _, testString := range testStrings {
			if len(standardizeSpaces(testString))> longest{
			 longest = len(standardizeSpaces(testString))
			 finalSting = standardizeSpaces(testString)
			}
		}
		c<-finalSting
	}


	func main(){

		var longestword chan string = make(chan string)
		var palin chan string = make(chan string)
		var testString string
		testString = " Learning     Go.It is interesting! madam"
		
		go find_longest_word(longestword, testString)
		go palindrome(palin, testString)
		
		fmt.Println("Longest word", <-longestword)
		fmt.Println("Palindrome word", <-palin)
	}