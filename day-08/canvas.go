package main

import (
	"fmt"

	"github.com/gookit/color"
)

type Canvas struct {
	Height int
	Width  int
	Pixels []int
}

func (canvas Canvas) Check() {
	fmt.Println("Checking canvas data...")
	fmt.Printf("Height = %v\n", canvas.Height)
	fmt.Printf("Width = %v\n", canvas.Width)
	fmt.Printf("Pixels = %v\n", len(canvas.Pixels))
}

func (canvas Canvas) Paint() {

	fmt.Println("Painting canvas...")

	for i := 0; i < len(canvas.Pixels); i++ {
		switch canvas.Pixels[i] {

		case 0:
			color.BgBlack.Print(" ")
			break

		case 1:
			color.BgWhite.Print(" ")
			break

		default:
			break
		}

		if i%canvas.Width == canvas.Width-1 {
			fmt.Println("")
		}
	}
}
