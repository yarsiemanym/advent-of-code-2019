package main

import (
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	file := os.Args[1]
	height, err := strconv.Atoi(os.Args[2])
	check(err)
	width, err := strconv.Atoi(os.Args[3])
	check(err)

	image := readFile(file, height, width)
	image.Check()
	canvas := image.Render()
	canvas.Check()
	canvas.Paint()
}

func readFile(file string, height int, width int) Image {
	bytes, err := ioutil.ReadFile(file)
	check(err)

	text := string(bytes)

	image := Image{
		Height: height,
		Width:  width,
	}

	for i := 0; i < len(text); i++ {

		layerIndex := i / (height * width)

		if len(image.Layers) <= layerIndex {
			image.Layers = append(image.Layers, Layer{
				Id: layerIndex + 1,
			})
		}

		digit, err := strconv.Atoi(string(text[i]))
		check(err)

		image.Layers[layerIndex].Pixels = append(image.Layers[layerIndex].Pixels, digit)
	}

	return image
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
