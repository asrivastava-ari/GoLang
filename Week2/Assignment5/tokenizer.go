
/*
Read a string and then do the following, 
i.	Tokenize and store all tokens in a suitable Go data type, along with their frequency of occurrence. 
ii.	Display the list of tokens, along with their frequency of occurrence (sorted in decreasing order)
*/
package main

import (
	"fmt"
	"strings"
)

func tokenizer(str string){
	var chunks  []string
	var count int
	var tokenMap = make(map[string]  int)
	 
	splitString:=strings.Fields(str)
	chunks=append(chunks,splitString[0])
	for i:=1;i<len(splitString);i++{
		for j:=0; j<i; j++{
			if chunks[j] == splitString[i]{
				continue
			} else {
				chunks=append(chunks,splitString[i])
			}
		}
	}
	
	for i:=0; i<len(chunks); i++{
		count=0	
		for j:=0; j<len(splitString); j++{
			if chunks[i] == splitString[j]{
				count++
			}
		}
		tokenMap[chunks[i]]=count
	}
	
	fmt.Println(tokenMap)
	
}


func main(){
	
	str:="I am the best. I am the best. Yes you are best"
	tokenizer(str)
}

