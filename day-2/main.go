package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	file := os.Args[1]
	values := readFile(file)

	arg1 := os.Args[2]
	fmt.Printf("arg1 = %v\n", arg1)
	iArg1, err := strconv.Atoi(arg1)
	check(err)
	values[1] = iArg1

	arg2 := os.Args[3]
	fmt.Printf("arg2 = %v\n", arg2)
	iArg2, err := strconv.Atoi(arg2)
	check(err)
	values[2] = iArg2

	for pointer := 0; pointer < len(values); pointer += 4 {

		opcode := values[pointer]

		fmt.Printf("opcode '%v'.\n", opcode)

		if opcode == 1 || opcode == 2 {
			x := values[values[pointer+1]]
			fmt.Printf("values[%v] = %v\n", pointer+1, x)
			y := values[values[pointer+2]]
			fmt.Printf("values[%v] = %v\n", pointer+2, y)
			result := 0

			if opcode == 1 {
				result = x + y
				fmt.Printf("values[%v] = %v + %v = %v\n", values[pointer+3], x, y, result)
			} else {
				result = x * y
				fmt.Printf("values[%v] = %v * %v = %v\n", values[pointer+3], x, y, result)
			}
			values[values[pointer+3]] = result
		} else if opcode == 99 {
			fmt.Println("Terminating.")
			break
		} else {
			panic(fmt.Sprintf("Invalid opcode '%v'.", opcode))
		}
	}

	fmt.Printf("values[0] = '%v'.\n", values[0])
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
	values := strings.Split(text, ",")

	fmt.Printf("%v values detected.\n", len(values))

	intValues := make([]int, len(values))

	for index, strValue := range values {

		fmt.Printf("Parsing values %v: '%v'.\n", index, strValue)

		intValue, err := strconv.Atoi(strValue)
		check(err)

		fmt.Printf("Value '%v' parsed.\n", intValue)

		intValues[index] = intValue
	}

	fmt.Printf("Returning '%v' values.\n", len(intValues))
	return intValues
}
