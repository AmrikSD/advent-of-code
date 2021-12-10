package main

import (
	"bufio"
	"os"
)

func ParseInput(filePath string) (heightmap [][]int) {

	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		text := scanner.Text()
		row := []int{}
		for _, numRune := range text {
			row = append(row, int(numRune-'0'))
		}
		heightmap = append(heightmap, row)
	}
	return heightmap
}

func FindLowPoints(heightmap [][]int) (lowpointsCoords []Coord, lowpointsVal []int) {

	for rowIndex, row := range heightmap {

		//do inner safely...
		for colIndex, cell := range row {
			if rowIndex == 0 || rowIndex == len(heightmap)-1 {
				continue
			}
			if colIndex == 0 || colIndex == len(row)-1 {
				continue
			}
			up := heightmap[rowIndex+1][colIndex]
			down := heightmap[rowIndex-1][colIndex]
			left := heightmap[rowIndex][colIndex-1]
			right := heightmap[rowIndex][colIndex+1]

			if cell < up && cell < down && cell < left && cell < right {
				lowpointsVal = append(lowpointsVal, cell)
				lowpointsCoords = append(lowpointsCoords, Coord{rowIndex, colIndex})
			}

		}

	}

	//do outer (not corners)

	// top row
	for i := 1; i < len(heightmap[0])-1; i++ {
		cell := heightmap[0][i]
		down := heightmap[1][i]
		left := heightmap[0][i-1]
		right := heightmap[0][i+1]
		if cell < down && cell < left && cell < right {
			lowpointsVal = append(lowpointsVal, cell)
			lowpointsCoords = append(lowpointsCoords, Coord{0, i})
		}
	}

	//bottom row
	for i := 1; i < len(heightmap[0])-1; i++ {
		cell := heightmap[len(heightmap)-1][i]
		up := heightmap[len(heightmap)-2][i]
		left := heightmap[len(heightmap)-1][i-1]
		right := heightmap[len(heightmap)-1][i+1]

		if cell < up && cell < left && cell < right {
			lowpointsVal = append(lowpointsVal, cell)
			lowpointsCoords = append(lowpointsCoords, Coord{len(heightmap) - 1, i})
		}
	}

	//left column
	for i := 1; i < len(heightmap)-1; i++ {
		cell := heightmap[i][0]
		up := heightmap[i+1][0]
		down := heightmap[i-1][0]
		right := heightmap[i][1]
		if cell < up && cell < down && cell < right {
			lowpointsVal = append(lowpointsVal, cell)
			lowpointsCoords = append(lowpointsCoords, Coord{i, 0})
		}
	}

	//right column
	for i := 1; i < len(heightmap)-1; i++ {
		cell := heightmap[i][len(heightmap)-1]
		up := heightmap[i+1][len(heightmap)-1]
		down := heightmap[i-1][len(heightmap)-1]
		left := heightmap[i][len(heightmap)-2]
		if cell < up && cell < down && cell < left {
			lowpointsVal = append(lowpointsVal, cell)
			lowpointsCoords = append(lowpointsCoords, Coord{i, len(heightmap) - 1})
		}
	}
	//do corners....

	//topRight
	if heightmap[0][0] < heightmap[0][1] && heightmap[0][0] < heightmap[1][0] {
		lowpointsVal = append(lowpointsVal, heightmap[0][0])
		lowpointsCoords = append(lowpointsCoords, Coord{0, 0})
	}
	//topLeft
	if heightmap[0][len(heightmap[0])-1] < heightmap[1][len(heightmap[0])-1] && heightmap[0][len(heightmap[0])-1] < heightmap[0][len(heightmap[0])-2] {
		lowpointsVal = append(lowpointsVal, heightmap[0][len(heightmap[0])-1])
		lowpointsCoords = append(lowpointsCoords, Coord{0, len(heightmap[0]) - 1})
	}
	//bottomleft
	if heightmap[len(heightmap)-1][0] < heightmap[len(heightmap)-1][1] && heightmap[len(heightmap)-1][0] < heightmap[len(heightmap)-2][0] {
		lowpointsVal = append(lowpointsVal, heightmap[len(heightmap)-1][0])
		lowpointsCoords = append(lowpointsCoords, Coord{len(heightmap) - 1, 0})
	}

	//bottomRight
	if heightmap[len(heightmap)-1][len(heightmap[0])-1] < heightmap[len(heightmap)-2][len(heightmap[0])-1] && heightmap[len(heightmap)-1][len(heightmap[0])-1] < heightmap[len(heightmap)-1][len(heightmap[0])-2] {
		lowpointsVal = append(lowpointsVal, heightmap[len(heightmap)-1][len(heightmap[0])-1])
		lowpointsCoords = append(lowpointsCoords, Coord{len(heightmap) - 1, len(heightmap[0]) - 1})
	}

	return lowpointsCoords, lowpointsVal
}

//DFS to find the area of the current basin, as we know they all belong to only one basin we can do this.... I think..
func findBasinArea(heightmap [][]int, visited [][]bool, basinStart Coord) int {

	x := basinStart.x
	y := basinStart.y

	up := Coord{y - 1, x}
	down := Coord{y + 1, x}
	left := Coord{y, x - 1}
	right := Coord{y, x + 1}

	neighbors := []Coord{
		up, down, left, right,
	}

	sum := 1
	visited[y][x] = true
	for _, c := range neighbors {
		if isSafe(c, heightmap, visited) {
			sum += findBasinArea(heightmap, visited, c)
		}
	}

	return sum
}

func isSafe(coord Coord, hm [][]int, visited [][]bool) bool {
	maxY := len(hm) - 1
	maxX := len(hm[0]) - 1

	return coord.y >= 0 && coord.y <= maxY && coord.x >= 0 && coord.x <= maxX && !visited[coord.y][coord.x] && hm[coord.y][coord.x] != 9
}

type Coord struct {
	y int
	x int
}
