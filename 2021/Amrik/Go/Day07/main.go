package main

import (
	"fmt"
)

func main() {

	input := "input.txt"
	fmt.Println("part 1: ", partOne(input))
	fmt.Println("part 2: ", partTwo(input))

}
func partOne(input string) (answer int) {

	crabs := parseInput(input)

	return getDistanceForCrabsToHorizonalSpaceP1(crabs)

}

func partTwo(input string) (answer int) {

	crabs := parseInput(input)

	return getDistanceForCrabsToHorizonalSpaceP2(crabs)

}
