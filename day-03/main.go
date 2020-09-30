package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var centralPort = Point{x: 0, y: 0}

func main() {
	file := os.Args[1]
	wires := readFile(file)

	var nearestIntersection *Point
	var fewestSteps = 0
	wire1 := wires[0]
	wire2 := wires[1]
	intersections := wire1.Intersections(wire2)

	for i := 0; i < len(intersections); i++ {

		steps := wire1.StepsTo(intersections[i]) + wire2.StepsTo(intersections[i])

		if centralPort != intersections[i] && (nearestIntersection == nil || steps < fewestSteps) {
			nearestIntersection = &intersections[i]
			fewestSteps = steps
		}
	}

	if nearestIntersection == nil {
		fmt.Println("No wires intersect.")
	} else {
		fmt.Printf("Nearest intersection = {x: %v, y: %v}\n", nearestIntersection.x, nearestIntersection.y)
		fmt.Printf("Distance from central port = %v\n", nearestIntersection.ManhattanDistance(centralPort))
		fmt.Printf("Steps from central port = %v\n", fewestSteps)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(path string) []Wire {

	fmt.Printf("Reading file '%v'.\n", path)

	bytes, err := ioutil.ReadFile(path)
	check(err)

	fmt.Printf("%v bytes read.\n", len(bytes))

	text := string(bytes)
	lines := strings.Split(text, "\n")
	fmt.Printf("%v lines read.\n", len(lines))

	wires := make([]Wire, len(lines))

	for index, line := range lines {
		wires[index] = parseWire(line)
	}

	fmt.Printf("%v wires constructed.\n", len(wires))

	return wires
}

func parseWire(line string) Wire {
	fmt.Println("Parsing wire.")

	instructions := strings.Split(line, ",")
	wire := Wire{segments: make([]Segment, len(instructions))}

	fmt.Printf("%v instructions detected.\n", len(instructions))

	for index, instruction := range instructions {

		var start Point

		if index == 0 {
			start = centralPort
		} else {
			start = wire.segments[index-1].end
		}

		segment := parseSegment(start, instruction)
		wire.segments[index] = segment
	}

	return wire
}

func parseSegment(start Point, instruction string) Segment {
	fmt.Println("Parsing segment.")
	fmt.Printf("start = { x: %v, y: %v}\n", start.x, start.y)
	fmt.Printf("instruction = '%v'\n", instruction)

	direction := string(instruction[0])
	fmt.Printf("direction = '%v'\n", direction)

	distance, err := strconv.Atoi(string(instruction[1:]))
	check(err)
	fmt.Printf("distance = '%v'\n", distance)

	end := Point{x: start.x, y: start.y}

	if direction == "U" {
		end.y += distance
	} else if direction == "D" {
		end.y -= distance
	} else if direction == "L" {
		end.x -= distance
	} else if direction == "R" {
		end.x += distance
	} else {
		panic(fmt.Sprintf("Invalid direction '%v'.", direction))
	}

	fmt.Printf("end = { x: %v, y: %v}\n", end.x, end.y)

	return Segment{start: start, end: end}
}
