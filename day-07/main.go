package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

var ampA IntCode
var ampB IntCode
var ampC IntCode
var ampD IntCode
var ampE IntCode

var waitGroup sync.WaitGroup

func main() {
	file := os.Args[1]
	min, err := strconv.Atoi(os.Args[2])
	check(err)
	max, err := strconv.Atoi(os.Args[3])
	check(err)
	maxSignal := determineMaxSignal(file, min, max)
	fmt.Printf("Max Signal = %v\n", maxSignal)
}

func determineMaxSignal(file string, min int, max int) int {
	configurations := buildTests(min, max)
	maxSignal := 0

	for _, configuration := range configurations {
		signal := testPhaseConfigration(configuration, file)

		if signal > maxSignal {
			fmt.Println("New max signal.")
			maxSignal = signal
		}
	}

	return maxSignal
}

func buildTests(min int, max int) []Configuration {
	var configrations []Configuration

	for a := min; a <= max; a++ {

		for b := min; b <= max; b++ {
			if b == a {
				continue
			}

			for c := min; c <= max; c++ {
				if c == a || c == b {
					continue
				}

				for d := min; d <= max; d++ {
					if d == a || d == b || d == c {
						continue
					}

					for e := min; e <= max; e++ {
						if e == a || e == b || e == c || e == d {
							continue
						}

						configrations = append(configrations, Configuration{
							A: a,
							B: b,
							C: c,
							D: d,
							E: e,
						})
					}
				}
			}
		}
	}

	return configrations
}

func testPhaseConfigration(configuration Configuration, file string) int {

	fmt.Printf("Testing configuration %v\n", configuration)

	amps := initialize()

	for _, amp := range amps {
		waitGroup.Add(1)
		go amp.Run(file)
	}

	amps[0].InputWriter.Write(IntToBytes(configuration.A))
	amps[1].InputWriter.Write(IntToBytes(configuration.B))
	amps[2].InputWriter.Write(IntToBytes(configuration.C))
	amps[3].InputWriter.Write(IntToBytes(configuration.D))
	amps[4].InputWriter.Write(IntToBytes(configuration.E))

	signal := 0

	go func() {
		for {

			amps[0].InputWriter.Write(IntToBytes(signal))
			bytes := make([]byte, 4)
			_, err := amps[4].OutputReader.Read(bytes)
			check(err)
			signal = BytesToInt(bytes)
		}
	}()

	waitGroup.Wait()

	fmt.Printf("Signal = %v\n", signal)

	return signal
}

func initialize() []IntCode {
	rA, wA := io.Pipe()
	rAB, wAB := io.Pipe()
	rBC, wBC := io.Pipe()
	rCD, wCD := io.Pipe()
	rDE, wDE := io.Pipe()
	rE, wE := io.Pipe()

	ampA = IntCode{
		Name:         "Amp A",
		WaitGroup:    &waitGroup,
		InputWriter:  wA,
		InputReader:  rA,
		OutputWriter: wAB,
		OutputReader: rAB,
	}

	ampB = IntCode{
		Name:         "Amp B",
		WaitGroup:    &waitGroup,
		InputWriter:  wAB,
		InputReader:  rAB,
		OutputWriter: wBC,
		OutputReader: rBC,
	}

	ampC = IntCode{
		Name:         "Amp C",
		WaitGroup:    &waitGroup,
		InputWriter:  wBC,
		InputReader:  rBC,
		OutputWriter: wCD,
		OutputReader: rCD,
	}

	ampD = IntCode{
		Name:         "Amp D",
		WaitGroup:    &waitGroup,
		InputWriter:  wCD,
		InputReader:  rCD,
		OutputWriter: wDE,
		OutputReader: rDE,
	}

	ampE = IntCode{
		Name:         "Amp E",
		WaitGroup:    &waitGroup,
		InputWriter:  wDE,
		InputReader:  rDE,
		OutputWriter: wE,
		OutputReader: rE,
	}

	return []IntCode{
		ampA,
		ampB,
		ampC,
		ampD,
		ampE,
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
