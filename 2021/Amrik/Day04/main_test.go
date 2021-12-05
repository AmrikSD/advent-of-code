package main

import (
	"reflect"
	"testing"
)

func TestParseInputBoard(t *testing.T) {
	boards, _ := ParseInput("sample.txt")

	expected := []Board{
		{
			[][]Node{
				{
					{value: 22, called: false},
					{value: 13, called: false},
					{value: 17, called: false},
					{value: 11, called: false},
					{value: 0, called: false},
				},
				{
					{value: 8, called: false},
					{value: 2, called: false},
					{value: 23, called: false},
					{value: 4, called: false},
					{value: 24, called: false},
				},
				{
					{value: 21, called: false},
					{value: 9, called: false},
					{value: 14, called: false},
					{value: 16, called: false},
					{value: 7, called: false},
				},
				{
					{value: 6, called: false},
					{value: 10, called: false},
					{value: 3, called: false},
					{value: 18, called: false},
					{value: 5, called: false},
				},
				{
					{value: 1, called: false},
					{value: 12, called: false},
					{value: 20, called: false},
					{value: 15, called: false},
					{value: 19, called: false},
				},
			},
		},
		{
			[][]Node{
				{
					{value: 3, called: false},
					{value: 15, called: false},
					{value: 0, called: false},
					{value: 2, called: false},
					{value: 22, called: false},
				},
				{
					{value: 9, called: false},
					{value: 18, called: false},
					{value: 13, called: false},
					{value: 17, called: false},
					{value: 5, called: false},
				},
				{
					{value: 19, called: false},
					{value: 8, called: false},
					{value: 7, called: false},
					{value: 25, called: false},
					{value: 23, called: false},
				},
				{
					{value: 20, called: false},
					{value: 11, called: false},
					{value: 10, called: false},
					{value: 24, called: false},
					{value: 4, called: false},
				},
				{
					{value: 14, called: false},
					{value: 21, called: false},
					{value: 16, called: false},
					{value: 12, called: false},
					{value: 6, called: false},
				},
			},
		},
		{
			[][]Node{
				{
					{value: 14, called: false},
					{value: 21, called: false},
					{value: 17, called: false},
					{value: 24, called: false},
					{value: 4, called: false},
				},
				{
					{value: 10, called: false},
					{value: 16, called: false},
					{value: 15, called: false},
					{value: 9, called: false},
					{value: 19, called: false},
				},
				{
					{value: 18, called: false},
					{value: 8, called: false},
					{value: 23, called: false},
					{value: 26, called: false},
					{value: 20, called: false},
				},
				{
					{value: 22, called: false},
					{value: 11, called: false},
					{value: 13, called: false},
					{value: 6, called: false},
					{value: 5, called: false},
				},
				{
					{value: 2, called: false},
					{value: 0, called: false},
					{value: 12, called: false},
					{value: 3, called: false},
					{value: 7, called: false},
				},
			},
		},
	}

	if !reflect.DeepEqual(expected, boards) {
		t.Errorf("Expected %v \n actual %v", expected, boards)
	}

}

func TestParseInputDeck(t *testing.T) {
	_, deck := ParseInput("sample.txt")

	expected := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}

	if !reflect.DeepEqual(expected, deck) {
		t.Errorf("Expected %v\n actual %v", expected, deck)
	}

}

func TestGetSumOfNotVisited(t *testing.T) {

	board := Board{
		[][]Node{
			{
				{value: 14, called: true},
				{value: 21, called: true},
				{value: 17, called: true},
				{value: 24, called: true},
				{value: 4, called: true},
			},
			{
				{value: 10, called: false},
				{value: 16, called: false},
				{value: 15, called: false},
				{value: 9, called: true},
				{value: 19, called: false},
			},
			{
				{value: 18, called: false},
				{value: 8, called: false},
				{value: 23, called: true},
				{value: 26, called: false},
				{value: 20, called: false},
			},
			{
				{value: 22, called: false},
				{value: 11, called: true},
				{value: 13, called: false},
				{value: 6, called: false},
				{value: 5, called: true},
			},
			{
				{value: 2, called: true},
				{value: 0, called: true},
				{value: 12, called: false},
				{value: 3, called: false},
				{value: 7, called: true},
			},
		},
	}

	expected := 188
	actual := board.sumOfNotVisited()

	if actual != expected {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

}

func TestIsBingo(t *testing.T) {

	noBingo := Board{[][]Node{
		{
			{value: 9, called: false},
			{value: 9, called: true},
			{value: 9, called: true},
		},
		{
			{value: 9, called: false},
			{value: 9, called: false},
			{value: 9, called: false},
		},
		{
			{value: 9, called: false},
			{value: 9, called: false},
			{value: 9, called: false},
		},
	}}
	ColumnBingo := Board{[][]Node{
		{
			{value: 9, called: true},
			{value: 9, called: false},
			{value: 9, called: false},
		},
		{
			{value: 9, called: true},
			{value: 9, called: false},
			{value: 9, called: false},
		},
		{
			{value: 9, called: true},
			{value: 9, called: false},
			{value: 9, called: false},
		},
	}}
	RowBingo := Board{[][]Node{
		{
			{value: 9, called: true},
			{value: 9, called: true},
			{value: 9, called: true},
		},
		{
			{value: 9, called: false},
			{value: 9, called: false},
			{value: 9, called: false},
		},
		{
			{value: 9, called: false},
			{value: 9, called: false},
			{value: 9, called: false},
		},
	}}

	if noBingo.isBingo() {
		t.Errorf("Expected %v, actual %v", noBingo.isBingo(), false)
	}

	if !RowBingo.isBingo() {
		t.Errorf("Expected %v, actual %v", RowBingo.isBingo(), true)
	}

	if !ColumnBingo.isBingo() {
		t.Errorf("Expected %v, actual %v", ColumnBingo.isBingo(), true)
	}
}

func TestProcessDraw(t *testing.T) {

	board := Board{[][]Node{
		{
			{value: 1, called: true},
			{value: 2, called: true},
			{value: 3, called: true},
		},
		{
			{value: 4, called: false},
			{value: 5, called: false},
			{value: 6, called: false},
		},
		{
			{value: 7, called: false},
			{value: 8, called: false},
			{value: 9, called: false},
		},
	}}

	expected := Board{[][]Node{
		{
			{value: 1, called: true},
			{value: 2, called: true},
			{value: 3, called: true},
		},
		{
			{value: 4, called: false},
			{value: 5, called: false},
			{value: 6, called: false},
		},
		{
			{value: 7, called: false},
			{value: 8, called: false},
			{value: 9, called: true},
		},
	}}

	board.processDraw(9)

	if !reflect.DeepEqual(board, expected) {
		t.Errorf("Expected %v, actual %v", expected, board)
	}

}

func TestPartOne(t *testing.T) {
	expected := 4512
	actual := partOne("sample.txt")

	if expected != actual {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}

func TestPartTwo(t *testing.T) {
	expected := 1924
	actual := partTwo("sample.txt")

	if expected != actual {
		t.Errorf("Expected %v, actual %v", expected, actual)
	}
}
