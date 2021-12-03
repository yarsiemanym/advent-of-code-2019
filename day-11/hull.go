package main

import "fmt"

const hullWidth = 200
const hullHeight = 100

type Hull struct {
	Panels [][]*int
}

func (hull *Hull) Init() {
	hull.Panels = make([][]*int, hullHeight)

	for row := range hull.Panels {
		hull.Panels[row] = make([]*int, hullWidth)
	}
}

func (hull *Hull) Draw(panel Point, color int) {
	hull.Panels[panel.y+(hullHeight/2)][panel.x+(hullWidth/2)] = &color
}

func (hull *Hull) Panel(panel Point) int {
	color := hull.Panels[panel.y+(hullHeight/2)][panel.x+(hullWidth/2)]

	if color == nil {
		color = new(int)
		*color = 0
	}

	return *color
}

func (hull *Hull) Print() {
	for row := range hull.Panels {
		for col := range hull.Panels[row] {
			panelCoords := Point{
				x: col - (hullWidth / 2),
				y: row - (hullHeight / 2),
			}
			if hull.Panel(panelCoords) == 0 {
				fmt.Print(".")
			} else if hull.Panel(panelCoords) == 1 {
				fmt.Print("#")
			}
		}
		fmt.Println("")
	}
	fmt.Println("")
}
