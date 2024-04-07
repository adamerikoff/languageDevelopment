package main

import (
	"math"
	"math/rand"
	"time"
)

// Bird represents a bird in the simulation.
type Bird struct {
	position Vector // Position of the bird.
	velocity Vector // Velocity of the bird.
	id       int    // Unique identifier for the bird.
}

// calcAcceleration calculates the acceleration of the bird based on nearby birds.
func (bird *Bird) calcAcceleration() Vector {
	// Calculate the view area of the bird.
	upper, lower := bird.position.AddConstant(viewRadius), bird.position.AddConstant(-viewRadius)
	// Initialize variables for calculating average velocity and count of nearby birds.
	avgVelocity := Vector{0, 0}
	count := 0.0
	// Iterate through the nearby area to find other birds.
	for i := math.Max(lower.x, 0); i <= math.Min(upper.x, screenWidth); i++ {
		for j := math.Max(lower.y, 0); j <= math.Min(upper.y, screenHeight); j++ {
			// Check if there is another bird at the current position.
			if otherBirdId := birdsMap[int(i)][int(j)]; otherBirdId != -1 && otherBirdId != bird.id {
				// Calculate the distance between the current bird and the other bird.
				if dist := birds[otherBirdId].position.EuclideanDistance(bird.position); dist < viewRadius {
					// If the other bird is within view radius, add its velocity to the average velocity.
					count++
					avgVelocity = avgVelocity.Add(birds[otherBirdId].velocity)
				}
			}
		}
	}
	// Calculate acceleration based on the average velocity of nearby birds.
	accel := Vector{0, 0}
	if count > 0 {
		avgVelocity = avgVelocity.DivideConstant(count)
		accel = avgVelocity.Subtract(bird.velocity).MultiplyConstant(adjRate)
	}
	return accel
}

// moveOne moves the bird by updating its velocity and position.
func (bird *Bird) moveOne() {
	// Update velocity based on acceleration.
	bird.velocity = bird.velocity.Add(bird.calcAcceleration()).Limit(-1, 1)
	// Clear the bird's previous position on the map.
	birdsMap[int(bird.position.x)][int(bird.position.y)] = -1
	// Update the bird's position.
	bird.position = bird.position.Add(bird.velocity)
	// Check for screen boundaries and update velocity accordingly.
	next := bird.position.Add(bird.velocity)
	if next.x >= screenWidth || next.x < 0 {
		bird.velocity = Vector{x: -bird.velocity.x, y: bird.velocity.y}
	}
	if next.y >= screenHeight || next.y < 0 {
		bird.velocity = Vector{x: bird.velocity.x, y: -bird.velocity.y}
	}
	// Update the bird's position on the map.
	birdsMap[int(bird.position.x)][int(bird.position.y)] = bird.id
}

// start starts the bird's movement simulation.
func (bird *Bird) start() {
	for {
		// Move the bird and wait for a short interval.
		bird.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

// createBird creates a new bird with the given ID and starts its movement simulation.
func createBird(bird_id int) {
	// Create a new bird with random initial position and velocity.
	bird := Bird{
		position: Vector{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
		velocity: Vector{x: (rand.Float64() * 2) - 1.0, y: (rand.Float64() * 2) - 1.0},
		id:       bird_id,
	}
	// Add the bird to the global birds map and update its position on the map.
	birds[bird_id] = &bird
	birdsMap[int(bird.position.x)][int(bird.position.y)] = bird_id
	// Start the bird's movement simulation in a separate goroutine.
	go bird.start()
}
