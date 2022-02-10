package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	x1 int
	y1 int

	x2 int
	y2 int
}

// Parses an input file returns a slice of lines
func GetLines(filePath string) (lines []Line) {
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		line := strings.Split(text, " -> ")

		startPoint := strings.Split(line[0], ",")
		endPoint := strings.Split(line[1], ",")

		x1, _ := strconv.Atoi(startPoint[0])
		y1, _ := strconv.Atoi(startPoint[1])
		x2, _ := strconv.Atoi(endPoint[0])
		y2, _ := strconv.Atoi(endPoint[1])

		lines = append(lines, Line{
			x1: x1,
			y1: y1,
			x2: x2,
			y2: y2,
		})

	}
	return lines
}

func CreateGrid(lines []Line) (grid [][]int) {

	width := 0
	height := 0

	for _, l := range lines {
		width = max(width, l.x1)
		width = max(width, l.x2)
		height = max(height, l.y1)
		height = max(height, l.y1)
	}
	for row := 0; row <= height; row++ {
		CurrRow := []int{}
		for col := 0; col <= width; col++ {
			CurrRow = append(CurrRow, 0)
		}
		grid = append(grid, CurrRow)
	}

	// convert each line to a an x,y coord to increment.
	for _, l := range lines {
		grid = AddLineToGrid(grid, l)
	}

	return grid
}

func AddLineToGrid(oldGrid [][]int, line Line) [][]int {

	newGrid := oldGrid

	if line.x1 == line.x2 {
		//Veritcal Line
		x := line.x1
		start := min(line.y1, line.y2)
		end := max(line.y1, line.y2)
		for i := start; i <= end; i++ {
			newGrid[i][x] = newGrid[i][x] + 1
		}
	} else if line.y1 == line.y2 {
		//Horizontal Line
		y := line.y1
		start := min(line.x1, line.x2)
		end := max(line.x1, line.x2)
		for i := start; i <= end; i++ {
			newGrid[y][i] = newGrid[y][i] + 1
		}
	} else {
		// do diagonal
		if line.y2 < line.y1 {
			tx := line.x1
			ty := line.y1
			line.x1 = line.x2
			line.y1 = line.y2
			line.x2 = tx
			line.y2 = ty
		}

		if line.x1 < line.x2 {
			y := line.y1
			for x := line.x1; x <= line.x2; x++ {
				newGrid[y][x] = newGrid[y][x] + 1
				y++
			}
		} else {

			y := line.y1
			for x := line.x1; x >= line.x2; x-- {
				newGrid[y][x] = newGrid[y][x] + 1
				y++
			}
		}

	}

	return newGrid
}

func RemoveNonHorizontalOrVericalLines(allLines []Line) (filteredLines []Line) {
	for _, l := range allLines {
		if l.x1 == l.x2 || l.y1 == l.y2 {
			filteredLines = append(filteredLines, l)
		}
	}
	return filteredLines
}

func GetNumberOfNodesGreaterThanTwo(grid [][]int) (sum int) {

	for _, row := range grid {
		for _, node := range row {
			if node >= 2 {
				sum++
			}
		}
	}
	return sum
}

func min(a, b int) (max int) {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) (max int) {
	if a > b {
		return a
	}
	return b
}
