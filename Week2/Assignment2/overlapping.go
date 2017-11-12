/*
Define a function overlapping() that takes two arrays and 
returns True if they have at least one member in common, False otherwise.
 You may use your is_member() function, but for the sake of the exercise, 
 you should (also) write it using two nested for-loops.
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


func isMemberContainer(container1 []int, container2 []int) bool{
	
	for _, container := range container1 {
		if isMember(container2,container) {
			return true
		}
	   }
	   return false
}

func main(){

	var container1=[]int{1,2,3,4,5,11,50}
	var container2=[]int{10,20,30,40,51,11,51}
	

	fmt.Printf("\nResult:%t\n",isMemberContainer(container1,container2))


  
}