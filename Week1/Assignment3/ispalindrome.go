/*
Define a function is_palindrome() that recognizes palindromes 
(i.e. words that look the same written backwards). 
For example, is_palindrome("radar") should return True
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

func main () {

	var inputStr string
	fmt.Println("\nEnter String for test:")
	fmt.Scanf("%s\n", &inputStr)
	inputStr = strings.ToLower(inputStr)
	fmt.Printf("\nTest result for palindrome: %t\n",isPalindrome(inputStr))
}