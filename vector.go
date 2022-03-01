package main

import "math"

type Vector struct {
	x, y float64
}

func VectorFromAngle(angle float64) *Vector {
	return &Vector{
		x: math.Cos(angle),
		y: math.Sin(angle),
	}
}
