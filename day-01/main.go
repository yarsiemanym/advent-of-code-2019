package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {

	masses := readFile("./input.txt")
	totalFuel := 0

	fmt.Printf("Total fuel '%v'.'\n", totalFuel)

	for index, mass := range masses {

		fmt.Printf("Calculating fuel for module '%v'.\n", index)
		fuel := calcFuel(mass)
		fmt.Printf("'%v' fuel required for module %v.\n", fuel, index)
		totalFuel += fuel

		fmt.Printf("Total fuel '%v'.\n", totalFuel)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(path string) []int {

	fmt.Printf("Reading file '%v'.\n", path)

	bytes, err := ioutil.ReadFile(path)
	check(err)

	fmt.Printf("%v bytes read.\n", len(bytes))

	text := string(bytes)
	lines := strings.Split(text, "\n")

	fmt.Printf("%v lines detected.\n", len(lines))

	intValues := make([]int, len(lines))

	for index, strValue := range lines {

		fmt.Printf("Parsing line %v: '%v'.\n", index, strValue)

		intValue, err := strconv.Atoi(strValue)
		check(err)

		fmt.Printf("Value '%v' parsed.\n", intValue)

		intValues[index] = intValue
	}

	fmt.Printf("Returning '%v' values.\n", len(intValues))
	return intValues
}

func calcFuel(mass int) int {

	if mass <= 0 {
		return 0
	}

	fmt.Printf("Calculating fuel required for module of mass '%v'.\n", mass)
	moduleFuel := int(math.Max(float64(mass)/3-2, 0))
	fmt.Printf("'%v' fuel required.\n", moduleFuel)
	fuelFuel := calcFuel(moduleFuel)
	totalFuel := moduleFuel + fuelFuel

	return totalFuel
}
