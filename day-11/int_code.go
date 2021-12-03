package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type OutputHandler func(value int)
type InputHandler func() int

type IntCode struct {
	Name         string
	values       []int
	relativeBase int
	OuputHandler OutputHandler
	InputHandler InputHandler
}

func (intCode *IntCode) Run(file string) {
	intCode.readFile(file)

	for pointer := 0; pointer < len(intCode.values); {
		intCode.executeInstruction(&pointer)
	}
}

func (intCode *IntCode) readFile(path string) {

	bytes, err := ioutil.ReadFile(path)
	Check(err)

	text := string(bytes)
	strValues := strings.Split(text, ",")

	for _, strValue := range strValues {

		intValue, err := strconv.Atoi(strValue)
		Check(err)

		intCode.values = append(intCode.values, intValue)
	}

	additionalMemory := make([]int, 1024)
	intCode.values = append(intCode.values, additionalMemory...)
}

func (intCode *IntCode) executeInstruction(pointer *int) {
	instruction := intCode.values[*pointer]
	opCode, parameterModes := intCode.parseInstruction(instruction)
	*pointer++

	if opCode == "99" {
		fmt.Printf("[%v] Terminating.\n", intCode.Name)
		*pointer = len(intCode.values)
	} else {

		switch opCode {

		case "01":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			arg1 := intCode.get(true, parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 1)
			arg2 := intCode.get(true, parameterMode, intCode.values[*pointer])
			*pointer++

			sum := arg1 + arg2

			parameterMode = intCode.parseParameterMode(parameterModes, 2)
			address := intCode.get(false, parameterMode, intCode.values[*pointer])
			*pointer++

			intCode.set(address, sum)
			fmt.Printf("[%v] @%v = %v + %v = %v\n", intCode.Name, address, arg1, arg2, sum)

		case "02":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			arg1 := intCode.get(true, parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 1)
			arg2 := intCode.get(true, parameterMode, intCode.values[*pointer])
			*pointer++

			product := arg1 * arg2

			parameterMode = intCode.parseParameterMode(parameterModes, 2)
			address := intCode.get(false, parameterMode, intCode.values[*pointer])
			*pointer++

			intCode.set(address, product)
			fmt.Printf("[%v] @%v = %v * %v = %v\n", intCode.Name, address, arg1, arg2, product)

		case "03":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			address := intCode.get(false, parameterMode, intCode.values[*pointer])
			*pointer++
			intCode.input(address)

		case "04":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			value := intCode.get(true, parameterMode, intCode.values[*pointer])
			intCode.output(value)
			*pointer++

		case "05":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			arg1 := intCode.get(true, parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 1)
			target := intCode.get(true, parameterMode, intCode.values[*pointer])
			*pointer++

			if arg1 != 0 {
				*pointer = target
				fmt.Printf("[%v] %v != 0 => goto @%v\n", intCode.Name, arg1, target)
			}

		case "06":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			arg1 := intCode.get(true, parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 1)
			target := intCode.get(true, parameterMode, intCode.values[*pointer])
			*pointer++

			if arg1 == 0 {
				*pointer = target
				fmt.Printf("[%v] %v == 0 => goto @%v\n", intCode.Name, arg1, target)
			}

		case "07":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			arg1 := intCode.get(true, parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 1)
			arg2 := intCode.get(true, parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 2)
			address := intCode.get(false, parameterMode, intCode.values[*pointer])
			*pointer++

			if arg1 < arg2 {
				intCode.set(address, 1)
				fmt.Printf("[%v] %v < %v => @%v = 1\n", intCode.Name, arg1, arg2, address)
			} else {
				intCode.set(address, 0)
				fmt.Printf("[%v] %v >= %v => @%v = 0\n", intCode.Name, arg1, arg2, address)
			}

		case "08":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			arg1 := intCode.get(true, parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 1)
			arg2 := intCode.get(true, parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 2)
			address := intCode.get(false, parameterMode, intCode.values[*pointer])
			*pointer++

			if arg1 == arg2 {
				intCode.set(address, 1)
				fmt.Printf("[%v] %v == %v => @%v = 1\n", intCode.Name, arg1, arg2, address)
			} else {
				intCode.set(address, 0)
				fmt.Printf("[%v] %v != %v => @%v = 0\n", intCode.Name, arg1, arg2, address)
			}

		case "09":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			arg1 := intCode.get(true, parameterMode, intCode.values[*pointer])
			*pointer++

			fmt.Printf("[%v] Relative Base = %v + %v = %v\n", intCode.Name, intCode.relativeBase, arg1, intCode.relativeBase+arg1)
			intCode.relativeBase += arg1

		default:
			panic(fmt.Sprintf("[%v] Invalid opcode '%v'.", intCode.Name, opCode))
		}
	}
}

func (intCode *IntCode) parseInstruction(instruction int) (string, string) {
	strInstruction := fmt.Sprintf("%02d", instruction)
	strOpCode := string(strInstruction[len(strInstruction)-2:])
	strParameterModes := string(strInstruction[:len(strInstruction)-2])
	return strOpCode, strParameterModes
}

func (intCode *IntCode) parseParameterMode(parameterModes string, index int) int {
	iParameterMode := 0
	var err error

	if len(parameterModes)-1-index >= 0 {
		iParameterMode, err = strconv.Atoi(string(parameterModes[len(parameterModes)-1-index]))
		Check(err)
	}

	return iParameterMode
}

func (intCode *IntCode) get(input bool, parameterMode int, parameter int) int {
	var value int

	if input {
		if parameterMode == 0 {
			value = intCode.values[parameter]
		} else if parameterMode == 1 {
			value = parameter
		} else if parameterMode == 2 {
			value = intCode.values[intCode.relativeBase+parameter]
		} else {
			panic(fmt.Sprintf("[%v] Invalid input parameter mode '%v'.\n", intCode.Name, parameterMode))
		}
	} else {
		if parameterMode == 0 {
			value = parameter
		} else if parameterMode == 2 {
			value = intCode.relativeBase + parameter
		} else {
			panic(fmt.Sprintf("[%v] Invalid output parameter mode '%v'.\n", intCode.Name, parameterMode))
		}
	}

	return value
}

func (intCode *IntCode) set(address int, value int) {
	intCode.values[address] = value
}

func (intCode IntCode) input(address int) {
	value := intCode.InputHandler()
	fmt.Printf("[%v] Input: %v\n", intCode.Name, value)
	intCode.values[address] = value
	fmt.Printf("[%v] @%v = %v\n", intCode.Name, address, value)
}

func (intCode *IntCode) output(value int) {
	fmt.Printf("[%v] Output: %v\n", intCode.Name, value)
	intCode.OuputHandler(value)
}
