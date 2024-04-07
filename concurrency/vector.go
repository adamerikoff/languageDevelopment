package main

import "math"

type Vector struct {
	x float64
	y float64
}

func (v1 Vector) Add(v2 Vector) Vector {
	return Vector{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
	}
}

func (v1 Vector) Substract(v2 Vector) Vector {
	return Vector{
		x: v1.x - v2.x,
		y: v1.y - v2.y,
	}
}

func (v1 Vector) Multiply(v2 Vector) Vector {
	return Vector{
		x: v1.x * v2.x,
		y: v1.y * v2.y,
	}
}

func (v1 Vector) AddConstant(d float64) Vector {
	return Vector{
		x: v1.x + d,
		y: v1.y + d,
	}
}

func (v1 Vector) SubstractConstant(d float64) Vector {
	return Vector{
		x: v1.x - d,
		y: v1.y - d,
	}
}

func (v1 Vector) MultiplyConstant(d float64) Vector {
	return Vector{
		x: v1.x * d,
		y: v1.y * d,
	}
}

func (v1 Vector) DivideConstant(d float64) Vector {
	return Vector{
		x: v1.x / d,
		y: v1.y / d,
	}
}

func (v1 Vector) Limit(lower, higher float64) Vector {
	return Vector{
		x: math.Min(math.Max(v1.x, lower), higher),
		y: math.Min(math.Max(v1.y, lower), higher),
	}
}

func (v1 Vector) EuclideanDistance(v2 Vector) float64 {
	return math.Sqrt(math.Pow(v1.x-v2.x, 2) + math.Pow(v1.y-v2.y, 2))
}
