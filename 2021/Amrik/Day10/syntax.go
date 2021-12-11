package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ParseInput(filePath string) (strings []string) {

	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}

	return strings
}

func findBadBracket(str string) (rune, error) {

	stack := []rune{}

	for _, char := range str {
		switch char {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case '<':
			stack = append(stack, '>')
		default:
			if char != stack[len(stack)-1] {
				return char, nil
			}
			stack = stack[:len(stack)-1]

		}
	}
	var no rune
	return no, errors.New("no problem found")
}

func isIncomplete(str string) bool {

	stack := []rune{}

	for _, char := range str {
		switch char {
		case '(':
			stack = append(stack, ')')
		case '[':
			stack = append(stack, ']')
		case '{':
			stack = append(stack, '}')
		case '<':
			stack = append(stack, '>')
		default:
			stack = stack[:len(stack)-1]

		}
	}
	fmt.Println(str, stack)
	return false
}
