package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := "input.txt"
	fmt.Println("part 1: ", partOne(input))

}

func partOne(input string) (answer int) {

	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)

	var length int
	used := make(map[int]int)

	for scanner.Scan() {
		length++
		number := scanner.Text()

		for pos, char := range number {
			used[pos] += int(char - '0')
		}

	}

	gamma := ""
	epsilon := ""

	for i := 0; i < len(used); i++ {
		if used[i] >= length/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	return int(gammaInt * epsilonInt)

}

func partTwo(input string) (answer int) {

	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		in := scanner.Text()
		fmt.Println(in)
	}

	return 0

}
