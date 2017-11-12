/*
Write a Go program to compute square root of a number using Newton’s method
 (https://math.mit.edu/~stevenj/18.335/newton-sqrt.pdf)
*/


package main

import (
	"fmt"
	"math"
)

var p = 0.0

func Sqrt(x float64) float64 {
	z := 1.0
	for math.Abs(p - z) > 0.000001 { 
		z = newton(z, x)
		p = newton(z, x)
	}
	return z
}

func newton(z, x float64) float64 {
	return z - ( ((z*z) - x) / (2*z) )
}

func main() {
	var number float64
	fmt.Println("Enter number : ")
    fmt.Scan(&number)
	fmt.Println(Sqrt(number))
}