/*
Write a variadic function in Go, which concatenates an arbitrary number of strings.
(References: https://www.golang-book.com/books/intro/7
https://en.wikipedia.org/wiki/Variadic_function
*/

package main

import "fmt"

func concatenate(str ...string) {
	var finalStr string
    for _, s := range str {
        finalStr += s
    }
    fmt.Printf("Final concatenated string:\n%s\n",finalStr)
}

func main() {

    strs := []string{"AC", "A", "BBB", "CCC"}
    concatenate(strs...)
}
