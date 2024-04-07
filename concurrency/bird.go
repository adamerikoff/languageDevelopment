package main

import (
	"math/rand"
	"time"
)

type Bird struct {
	position Vector
	velocity Vector
	id       int
}

func (bird *Bird) moveOne() {
	bird.position = bird.position.Add(bird.velocity)
	next := bird.position.Add(bird.velocity)
	if next.x >= screenWidth || next.x < 0 {
		bird.velocity = Vector{x: -bird.velocity.x, y: bird.velocity.y}
	}
	if next.y >= screenHeight || next.y < 0 {
		bird.velocity = Vector{x: bird.velocity.x, y: -bird.velocity.y}
	}
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
	go bird.start()
}
