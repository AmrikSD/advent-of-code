package main

import (
	amrik "2024/amrik"
	"fmt"
)

func main() {
	fmt.Println("Amrik:")
	fmt.Printf("\tDay 1 Part 1: %d\n", amrik.Day01Part01("./amrik/inputs/day01/input.txt"))
	fmt.Printf("\tDay 1 Part 2: %d\n", amrik.Day01Part02("./amrik/inputs/day01/input.txt"))
}
