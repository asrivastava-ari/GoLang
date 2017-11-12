/*

Given a number, compute the digital sum of the number using Go.
Definition of Digital Root:
The digital root (also repeated digital sum) of a number is the (single digit) 
value obtained by an iterative process of summing digits, on each iteration using 
the result from the previous iteration to compute a digit sum. The process continues 
until a single-digit number is reached.
Examples:
Digital root of 24566 is 2 + 4 + 5 + 6 + 6 = 23 = 2 + 3 = 5
Digital root of 65536 is 6 + 5 + 5 + 3 + 6 = 25 = 2 + 5 = 7


*/
package main

import (
	"fmt"
	)

func finalDigitalSum(x int) int {

	var result int
	result = digitalSum(x)
	
	for result > 9{
		result = digitalSum(result)
		
	}
	return result
}

func digitalSum( x int) int {
	var digit,sum int
	for x > 0 {
		digit = x%10
		sum = sum + digit
		x = x/10
	}
	return sum

}

func main() {
	var number int
	fmt.Println("Enter number : ")
    fmt.Scan(&number)
	fmt.Println(finalDigitalSum(number))
}