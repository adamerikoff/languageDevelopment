package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type Point struct {
	x int
	y int
}

var (
	r = regexp.MustCompile(`\((\d*),(\d*)\)`)
)

func findArea(inputChannel chan string) {
	for pointsString := range inputChannel {
		var points []Point
		for _, subString := range r.FindAllStringSubmatch(pointsString, -1) {
			x, _ := strconv.Atoi(subString[1])
			y, _ := strconv.Atoi(subString[2])
			points = append(points, Point{x, y})
		}

		area := 0.0
		for i := 0; i < len(points); i++ {
			a, b := points[i], points[(i+1)%len(points)]
			area += float64(a.x*b.y) - float64(a.y*b.x)
		}
		fmt.Println(math.Abs(float64(area)) / 2.0)
	}
	waitGroup.Done()
}
