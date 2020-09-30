package main

import "fmt"

type Wire struct {
	segments []Segment
}

func (wire1 Wire) Intersections(wire2 Wire) []Point {
	fmt.Println("Determining intersection of two wires.")
	var intersections []Point

	for _, segment1 := range wire1.segments {
		for _, segment2 := range wire2.segments {
			intersection := segment1.Intersection(segment2)

			if intersection != nil {
				intersections = append(intersections, *intersection)
			}
		}
	}

	fmt.Printf("%v intersections found.\n", len(intersections))

	return intersections
}

func (wire1 Wire) StepsTo(point1 Point) int {
	steps := 0

	for i := 0; i < len(wire1.segments); i++ {

		segment := wire1.segments[i]

		if segment.Contains(point1) {
			steps += segment.Split(point1)[0].Length()
			break
		} else {
			steps += segment.Length()
		}
	}

	return steps
}
