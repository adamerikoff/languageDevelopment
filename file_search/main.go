package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	matches   []string
	waitGroup = sync.WaitGroup{}
	lock      = sync.Mutex{}
)

func fileSearch(root string, filename string, interactive bool) {
	if interactive {
		fmt.Println("Searching in", root)
	}
	files, _ := os.ReadDir(root)
	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			lock.Lock()
			matches = append(matches, filepath.Join(root, filename))
			lock.Unlock()
		}
		if file.IsDir() {
			waitGroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename, interactive)

		}
	}
	waitGroup.Done()
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <directory> <filename> <flags>")
		return
	}

	directory := os.Args[1]
	filename := os.Args[2]

	interactive := false

	if len(os.Args) > 3 && os.Args[3] == "i" {
		interactive = true
	}

	startTime := time.Now()

	waitGroup.Add(1)
	go fileSearch(directory, filename, interactive)
	waitGroup.Wait()

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	for _, file := range matches {
		fmt.Println("Matched:", file)
	}
	fmt.Println("Search completed in", duration.Seconds(), "seconds")
}
