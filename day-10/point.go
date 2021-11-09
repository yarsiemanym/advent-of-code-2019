package main

import (
	"math"
)

type Point struct {
	X int
	Y int
}

func (point1 Point) Distance(point2 Point) float64 {
	deltaX := float64(point1.X - point2.X)
	deltaY := float64(point1.Y - point2.Y)
	distance := math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	return distance
}

func (point1 Point) ManhattanDistance(point2 Point) int {
	deltaX := math.Abs(float64(point1.X - point2.X))
	deltaY := math.Abs(float64(point1.Y - point2.Y))
	distance := int(deltaX + deltaY)
	return distance
}

func (point1 Point) Add(point2 Point) Point {
	return Point{X: point1.X + point2.X, Y: point1.Y + point2.Y}
}

func (point1 Point) Subtract(point2 Point) Point {
	return Point{X: point1.X - point2.X, Y: point1.Y - point2.Y}
}

func (point1 Point) Scale(magnitude int) Point {
	return Point{X: point1.X * magnitude, Y: point1.Y * magnitude}
}

func (point1 Point) Dot(point2 Point) int {
	return point1.X*point2.X + point1.Y*point2.Y
}

func (point1 Point) Cross(point2 Point) int {
	return point1.X*point2.X - point1.Y*point2.Y
}

func (point1 Point) Angle(point2 Point) float64 {
	slope := point2.Subtract(point1)
	angle := math.Atan2(float64(0-slope.Y), float64(slope.X))
	return angle
}
