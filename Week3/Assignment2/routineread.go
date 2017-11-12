
/*
Modify the program  goroutines.go to make the routines return a receive only channel to main. 
The main routine reads from this channel. Create channel in the routine
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

func palindrome(s string) chan string{

	var palindrm chan string = make(chan string)
	go func(){
		
	str :=strings.Fields(s)
		var palin string
		for i:=0; i<len(str); i++{
			if isPalindrome(str[i]){
				palin=str[i]
			} 
		}
		palindrm<-palin
	}()
	return palindrm
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
	
	func find_longest_word(str string) chan string{

		var longst chan string = make(chan string)
		go func(){
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
		longst<-finalSting
		}()
		return longst
	}


	func main(){

		var longestword chan string = make(chan string)
		var palin chan string = make(chan string)
		var testString string
		testString = " Learning     Go.It is interesting! madam"
		
		longestword =  find_longest_word(testString)
		palin =  palindrome(testString)
		
		fmt.Println("Longest word", <-longestword)
		fmt.Println("Palindrome word", <-palin)
	}