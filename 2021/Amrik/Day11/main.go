package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := "input.txt"

	fmt.Println(PartOne(input))
	fmt.Println(PartTwo(input))

}

func PartOne(input string) (flashes int) {

	octopodes := ParseInput(input)
	for i := 0; i < 100; i++ {
		flashes += step(octopodes)
	}

	printGrid(octopodes)
	return flashes

}

func PartTwo(input string) int {

	octopodes := ParseInput(input)

	currStep := 0
	flashes := 10000
	for flashes != 100 {
		flashes = step(octopodes)
		currStep++
	}

	printGrid(octopodes)
	return currStep

}
func step(input [][]int) int {
	flashed := [10][10]bool{}

	// First, the energy level of each octopus increases by 1.
	for y := range input {
		for x := range input[y] {
			input[y][x]++
		}
	}

	// Then, any octopus with an energy level greater than 9 flashes.
	flashHappened := true
	for flashHappened {
		flashHappened = false
		for y := range input {
			for x := range input[y] {
				if flashed[y][x] {
					continue
				}
				if input[y][x] > 9 {
					flashHappened = true
					flashed[y][x] = true
					neighbours := findNeighbours(input, y, x)
					for _, n := range neighbours {
						input[n[0]][n[1]]++
					}
				}
			}
		}

	}

	//count flashes and set flashed octopodes to 0
	count := 0
	for y := range flashed {
		for x := range flashed[y] {
			if flashed[y][x] {
				count++
				input[y][x] = 0
			}
		}
	}

	return count
}

func findNeighbours(input [][]int, y, x int) (out [][2]int) {
	out = [][2]int{}

	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if i < 0 || j < 0 {
				continue
			}
			if i > len(input)-1 || j > len(input[0])-1 {
				continue
			}
			out = append(out, [2]int{i, j})
		}
	}
	return out
}

func ParseInput(filePath string) [][]int {
	out := [][]int{}

	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		row := []int{}
		for _, strAsNum := range scanner.Text() {
			row = append(row, int(strAsNum-48))
		}
		out = append(out, row)
	}

	return out
}

func printGrid(input [][]int) {
	for y := range input {
		for _, v := range input[y] {
			fmt.Printf("%v", v)
		}
		fmt.Println()
	}
	fmt.Println()
}
