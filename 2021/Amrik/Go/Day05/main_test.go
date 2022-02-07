package main

import (
	"reflect"
	"testing"
)

func TestFilterLinesRemovesNonHorizontalOrVerticalLines(t *testing.T) {

	input := []Line{
		{
			x1: 0,
			y1: 9,
			x2: 5,
			y2: 9,
		},
		{
			x1: 8,
			y1: 0,
			x2: 0,
			y2: 8,
		},
		{
			x1: 9,
			y1: 4,
			x2: 3,
			y2: 4,
		},
		{
			x1: 2,
			y1: 2,
			x2: 2,
			y2: 1,
		},
		{
			x1: 7,
			y1: 0,
			x2: 7,
			y2: 4,
		},
		{
			x1: 6,
			y1: 4,
			x2: 2,
			y2: 0,
		},
		{
			x1: 0,
			y1: 9,
			x2: 2,
			y2: 9,
		},
		{
			x1: 3,
			y1: 4,
			x2: 1,
			y2: 4,
		},
		{
			x1: 0,
			y1: 0,
			x2: 8,
			y2: 8,
		},
		{
			x1: 5,
			y1: 5,
			x2: 8,
			y2: 2,
		},
	}

	expected := []Line{
		{
			x1: 0,
			y1: 9,
			x2: 5,
			y2: 9,
		},
		{
			x1: 9,
			y1: 4,
			x2: 3,
			y2: 4,
		},
		{
			x1: 2,
			y1: 2,
			x2: 2,
			y2: 1,
		},
		{
			x1: 7,
			y1: 0,
			x2: 7,
			y2: 4,
		},
		{
			x1: 0,
			y1: 9,
			x2: 2,
			y2: 9,
		},
		{
			x1: 3,
			y1: 4,
			x2: 1,
			y2: 4,
		},
	}

	actual := RemoveNonHorizontalOrVericalLines(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestGetLines(t *testing.T) {
	input := "sample.txt"

	expected := []Line{
		{
			x1: 0,
			y1: 9,
			x2: 5,
			y2: 9,
		},
		{
			x1: 8,
			y1: 0,
			x2: 0,
			y2: 8,
		},
		{
			x1: 9,
			y1: 4,
			x2: 3,
			y2: 4,
		},
		{
			x1: 2,
			y1: 2,
			x2: 2,
			y2: 1,
		},
		{
			x1: 7,
			y1: 0,
			x2: 7,
			y2: 4,
		},
		{
			x1: 6,
			y1: 4,
			x2: 2,
			y2: 0,
		},
		{
			x1: 0,
			y1: 9,
			x2: 2,
			y2: 9,
		},
		{
			x1: 3,
			y1: 4,
			x2: 1,
			y2: 4,
		},
		{
			x1: 0,
			y1: 0,
			x2: 8,
			y2: 8,
		},
		{
			x1: 5,
			y1: 5,
			x2: 8,
			y2: 2,
		},
	}

	actual := GetLines(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual %v", expected, actual)
	}

}

func TestCreateGridPartOne(t *testing.T) {

	input := []Line{
		{
			x1: 0,
			y1: 9,
			x2: 5,
			y2: 9,
		},
		{
			x1: 9,
			y1: 4,
			x2: 3,
			y2: 4,
		},
		{
			x1: 2,
			y1: 2,
			x2: 2,
			y2: 1,
		},
		{
			x1: 7,
			y1: 0,
			x2: 7,
			y2: 4,
		},
		{
			x1: 0,
			y1: 9,
			x2: 2,
			y2: 9,
		},
		{
			x1: 3,
			y1: 4,
			x2: 1,
			y2: 4,
		},
	}

	expected := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, //0
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, //1
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, //2
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, //3
		{0, 1, 1, 2, 1, 1, 1, 2, 1, 1}, //4
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //5
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //6
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //7
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //8
		{2, 2, 2, 1, 1, 1, 0, 0, 0, 0}, //9
	}

	actual := CreateGrid(input)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual %v", expected, actual)
	}

}

func TestGetNumberOfNodesGreaterThanTwo(t *testing.T) {

	input := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, //0
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, //1
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0}, //2
		{0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, //3
		{0, 1, 1, 2, 1, 1, 1, 2, 1, 1}, //4
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //5
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //6
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //7
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, //8
		{2, 2, 2, 1, 1, 1, 0, 0, 0, 0}, //9
	}

	expected := 5

	actual := GetNumberOfNodesGreaterThanTwo(input)

	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

}

func TestPartOne(t *testing.T) {

	expected := 5
	actual := partOne("sample.txt")

	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

}

func TestPartTwo(t *testing.T) {

	expected := 12
	actual := partTwo("sample.txt")

	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

}
