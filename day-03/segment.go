package main

import (
	"fmt"
	"math"
)

type Segment struct {
	start Point
	end   Point
}

const epsilon = 1e-10

func (segment1 Segment) IsVertical() bool {
	return segment1.start.y != segment1.end.y && segment1.start.x == segment1.end.x
}

func (segment1 Segment) IsHorizontal() bool {
	return segment1.start.y == segment1.end.y && segment1.start.x != segment1.end.x
}

func (segment1 Segment) IsParallel(segment2 Segment) bool {
	return segment1.IsVertical() == segment2.IsVertical()
}

func (segment1 Segment) Contains(point1 Point) bool {
	a := segment1.start
	b := segment1.end
	c := point1

	divergence := math.Abs(a.Distance(b) - a.Distance(c) - c.Distance(b))

	return divergence < epsilon
}

func (segment1 Segment) Split(point1 Point) [2]Segment {
	if !segment1.Contains(point1) {
		panic("Segment does not contain point.")
	}

	segment2 := Segment{start: segment1.start, end: point1}
	segment3 := Segment{start: point1, end: segment1.end}

	return [2]Segment{segment2, segment3}
}

func (segment1 Segment) Intersection(segment2 Segment) *Point {

	fmt.Println("Determining intersection of two segments.")
	fmt.Printf("segment1 = { start: { x: %v, y: %v }, end: { x: %v, y: %v } }\n", segment1.start.x, segment1.start.y, segment1.end.x, segment1.end.y)
	fmt.Printf("segment2 = { start: { x: %v, y: %v }, end: { x: %v, y: %v } }\n", segment2.start.x, segment2.start.y, segment2.end.x, segment2.end.y)

	var intersection *Point = nil

	if segment1.IsParallel(segment2) {
		fmt.Println("Segments are parallel.")
	} else {
		var verticalSegment *Segment
		var horizontalSegment *Segment

		if segment1.IsVertical() {
			verticalSegment = &segment1
			horizontalSegment = &segment2
		} else {
			verticalSegment = &segment2
			horizontalSegment = &segment1
		}

		potentialIntersection := Point{x: verticalSegment.start.x, y: horizontalSegment.start.y}

		fmt.Printf("potentialIntersection = { x: %v, y: %v }\n", potentialIntersection.x, potentialIntersection.y)

		if segment1.Contains(potentialIntersection) && segment2.Contains(potentialIntersection) {
			intersection = &potentialIntersection
		}
	}

	if intersection == nil {
		fmt.Println("Segments do not intersect.")
	} else {
		fmt.Printf("intersection = { x: %v, y: %v }\n", intersection.x, intersection.y)
	}

	return intersection
}

func (segment1 Segment) Length() int {
	return segment1.start.ManhattanDistance(segment1.end)
}
