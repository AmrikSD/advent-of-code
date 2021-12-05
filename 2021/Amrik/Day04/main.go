package main

import (
	"fmt"
	"reflect"
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

	boards, draws := ParseInput(input)

	winners := []Board{}

	for _, v := range draws {
		for _, b := range boards {
			b.processDraw(v)
			if b.isBingo() {
				if contains(winners, b) {
					continue
				}
				winners = append(winners, b)
				if len(winners) == len(boards) {
					return b.sumOfNotVisited() * v
				}
			}
		}
	}
	return 0
}

func contains(boards []Board, board Board) bool {
	for _, v := range boards {
		if reflect.DeepEqual(v, board) {
			return true
		}
	}
	return false
}
