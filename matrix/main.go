package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const matrixSize = 3000

var (
	matrixOne [matrixSize][matrixSize]int
	matrixTwo [matrixSize][matrixSize]int
	result    [matrixSize][matrixSize]int
	rwLock    = sync.RWMutex{}
	cond      = sync.NewCond(&rwLock)
	waitGroup = sync.WaitGroup{}
)

// initRandomMatrix initializes the given matrix with random values
func initRandomMatrix(matrix *[matrixSize][matrixSize]int) {
	for i := 0; i < matrixSize; i++ {
		for j := 0; j < matrixSize; j++ {
			matrix[i][j] = rand.Intn(100) // Generates a random integer between 0 and 99
		}
	}
}

func workOutRow(row int) {
	for column := 0; column < matrixSize; column++ {
		for i := 0; i < matrixSize; i++ {
			result[row][column] += matrixOne[row][i] * matrixTwo[i][column]
		}
	}
	waitGroup.Done()
}

func main() {
	start := time.Now()
	fmt.Println("STARTED.")

	// Initialize matrixOne and matrixTwo with random values
	initRandomMatrix(&matrixOne)
	initRandomMatrix(&matrixTwo)

	// Start goroutines to process each row
	for row := 0; row < matrixSize; row++ {
		waitGroup.Add(1)
		go workOutRow(row)
	}

	// Wait for all goroutines to finish
	waitGroup.Wait()

	elapsed := time.Since(start)

	fmt.Printf("Processing took %s \n", elapsed)
	fmt.Println(result) // Uncomment if you want to print the result
}
