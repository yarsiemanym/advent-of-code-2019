package main

type Layer struct {
	Id     int
	Pixels []int
}

func (layer Layer) Count(digit int) int {
	count := 0

	for i := 0; i < len(layer.Pixels); i++ {
		if layer.Pixels[i] == digit {
			count++
		}
	}

	return count
}
