package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"sync"
)

type IntCode struct {
	Name         string
	values       []int
	WaitGroup    *sync.WaitGroup
	InputWriter  *io.PipeWriter
	InputReader  *io.PipeReader
	OutputWriter *io.PipeWriter
	OutputReader *io.PipeReader
}

func (intCode IntCode) Run(file string) {
	intCode.values = intCode.readFile(file)

	for pointer := 0; pointer < len(intCode.values); {
		intCode.executeInstruction(&pointer)
	}
}

func (intCode IntCode) readFile(path string) []int {

	bytes, err := ioutil.ReadFile(path)
	check(err)

	text := string(bytes)
	strValues := strings.Split(text, ",")

	for _, strValue := range strValues {

		intValue, err := strconv.Atoi(strValue)
		check(err)

		intCode.values = append(intCode.values, intValue)
	}

	return intCode.values
}

func (intCode IntCode) executeInstruction(pointer *int) {
	instruction := intCode.values[*pointer]
	opCode, parameterModes := intCode.parseInstruction(instruction)
	*pointer++

	if opCode == "99" {
		fmt.Printf("[%v] Terminating.\n", intCode.Name)
		*pointer = len(intCode.values)
		intCode.WaitGroup.Done()
	} else {

		switch opCode {

		case "01":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			arg1 := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 1)
			arg2 := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			sum := arg1 + arg2

			parameterMode = 1
			address := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			intCode.set(address, sum)
			fmt.Printf("[%v] @%v = %v + %v = %v\n", intCode.Name, address, arg1, arg2, sum)
			break

		case "02":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			arg1 := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 1)
			arg2 := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			product := arg1 * arg2

			parameterMode = 1
			address := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			intCode.set(address, product)
			fmt.Printf("[%v] @%v = %v * %v = %v\n", intCode.Name, address, arg1, arg2, product)
			break

		case "03":
			parameterMode := 1
			address := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++
			intCode.input(address)
			break

		case "04":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			value := intCode.get(parameterMode, intCode.values[*pointer])
			intCode.output(value)
			*pointer++
			break

		case "05":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			arg1 := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 1)
			target := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			if arg1 != 0 {
				*pointer = target
				fmt.Printf("[%v] %v != 0 => goto @%v\n", intCode.Name, arg1, target)
			}
			break

		case "06":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			arg1 := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 1)
			target := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			if arg1 == 0 {
				*pointer = target
				fmt.Printf("[%v] %v == 0 => goto @%v\n", intCode.Name, arg1, target)
			}
			break

		case "07":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			arg1 := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 1)
			arg2 := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = 1
			address := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			if arg1 < arg2 {
				intCode.set(address, 1)
				fmt.Printf("[%v] %v < %v => @%v = 1\n", intCode.Name, arg1, arg2, address)
			} else {
				intCode.set(address, 0)
				fmt.Printf("[%v] %v >= %v => @%v = 0\n", intCode.Name, arg1, arg2, address)
			}
			break

		case "08":
			parameterMode := intCode.parseParameterMode(parameterModes, 0)
			arg1 := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = intCode.parseParameterMode(parameterModes, 1)
			arg2 := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			parameterMode = 1
			address := intCode.get(parameterMode, intCode.values[*pointer])
			*pointer++

			if arg1 == arg2 {
				intCode.set(address, 1)
				fmt.Printf("[%v] %v == %v => @%v = 1\n", intCode.Name, arg1, arg2, address)
			} else {
				intCode.set(address, 0)
				fmt.Printf("[%v] %v != %v => @%v = 0\n", intCode.Name, arg1, arg2, address)
			}
			break

		default:
			panic(fmt.Sprintf("[%v] Invalid opcode '%v'.", intCode.Name, opCode))
		}
	}
}

func (intCode IntCode) parseInstruction(instruction int) (string, string) {
	strInstruction := fmt.Sprintf("%02d", instruction)
	strOpCode := string(strInstruction[len(strInstruction)-2:])
	strParameterModes := string(strInstruction[:len(strInstruction)-2])
	return strOpCode, strParameterModes
}

func (intCode IntCode) parseParameterMode(parameterModes string, index int) int {
	iParameterMode := 0
	var err error

	if len(parameterModes)-1-index >= 0 {
		iParameterMode, err = strconv.Atoi(string(parameterModes[len(parameterModes)-1-index]))
		check(err)
	}

	return iParameterMode
}

func (intCode IntCode) get(parameterMode int, parameter int) int {
	var value int

	if parameterMode == 0 {
		value = intCode.values[parameter]
	} else if parameterMode == 1 {
		value = parameter
	} else {
		panic(fmt.Sprintf("[%v] Invalid parameter mode '%v'.\n", intCode.Name, parameterMode))
	}

	return value
}

func (intCode IntCode) set(address int, value int) {
	intCode.values[address] = value
}

func (intCode IntCode) input(address int) {
	bytes := make([]byte, 4)
	_, err := intCode.InputReader.Read(bytes)
	check(err)
	value := BytesToInt(bytes)
	fmt.Printf("[%v] Input = %v\n", intCode.Name, value)
	intCode.values[address] = value

}

func (intCode IntCode) output(value int) {
	bytes := IntToBytes(value)
	fmt.Printf("[%v] Output = %v\n", intCode.Name, value)
	_, err := intCode.OutputWriter.Write(bytes)
	check(err)
}
