package main

import (
	"math"
	"math/rand"
	"time"
)

type Bird struct {
	position Vector
	velocity Vector
	id       int
}

func (bird *Bird) calcAcceleration() Vector {
	upper, lower := bird.position.AddConstant(viewRadius), bird.position.AddConstant(-viewRadius)
	avgVelocity := Vector{0, 0}
	count := 0.0
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			if otherBirdId := birdsMap[int(i)][int(j)]; otherBirdId != -1 && otherBirdId != bird.id {
				if dist := birds[otherBirdId].position.EuclideanDistance(bird.position); dist < viewRadius {
					count++
					avgVelocity = avgVelocity.Add(birds[otherBirdId].velocity)
				}
			}
		}
	}
	accel := Vector{0, 0}
	if count > 0 {
		avgVelocity = avgVelocity.DivideConstant(count)
		accel = avgVelocity.Substract(bird.velocity).MultiplyConstant(adjRate)
	}
	return accel
}

func (bird *Bird) moveOne() {
	bird.velocity = bird.velocity.Add(bird.calcAcceleration()).Limit(-1, 1)
	birdsMap[int(bird.position.x)][int(bird.position.y)] = -1
	bird.position = bird.position.Add(bird.velocity)
	next := bird.position.Add(bird.velocity)
	if next.x >= screenWidth || next.x < 0 {
		bird.velocity = Vector{x: -bird.velocity.x, y: bird.velocity.y}
	}
	if next.y >= screenHeight || next.y < 0 {
		bird.velocity = Vector{x: bird.velocity.x, y: -bird.velocity.y}
	}
	birdsMap[int(bird.position.x)][int(bird.position.y)] = bird.id
}

func (bird *Bird) start() {
	for {
		bird.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBird(bird_id int) {
	bird := Bird{
		position: Vector{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
		velocity: Vector{x: (rand.Float64() * 2) - 1.0, y: (rand.Float64() * 2) - 1.0},
		id:       bird_id,
	}
	birds[bird_id] = &bird
	birdsMap[int(bird.position.x)][int(bird.position.y)] = bird_id
	go bird.start()
}
