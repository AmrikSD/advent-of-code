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

	boards, draws := ParseInput(input)

	for _, v := range draws {
		for _, b := range boards {
			b.processDraw(v)
			if b.isBingo() {
				return b.sumOfNotVisited() * v
			}
		}
	}
	return 0
}

func partTwo(input string) (answer int) {

	// file, _ := os.Open(input)
	// scanner := bufio.NewScanner(file)

	return 0
}
