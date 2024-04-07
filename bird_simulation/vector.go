package main

import "math"

// Vector represents a 2D vector with x and y components.
type Vector struct {
	x float64
	y float64
}

// Add returns the result of adding two vectors.
func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
	}
}

// Subtract returns the result of subtracting v2 from v1.
func (v1 Vector) Subtract(v2 Vector) Vector {
	return Vector{
		x: v1.x - v2.x,
		y: v1.y - v2.y,
	}
}

// Multiply returns the result of element-wise multiplication of two vectors.
func (v1 Vector) Multiply(v2 Vector) Vector {
	return Vector{
		x: v1.x * v2.x,
		y: v1.y * v2.y,
	}
}

// AddConstant returns the result of adding a constant value to each component of the vector.
func (v1 Vector) AddConstant(d float64) Vector {
	return Vector{
		x: v1.x + d,
		y: v1.y + d,
	}
}

// SubtractConstant returns the result of subtracting a constant value from each component of the vector.
func (v1 Vector) SubtractConstant(d float64) Vector {
	return Vector{
		x: v1.x - d,
		y: v1.y - d,
	}
}

// MultiplyConstant returns the result of multiplying each component of the vector by a constant value.
func (v1 Vector) MultiplyConstant(d float64) Vector {
	return Vector{
		x: v1.x * d,
		y: v1.y * d,
	}
}

// DivideConstant returns the result of dividing each component of the vector by a constant value.
func (v1 Vector) DivideConstant(d float64) Vector {
	return Vector{
		x: v1.x / d,
		y: v1.y / d,
	}
}

// Limit clamps each component of the vector to be within the range [lower, higher].
func (v1 Vector) Limit(lower, higher float64) Vector {
	return Vector{
		x: math.Min(math.Max(v1.x, lower), higher),
		y: math.Min(math.Max(v1.y, lower), higher),
	}
}

// EuclideanDistance calculates the Euclidean distance between two vectors.
func (v1 Vector) EuclideanDistance(v2 Vector) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
