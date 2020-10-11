package main

import (
	"fmt"
)

type Image struct {
	Height int
	Width  int
	Layers []Layer
}

func (image Image) Check() int {
	var fewestZeroLayer Layer
	fewestZeroCount := 2147483647

	fmt.Println("Checking image data...")
	fmt.Printf("Height = %v\n", image.Height)
	fmt.Printf("Width = %v\n", image.Width)
	fmt.Printf("Layers = %v\n", len(image.Layers))

	for i := 0; i < len(image.Layers); i++ {
		layer := image.Layers[i]
		zeroCount := image.Layers[i].Count(0)

		if zeroCount < fewestZeroCount {
			fewestZeroLayer = layer
			fewestZeroCount = zeroCount
		}
	}

	oneCount := fewestZeroLayer.Count(1)
	twoCount := fewestZeroLayer.Count(2)
	check := oneCount * twoCount

	fmt.Printf("Layer check = %v\n", check)

	return check
}

func (image Image) Render() Canvas {

	fmt.Println("Rendering image...")

	canvas := Canvas{
		Height: image.Height,
		Width:  image.Width,
		Pixels: make([]int, image.Width*image.Height),
	}

	for i := len(image.Layers) - 1; i >= 0; i-- {
		layer := image.Layers[i]

		for j := 0; j < len(layer.Pixels); j++ {
			if layer.Pixels[j] != 2 {
				canvas.Pixels[j] = layer.Pixels[j]
			}
		}
	}

	return canvas
}
