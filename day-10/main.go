package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	file := os.Args[1]
	asteroids := readFile(file)
	discoverAllSightLines(asteroids)
	newMonitoringStation := findAsteroidWithHighestVisibility(asteroids)
	fmt.Printf("New monitoring Station: {x: %v, y: %v}\n", newMonitoringStation.Location.X, newMonitoringStation.Location.Y)
	fmt.Printf("Can detect %v other asteroids.\n", len(newMonitoringStation.CanSee))

	laserLocation := newMonitoringStation.Location
	laserAngle := math.Pi / 2
	var nextTarget *Asteroid

	for i := 1; i <= 200; i++ {
		nextTarget, laserAngle = determineNextTarget(laserLocation, laserAngle, newMonitoringStation.CanSee)
		fmt.Printf("Target number %v is {x: %v, y: %v}\n", i, nextTarget.Location.X, nextTarget.Location.Y)
		asteroids = vaporizeTarget(asteroids, nextTarget)
		discoverSightLines(newMonitoringStation, asteroids)
	}
}

func readFile(file string) []*Asteroid {
	bytes, err := ioutil.ReadFile(file)
	Check(err)
	text := string(bytes)
	lines := strings.Split(text, "\n")

	var asteroids []*Asteroid

	for i := 0; i < len(lines); i++ {
		line := lines[i]

		for j := 0; j < len(line); j++ {
			char := string(line[j])

			if char == "#" {
				asteroids = append(asteroids, &Asteroid{
					Location: Point{
						X: j,
						Y: i,
					},
				})
			}
		}
	}

	return asteroids
}

func discoverAllSightLines(asteroids []*Asteroid) {
	for _, asteroid1 := range asteroids {
		discoverSightLines(asteroid1, asteroids)
	}
}

func discoverSightLines(asteroid1 *Asteroid, asteroids []*Asteroid) {
	asteroid1.CanSee = make([]*Asteroid, 0)

	for _, asteroid2 := range asteroids {

		if asteroid1 == asteroid2 {
			continue
		}

		line := Line{
			Start: asteroid1.Location,
			End:   asteroid2.Location,
		}

		canSee := true

		for _, asteroid3 := range asteroids {

			if asteroid3 == asteroid1 || asteroid3 == asteroid2 {
				continue
			}

			if line.Contains(asteroid3.Location) {
				canSee = false
				break
			}
		}

		if canSee {
			asteroid1.CanSee = append(asteroid1.CanSee, asteroid2)
		}
	}
}

func findAsteroidWithHighestVisibility(asteroids []*Asteroid) *Asteroid {
	highestVisibility := asteroids[0]

	for _, asteroid := range asteroids {
		if len(asteroid.CanSee) > len(highestVisibility.CanSee) {
			highestVisibility = asteroid
		}
	}

	return highestVisibility
}

func determineNextTarget(laserLocation Point, laserAngle float64, visibleAsteroids []*Asteroid) (*Asteroid, float64) {
	sort.Slice(visibleAsteroids, func(i, j int) bool {
		asteroid1Angle := laserLocation.Angle(visibleAsteroids[i].Location)

		if asteroid1Angle > laserAngle {
			asteroid1Angle -= 2 * math.Pi
		}

		asteroid2Angle := laserLocation.Angle(visibleAsteroids[j].Location)

		if asteroid2Angle > laserAngle {
			asteroid2Angle -= 2 * math.Pi
		}

		return asteroid1Angle > asteroid2Angle
	})

	return visibleAsteroids[0], laserLocation.Angle(visibleAsteroids[1].Location)
}

func vaporizeTarget(asteroids []*Asteroid, target *Asteroid) []*Asteroid {
	var remainingAsteroids []*Asteroid

	for _, asteroid := range asteroids {
		if asteroid != target {
			remainingAsteroids = append(remainingAsteroids, asteroid)
		}
	}

	return remainingAsteroids
}
