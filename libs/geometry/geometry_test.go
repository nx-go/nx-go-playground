package geometry

import "testing"

func TestCreatePoint(t *testing.T) {
	p := CreatePoint(0, 0)
	if p.X != 0 {
		t.Errorf("X should be 0, got %f", p.X)
	}
	if p.Y != 0 {
		t.Errorf("Y should be 0, got %f", p.Y)
	}
}

func TestDistance(t *testing.T) {
	p1 := CreatePoint(0, 0)
	p2 := CreatePoint(0, 10)
	if Distance(p1, p2) != 10 {
		t.Errorf("Distance should be 10, got %f", Distance(p1, p2))
	}
}

func TestTriangleArea(t *testing.T) {
	p1 := CreatePoint(0, 0)
	p2 := CreatePoint(0, 10)
	p3 := CreatePoint(10, 0)
	t1 := CreateTriangle(p1, p2, p3)
	if t1.Area() != 50 {
		t.Errorf("Area should be 50, got %f", t1.Area())
	}
}

func TestTrianglePerimeter(t *testing.T) {
	p1 := CreatePoint(-2, 6)
	p2 := CreatePoint(2, 4)
	p3 := CreatePoint(-1, 0)
	t1 := CreateTriangle(p1, p2, p3)
	if t1.Perimeter()-15.55 <= 1e-5 {
		t.Errorf("Perimeter should be 15.55, got %f", t1.Perimeter())
	}
}

func TestTriangleIsRightAngled(t *testing.T) {
	p1 := CreatePoint(0, 0)
	p2 := CreatePoint(0, 10)
	p3 := CreatePoint(10, 0)
	t1 := CreateTriangle(p1, p2, p3)
	if t1.IsRightAngled() {
		t.Errorf("IsRightAngled should be false, got %t", t1.IsRightAngled())
	}
}
