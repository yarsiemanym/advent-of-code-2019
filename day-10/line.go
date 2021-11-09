package main

import (
	"fmt"
	"math"
)

type Line struct {
	Start Point
	End   Point
}

const epsilon = 1e-10

func (line1 Line) IsVertical() bool {
	return line1.Start.Y != line1.End.Y && line1.Start.X == line1.End.X
}

func (line1 Line) IsHorizontal() bool {
	return line1.Start.Y == line1.End.Y && line1.Start.X != line1.End.X
}

func (line1 Line) IsParallel(line2 Line) bool {
	return line1.IsVertical() == line2.IsVertical()
}

func (line1 Line) Contains(point1 Point) bool {
	a := line1.Start
	b := line1.End
	c := point1

	divergence := math.Abs(a.Distance(b) - a.Distance(c) - c.Distance(b))

	return divergence < epsilon
}

func (line1 Line) Split(point1 Point) [2]Line {
	if !line1.Contains(point1) {
		panic("Line does not contain point.")
	}

	line2 := Line{Start: line1.Start, End: point1}
	line3 := Line{Start: point1, End: line1.End}

	return [2]Line{line2, line3}
}

func (line1 Line) Intersection(line2 Line) *Point {

	fmt.Println("Determining intersection of two lines.")
	fmt.Printf("line1 = { Start: { x: %v, y: %v }, End: { x: %v, y: %v } }\n", line1.Start.X, line1.Start.Y, line1.End.X, line1.End.Y)
	fmt.Printf("line2 = { Start: { x: %v, y: %v }, End: { x: %v, y: %v } }\n", line2.Start.X, line2.Start.Y, line2.End.X, line2.End.Y)

	var intersection *Point = nil

	if line1.IsParallel(line2) {
		fmt.Println("Lines are parallel.")
	} else {
		var verticalLine *Line
		var horizontalLine *Line

		if line1.IsVertical() {
			verticalLine = &line1
			horizontalLine = &line2
		} else {
			verticalLine = &line2
			horizontalLine = &line1
		}

		potentialIntersection := Point{X: verticalLine.Start.X, Y: horizontalLine.Start.Y}

		fmt.Printf("potentialIntersection = { x: %v, y: %v }\n", potentialIntersection.X, potentialIntersection.Y)

		if line1.Contains(potentialIntersection) && line2.Contains(potentialIntersection) {
			intersection = &potentialIntersection
		}
	}

	if intersection == nil {
		fmt.Println("Lines do not intersect.")
	} else {
		fmt.Printf("intersection = { x: %v, y: %v }\n", intersection.X, intersection.Y)
	}

	return intersection
}

func (line1 Line) Length() int {
	return line1.Start.ManhattanDistance(line1.End)
}
