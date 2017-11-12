
/*
Write a function is member() that takes a value (i.e. a number, string, etc) 
x and a array  of values a, and returns True if x is a member of a, False otherwise. 
*/
package main

import (
	"fmt"
)

func isMember(containers []int, num int) bool{
        
		for _, container := range containers {
			if container == num {
				return true
			}
   	    }
		   return false
}

func main(){

	var number int
	var container=[]int{1,2,3,4,5,10,50}
	fmt.Println("Enter number : ")
    fmt.Scan(&number)
	fmt.Printf("\nNumber present in Container %d \n%t\n", container,isMember(container,number))


  
}