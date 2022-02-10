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

	lines := RemoveNonHorizontalOrVericalLines(GetLines(input))
	grid := CreateGrid(lines)
	sum := 0
	for _, row := range grid {
		for _, v := range row {
			if v >= 2 {
				sum++
			}
		}
	}

	// fmt.Println()
	// fmt.Println()
	// for _, row := range grid {
	// 	for _, c := range row {
	// 		fmt.Printf("%v", c)
	// 	}
	// 	fmt.Println()
	// }

	return sum
}

func partTwo(input string) (answer int) {
	lines := GetLines(input)
	grid := CreateGrid(lines)
	sum := 0
	for _, row := range grid {
		for _, v := range row {
			if v >= 2 {
				sum++
			}
		}
	}

	// fmt.Println()
	// fmt.Println()
	// for _, row := range grid {
	// 	for _, c := range row {
	// 		fmt.Printf("%v", c)
	// 	}
	// 	fmt.Println()
	// }

	return sum
}
