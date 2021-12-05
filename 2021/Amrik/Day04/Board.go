package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	nodes [][]Node
}

type Node struct {
	value  int
	called bool
}

func (b Board) processDraw(numDrew int) {

	for i, v := range b.nodes {
		for j, n := range v {
			if n.value == numDrew {
				b.nodes[i][j] = Node{value: n.value, called: true}
			}
		}
	}

}

func (b Board) sumOfNotVisited() (sum int) {

	for _, v := range b.nodes {
		for _, n := range v {
			if !n.called {
				sum += n.value
			}
		}
	}

	return sum
}

func (b Board) isBingo() (isBingo bool) {
	return b.checkColumns() || b.checkRows()
}

func (b Board) checkRows() (isBingo bool) {

	for colIndex := 0; colIndex < len(b.nodes); colIndex++ {
		count := 0
		for rowIndex := 0; rowIndex < len(b.nodes); rowIndex++ {
			if b.nodes[rowIndex][colIndex].called {
				count++
			}
		}
		if count == len(b.nodes) {
			return true
		}
	}

	return false
}

func (b Board) checkColumns() (isBingo bool) {

	for colIndex := 0; colIndex < len(b.nodes); colIndex++ {
		count := 0
		for rowIndex := 0; rowIndex < len(b.nodes); rowIndex++ {
			if b.nodes[colIndex][rowIndex].called {
				count++
			}
		}
		if count == len(b.nodes) {
			return true
		}
	}

	return false
}

func ParseInput(filePath string) (boards []Board, draws []int) {

	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	firstLine := strings.Split(scanner.Text(), ",")
	for _, v := range firstLine {
		value, _ := strconv.Atoi(v)
		draws = append(draws, value)
	}

	scanner.Scan()
	rows := [][]Node{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			boards = append(boards, Board{rows})
			rows = [][]Node{}
		} else {
			row := []Node{}
			rowStr := strings.Fields(line)
			for _, v := range rowStr {
				val, _ := strconv.Atoi(v)
				row = append(row, Node{value: val, called: false})
			}
			rows = append(rows, row)
		}
	}
	return boards, draws
}
