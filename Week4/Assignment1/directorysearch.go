/*
Write a “search_file” function in Go which traverses through all paths, 
starting from the current working directory and lists all files that matches 
that matches the search criteria. The search criteria is the name of the file. 
The program prints the absolute pathname of all the files that matches
*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

)

func main() {
	fmt.Println(checkExistence("directorysearch.go"))
}

func checkExistence(ext string) []string {
	pathS, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var files []string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(ext, f.Name())
			if err == nil && r {
				files = append(files, path)
			}
		}
		return nil
	})

	return files
}