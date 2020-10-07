package main

import (
	"fmt"
)

type CelestialObject struct {
	Name       string
	Orbits     *CelestialObject
	Satellites []*CelestialObject
}

func (object CelestialObject) OrbitalCheckSum() int {
	return object.orbitalCheckSum(0)
}

func (object CelestialObject) orbitalCheckSum(degreesFromCenterOfMass int) int {
	checksum := degreesFromCenterOfMass

	for _, satellite := range object.Satellites {
		checksum += satellite.orbitalCheckSum(degreesFromCenterOfMass + 1)
	}

	return checksum
}

func (object CelestialObject) ToString(indent int) string {

	var str string

	for i := 0; i < indent; i++ {
		str = fmt.Sprintf("*%v", str)
	}

	str = fmt.Sprintf("%v%v\n", str, object.Name)

	for _, satellite := range object.Satellites {
		str = fmt.Sprintf("%v%v", str, satellite.ToString(indent+1))
	}

	return str
}

func (object *CelestialObject) FindSatellite(lookingFor *CelestialObject) []*CelestialObject {
	var path []*CelestialObject

	if object == lookingFor {
		path = append(path, object)
	} else {
		for _, satellite := range object.Satellites {
			path = satellite.FindSatellite(lookingFor)

			if len(path) > 0 {
				path = append([]*CelestialObject{object}, path...)
				break
			}
		}
	}

	return path
}

func (from *CelestialObject) PathTo(to *CelestialObject) []*CelestialObject {
	var path []*CelestialObject

	for pointer := from; ; {
		satellitePath := pointer.FindSatellite(to)

		if len(satellitePath) > 0 {
			path = append(path, satellitePath...)
			break
		} else {
			path = append(path, pointer)
			pointer = pointer.Orbits
		}
	}

	return path
}
