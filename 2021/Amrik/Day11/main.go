package main

import (
	"bufio"
	"fmt"
	"os"
)

type Coord struct {
	y int
	x int
}

func main() {
	input := "input.txt"

	fmt.Println(input)

}

func PartOne(input string) (flashes int) {

	return flashes

}

func ProgressGrid(input [][]Octopus) (flashes int, grid [][]Octopus) {
	flashes = 0

	// First, the energy level of each octopus increase by one
	grid = incrementAll(input)

	// Then, any octopus with an energy level greater than 9 flashes.
	// This increases the energy level of all adjacent octopuses by 1,
	// including octopuses that are diagonally adjacent.
	// If this causes an octopus to have an energy level greater than 9, it also flashes.
	// This process continues as long as new octopuses keep having their energy level increased beyond 9. (
	// An octopus can only flash at most once per step.)
	for y := range grid {
		for x := range grid[y] {
			flashes, grid = flash(y, x, flashes, grid)
		}
	}

	return flashes, grid
}

func flash(y, x, flashes int, grid [][]Octopus) (int, [][]Octopus) {

	if grid[y][x].flashed {
		return flashes, grid
	}

	out := [][]Octopus{}
	for i := range grid {
		OctoRow := []Octopus{}
		for j, o := range grid[y] {
			if i == y && j == x {
				OctoRow = append(OctoRow, Octopus{o.value, true})
			} else {
				OctoRow = append(OctoRow, Octopus{o.value, o.flashed})
			}
		}
		out = append(out, OctoRow)
	}

	neighbours := findNeighbours(y, x, out)

	return flashes, grid
}

func findNeighbours(y, x int, input [][]Octopus) []Coord {
	out := []Coord{}
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if i < 0 || j < 0 {
				continue
			}
			if i > len(input)-1 || j > len(input[i])-1 {
				continue
			}
			out = append(out, Coord{i, j})
		}
	}
	return out
}

func incrementAll(input [][]Octopus) [][]Octopus {
	out := [][]Octopus{}
	for _, row := range input {
		OctoRow := []Octopus{}
		for _, o := range row {
			OctoRow = append(OctoRow, Octopus{o.value + 1, o.flashed})
		}
		out = append(out, OctoRow)
	}
	return out
}

func ParseInput(filePath string) [][]Octopus {
	out := [][]Octopus{}

	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		OctoRow := []Octopus{}
		for _, strAsNum := range scanner.Text() {
			OctoRow = append(OctoRow, Octopus{int(strAsNum - '0'), false})
		}
		out = append(out, OctoRow)
	}

	return out
}
