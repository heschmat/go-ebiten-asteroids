package main

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v Vector) Normalize() Vector {
	mag := math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))

	return Vector{v.X/mag, v.Y/mag}
}
