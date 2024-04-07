package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	matches []string
)

func fileSearch(root string, filename string) {
	fmt.Println("Searching in", root)
	files, _ := os.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			matches = append(matches, filepath.Join(root, filename))
		}
		if file.IsDir() {
			fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}
}

func main() {
	startTime := time.Now()
	fileSearch("/Users/adamerik/Documents/code/python/ml_algorithms", "voting.ipynb")
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	for _, file := range matches {
		fmt.Println("Matched:", file)
	}
	fmt.Println("Search completed in", duration.Seconds(), "seconds")
}
