package main

import "math"

type Point struct {
	x int
	y int
}

func (point1 Point) Distance(point2 Point) float64 {
	deltaX := float64(point1.x - point2.x)
	deltaY := float64(point1.y - point2.y)
	distance := math.Sqrt(math.Pow(deltaX, 2) + math.Pow(deltaY, 2))
	return distance
}

func (point1 Point) ManhattanDistance(point2 Point) int {
	deltaX := math.Abs(float64(point1.x - point2.x))
	deltaY := math.Abs(float64(point1.y - point2.y))
	distance := int(deltaX + deltaY)
	return distance
}

func (point1 Point) Add(point2 Point) Point {
	return Point{x: point1.x + point2.x, y: point1.y + point2.y}
}

func (point1 Point) Subtract(point2 Point) Point {
	return Point{x: point1.x - point2.x, y: point1.y - point2.y}
}

func (point1 Point) Scale(magnitude int) Point {
	return Point{x: point1.x * magnitude, y: point1.y * magnitude}
}

func (point1 Point) Dot(point2 Point) int {
	return point1.x*point2.x + point1.y*point2.y
}

func (point1 Point) Cross(point2 Point) int {
	return point1.x*point2.x - point1.y*point2.y
}
