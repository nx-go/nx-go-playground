package geometry

import "nx-go-playground/math"

type Point struct {
	X float64
	Y float64
}

type Triangle struct {
	A Point
	B Point
	C Point
}

func CreatePoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func CreateTriangle(a, b, c Point) Triangle {
	return Triangle{A: a, B: b, C: c}
}

func Distance(a, b Point) float64 {
	return math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}

func (t Triangle) Area() float64 {
	a := Distance(t.A, t.B)
	b := Distance(t.B, t.C)
	c := Distance(t.C, t.A)
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

func (t Triangle) Perimeter() float64 {
	a := Distance(t.A, t.B)
	b := Distance(t.B, t.C)
	c := Distance(t.C, t.A)
	return a + b + c
}

func (t Triangle) IsRightAngled() bool {
	a := Distance(t.A, t.B)
	b := Distance(t.B, t.C)
	c := Distance(t.C, t.A)
	return math.IsRightAngledTriangle(a, b, c)
}
