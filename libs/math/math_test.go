package math

import (
	"testing"
)

func TestClamp(t *testing.T) {
	if Clamp(0, 0, 1) != 0 {
		t.Error("Clamp(0, 0, 1) != 0")
	}
	if Clamp(1, 0, 1) != 1 {
		t.Error("Clamp(1, 0, 1) != 1")
	}
	if Clamp(0.5, 0, 1) != 0.5 {
		t.Error("Clamp(0.5, 0, 1) != 0.5")
	}
}

func TestSqrt(t *testing.T) {
	if Sqrt(4) != 2 {
		t.Error("Sqrt(4) != 2")
	}
	if Sqrt(9) != 3 {
		t.Error("Sqrt(9) != 3")
	}
}

func TestIsRightAngledTriangle(t *testing.T) {
	if !IsRightAngledTriangle(3, 4, 5) {
		t.Error("IsRightAngledTriangle(3, 4, 5) != true")
	}
	if IsRightAngledTriangle(3, 4, 6) {
		t.Error("IsRightAngledTriangle(3, 4, 6) != false")
	}
}
