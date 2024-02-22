package main

import (
	"fmt"
	"nx-go-playground/geometry"
	"nx-go-playground/logger"
)

func main() {
	triangle := geometry.CreateTriangle(geometry.CreatePoint(0, 0), geometry.CreatePoint(0, 10), geometry.CreatePoint(10, 0))
	logger.Info(fmt.Sprintf("area: %.2f", triangle.Area()))
	logger.Info(fmt.Sprintf("perimeter: %.2f", triangle.Perimeter()))
	logger.Warn(fmt.Sprintf("is right angled: %t", triangle.IsRightAngled()))
}
