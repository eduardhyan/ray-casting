package objects

import "math"

type Vector struct {
	X, Y float64
}

func VectorFromAngle(angle float64) *Vector {
	return &Vector{
		X: math.Cos(angle),
		Y: math.Sin(angle),
	}
}
