package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	min := os.Args[1]
	max := os.Args[2]

	passwords := findPasswords(min, max)

	fmt.Println(passwords)

	fmt.Printf("Number of passwords: %v\n", len(passwords))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findPasswords(min string, max string) []string {

	if len(min) != len(max) {
		panic("The length of min and max are not the same.")
	}

	var passwords []string

	for i := 1; i <= 9; i++ {
		root := fmt.Sprintf("%v", i)
		branches := buildPasswords(root, len(min), nil)

		for _, branch := range branches {
			if isValid(branch, min, max) {
				passwords = append(passwords, branch)
			}
		}
	}

	return passwords
}

func buildPasswords(root string, targetLength int, doubledDigit *int) []string {

	var passwords []string

	previous, err := strconv.Atoi(string(root[len(root)-1]))
	check(err)

	if targetLength-len(root) <= 0 {
		passwords = []string{root}
	} else if targetLength-len(root) == 1 && doubledDigit == nil {
		branch := fmt.Sprintf("%v%v", root, previous)
		passwords = []string{branch}
	} else {
		for i := previous; i <= 9; i++ {
			var newDoubledDigit *int

			if i == previous && doubledDigit == nil {
				newDoubledDigit = &i
			} else {
				newDoubledDigit = doubledDigit
			}

			branch := fmt.Sprintf("%v%v", root, i)
			passwords = append(passwords, buildPasswords(branch, targetLength, newDoubledDigit)...)
		}
	}

	return passwords
}

func isValid(password string, min string, max string) bool {
	return password >= min && password <= max
}
