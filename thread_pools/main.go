package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}

const nThreads int = 8

// 1.446760417s  without pool
// 0.591112917s with pool
func main() {
	dat, _ := os.ReadFile(filepath.Join("./", "polygons.txt"))
	text := string(dat)

	// line := "(4,10),(12,8),(10,3),(2,2),(7,5)"
	inputChannel := make(chan string, 1000)
	for i := 0; i < nThreads; i++ {
		go findArea(inputChannel)
	}
	waitGroup.Add(nThreads)
	start := time.Now()

	for _, line := range strings.Split(text, "\n") {
		inputChannel <- line
	}
	close(inputChannel)
	waitGroup.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Processing took %s \n", elapsed)

}
