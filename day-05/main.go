package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var values []int

func main() {

	file := os.Args[1]
	values := readFile(file)

	for pointer := 0; pointer < len(values); {
		executeInstruction(&pointer)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readFile(path string) []int {

	bytes, err := ioutil.ReadFile(path)
	check(err)

	text := string(bytes)
	strValues := strings.Split(text, ",")

	for _, strValue := range strValues {

		intValue, err := strconv.Atoi(strValue)
		check(err)

		values = append(values, intValue)
	}

	return values
}

func executeInstruction(pointer *int) {
	instruction := values[*pointer]
	opCode, parameterModes := parseInstruction(instruction)
	*pointer++

	if opCode == "99" {
		fmt.Println("Terminating.")
		*pointer = len(values)
	} else {

		switch opCode {

		case "01":
			parameterMode := parseParameterMode(parameterModes, 0)
			arg1 := get(parameterMode, values[*pointer])
			*pointer++

			parameterMode = parseParameterMode(parameterModes, 1)
			arg2 := get(parameterMode, values[*pointer])
			*pointer++

			sum := arg1 + arg2

			parameterMode = 1
			address := get(parameterMode, values[*pointer])
			*pointer++

			set(address, sum)

			fmt.Printf("@%v = %v + %v = %v\n", address, arg1, arg2, sum)
			break

		case "02":
			parameterMode := parseParameterMode(parameterModes, 0)
			arg1 := get(parameterMode, values[*pointer])
			*pointer++

			parameterMode = parseParameterMode(parameterModes, 1)
			arg2 := get(parameterMode, values[*pointer])
			*pointer++

			product := arg1 * arg2

			parameterMode = 1
			address := get(parameterMode, values[*pointer])
			*pointer++

			set(address, product)

			fmt.Printf("@%v = %v * %v = %v\n", address, arg1, arg2, product)
			break

		case "03":
			parameterMode := 1
			address := get(parameterMode, values[*pointer])
			*pointer++

			fmt.Print("Input: ")
			reader := bufio.NewReader(os.Stdin)
			input, err := reader.ReadString('\n')
			check((err))
			input = strings.TrimSpace(input)
			value, err := strconv.Atoi(input)
			check(err)

			set(address, value)
			break

		case "04":
			parameterMode := parseParameterMode(parameterModes, 0)
			value := get(parameterMode, values[*pointer])
			fmt.Printf("TEST Result: %v\n", value)
			*pointer++
			break

		case "05":
			parameterMode := parseParameterMode(parameterModes, 0)
			arg1 := get(parameterMode, values[*pointer])
			*pointer++

			parameterMode = parseParameterMode(parameterModes, 1)
			target := get(parameterMode, values[*pointer])
			*pointer++

			if arg1 != 0 {
				*pointer = target
				fmt.Printf("%v != 0 => goto @%v\n", arg1, target)
			}
			break

		case "06":
			parameterMode := parseParameterMode(parameterModes, 0)
			arg1 := get(parameterMode, values[*pointer])
			*pointer++

			parameterMode = parseParameterMode(parameterModes, 1)
			target := get(parameterMode, values[*pointer])
			*pointer++

			if arg1 == 0 {
				*pointer = target
				fmt.Printf("%v == 0 => goto @%v\n", arg1, target)
			}
			break

		case "07":
			parameterMode := parseParameterMode(parameterModes, 0)
			arg1 := get(parameterMode, values[*pointer])
			*pointer++

			parameterMode = parseParameterMode(parameterModes, 1)
			arg2 := get(parameterMode, values[*pointer])
			*pointer++

			parameterMode = 1
			address := get(parameterMode, values[*pointer])
			*pointer++

			if arg1 < arg2 {
				set(address, 1)
				fmt.Printf("%v < %v => @%v = 1\n", arg1, arg2, address)
			} else {
				set(address, 0)
				fmt.Printf("%v >= %v => @%v = 0\n", arg1, arg2, address)
			}
			break

		case "08":
			parameterMode := parseParameterMode(parameterModes, 0)
			arg1 := get(parameterMode, values[*pointer])
			*pointer++

			parameterMode = parseParameterMode(parameterModes, 1)
			arg2 := get(parameterMode, values[*pointer])
			*pointer++

			parameterMode = 1
			address := get(parameterMode, values[*pointer])
			*pointer++

			if arg1 == arg2 {
				set(address, 1)
				fmt.Printf("%v == %v => @%v = 1\n", arg1, arg2, address)
			} else {
				set(address, 0)
				fmt.Printf("%v != %v => @%v = 0\n", arg1, arg2, address)
			}

			break

		default:
			panic(fmt.Sprintf("Invalid opcode '%v'.", opCode))
		}
	}
}

func parseInstruction(instruction int) (string, string) {
	strInstruction := fmt.Sprintf("%02d", instruction)
	strOpCode := string(strInstruction[len(strInstruction)-2:])
	strParameterModes := string(strInstruction[:len(strInstruction)-2])
	return strOpCode, strParameterModes
}

func parseParameterMode(parameterModes string, index int) int {
	iParameterMode := 0
	var err error

	if len(parameterModes)-1-index >= 0 {
		iParameterMode, err = strconv.Atoi(string(parameterModes[len(parameterModes)-1-index]))
		check(err)
	}

	return iParameterMode
}

func get(parameterMode int, parameter int) int {
	var value int

	if parameterMode == 0 {
		value = values[parameter]
	} else if parameterMode == 1 {
		value = parameter
	} else {
		panic(fmt.Sprintf("Invalid parameter mode '%v'.\n", parameterMode))
	}

	return value
}

func set(address int, value int) {
	values[address] = value
}

func print(address int) {
	fmt.Printf("values[%v] = %v", address, values[address])
}
