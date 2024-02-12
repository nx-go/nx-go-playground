package math

import "math"

func Clamp(f, low, high float64) float64 {
	if f < low {
		return low
	}
	if f > high {
		return high
	}
	return f
}

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func IsRightAngledTriangle(a, b, c float64) bool {
	return a*a+b*b == c*c || a*a+c*c == b*b || b*b+c*c == a*a
}
