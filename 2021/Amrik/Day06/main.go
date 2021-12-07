package main

import (
	"fmt"
)

func main() {

	input := "input.txt"
	fmt.Println("part 1: ", partOne(input, 80))
	fmt.Println("part 2: ", partOne(input, 256))

}
func partOne(input string, days int) (answer int) {
	startingSchool := parseInput(input)
	for i := 1; i <= days; i++ {
		startingSchool = getNextDay(startingSchool)
	}

	for _, v := range startingSchool {
		answer += v
	}
	return answer
}

func partTwo(input string) (answer int) {

	return answer
}
