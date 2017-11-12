/*
Write a function char_freq() that takes a string and builds a frequency 
listing of the characters contained in it. Represent the frequency 
listing as a Go Map. Try it with something like 
char_freq("abbabcbdbabdbdbabababcbcbab").
*/

package main

import (
	"fmt"	
)

func main(){
	char_freq("abbabcbdbabdbdbabababcbcbabdehhh")
}

func char_freq (str string){
	var dictionary = make(map[string] int)
	x :=0
	for x < len(str){
		eachChar := string([]rune(str)[x])
		count:=dictionary[eachChar]
		count++
		dictionary[eachChar]=count
		x+=1
	}
	fmt.Println(dictionary)
}
