package main

import (
	amrik "2024/amrik"
	nyk "2024/nyk"
	"fmt"
)

func main() {
	fmt.Println("Amrik:")
	fmt.Printf("\tDay 1 Part 1: %d\n", amrik.Day01Part01("./amrik/inputs/day01/input.txt"))
	fmt.Printf("\tDay 1 Part 2: %d\n", amrik.Day01Part02("./amrik/inputs/day01/input.txt"))
	fmt.Printf("\tDay 2 Part 1: %d\n", amrik.Day02Part01("./amrik/inputs/day02/input.txt"))
	fmt.Printf("\tDay 2 Part 2: %d\n", amrik.Day02Part02("./amrik/inputs/day02/input.txt"))

	fmt.Println("Nyk:")
	for day, solution := range nyk.Days {
		fmt.Printf("\tDay %d Part 1: %d\n", day+1, solution.Part1(solution.Input_path))
		fmt.Printf("\tDay %d Part 2: %d\n", day+1, solution.Part2(solution.Input_path))
	}

}
