package main

import "fmt"

const north = 0
const east = 1
const south = 2
const west = 3

type Robot struct {
	Brain           IntCode
	Hull            Hull
	position        Point
	color           *int
	turnDirection   *int
	facingDirection int
}

func (robot *Robot) Go(input string) {
	robot.position = Point{
		x: 0,
		y: 0,
	}
	robot.facingDirection = 0
	robot.Brain.Run(input)
}

func (robot *Robot) ProcessInstruction(instruction int) {
	if robot.color == nil {
		robot.color = &instruction
		fmt.Printf("Color = %v\n", *robot.color)
	} else if robot.turnDirection == nil {
		robot.turnDirection = &instruction
		fmt.Printf("Turn Direction = %v\n", *robot.turnDirection)
	}

	if robot.color != nil && robot.turnDirection != nil {
		robot.Paint()
		robot.Turn()
		robot.Move()
		robot.color = nil
		robot.turnDirection = nil
	}
}

func (robot *Robot) Paint() {

	robot.Hull.Draw(robot.position, *robot.color)
}

func (robot *Robot) Turn() {

	fmt.Printf("Facing direction before turn = %v\n", robot.facingDirection)

	switch *robot.turnDirection {
	case 0:
		robot.facingDirection = (robot.facingDirection + 3) % 4
	case 1:
		robot.facingDirection = (robot.facingDirection + 1) % 4
	default:
		panic(fmt.Sprintf("Invalid turning direction %v.", robot.turnDirection))
	}

	fmt.Printf("Direction after turn = %v\n", robot.facingDirection)
}
func (robot *Robot) Move() {
	fmt.Printf("Position before move = %v\n", robot.position)

	switch robot.facingDirection {
	case north:
		robot.position = Point{x: robot.position.x, y: robot.position.y - 1}
	case east:
		robot.position = Point{x: robot.position.x + 1, y: robot.position.y}
	case south:
		robot.position = Point{x: robot.position.x, y: robot.position.y + 1}
	case west:
		robot.position = Point{x: robot.position.x - 1, y: robot.position.y}
	default:
		panic(fmt.Sprintf("Invalid moving direction %v.", robot.facingDirection))
	}

	fmt.Printf("Position after move = %v\n", robot.position)
}
