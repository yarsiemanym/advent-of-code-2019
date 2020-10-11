package main

import "os"

func main() {
	file := os.Args[1]
	intCode := IntCode{
		Name: "BOOST",
	}

	intCode.Run(file)
}
