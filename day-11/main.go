package main

import (
	"fmt"
	"os"
)

func main() {
	painted := solvePart1()
	hull := solvePart2()

	fmt.Printf("Part 1 Answer: %v\n", painted)
	fmt.Println("Part 2 Answer:")
	hull.Print()
}

func solvePart1() int {
	robot := Robot{}

	hull := Hull{}
	hull.Init()

	brain := IntCode{
		Name: "Robot Brain",
		OuputHandler: func(value int) {
			robot.ProcessInstruction(value)
		},
		InputHandler: func() int {
			return hull.Panel(robot.position)
		},
	}

	robot.Brain = brain
	robot.Hull = hull

	input := os.Args[1]
	robot.Brain.Run(input)

	painted := 0

	for _, row := range hull.Panels {
		for _, panel := range row {
			if panel != nil {
				painted++
			}
		}
	}

	return painted
}

func solvePart2() Hull {

	robot := Robot{}

	hull := Hull{}
	hull.Init()
	hull.Draw(Point{x: 0, y: 0}, 1)

	brain := IntCode{
		Name: "Robot Brain",
		OuputHandler: func(value int) {
			robot.ProcessInstruction(value)
		},
		InputHandler: func() int {
			return hull.Panel(robot.position)
		},
	}

	robot.Brain = brain
	robot.Hull = hull

	input := os.Args[1]
	robot.Brain.Run(input)

	white := 0

	for _, row := range hull.Panels {
		for _, panel := range row {
			if panel != nil {
				white++
			}
		}
	}

	return hull
}
