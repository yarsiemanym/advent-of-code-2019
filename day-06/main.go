package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	file := os.Args[1]
	objects := readFile(file)
	centerOfMass := objects["COM"]
	you := objects["YOU"]
	santa := objects["SAN"]
	orbitalMap := centerOfMass.ToString(0)

	fmt.Println(orbitalMap)

	checksum := centerOfMass.OrbitalCheckSum()

	fmt.Printf("checksum = %v\n\n", checksum)

	path := you.Orbits.PathTo(santa.Orbits)

	fmt.Print("Path = YOU-")

	for _, step := range path {
		fmt.Printf("%v-", step.Name)
	}

	fmt.Print("SAN")

	fmt.Printf("\n\nOrbital Transfers = %v\n", len(path)-1)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(path string) map[string]*CelestialObject {
	bytes, err := ioutil.ReadFile(path)
	check(err)

	text := string(bytes)
	lines := strings.Split(text, "\n")

	objects := make(map[string]*CelestialObject)

	for _, line := range lines {
		tokens := strings.Split(line, ")")
		objectName := string(tokens[0])
		satelliteName := string(tokens[1])

		object, exists := objects[objectName]

		if !exists {
			object = new(CelestialObject)
			object.Name = objectName
			objects[objectName] = object
		}

		satellite, exists := objects[satelliteName]

		if !exists {
			satellite = new(CelestialObject)
			satellite.Name = satelliteName
			objects[satelliteName] = satellite
		}

		object.Satellites = append(object.Satellites, satellite)
		satellite.Orbits = object
	}

	_, exists := objects["COM"]

	if !exists {
		panic("No center of mass")
	}

	_, exists = objects["YOU"]

	if !exists {
		panic("No you")
	}

	_, exists = objects["SAN"]

	if !exists {
		panic("No santa")
	}

	return objects
}
