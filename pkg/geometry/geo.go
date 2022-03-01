package geometry

import "math"

func Radian(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

func Distance(x1, y1, x2, y2 float64) float64 {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy)
}
