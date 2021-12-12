package main

import (
	"reflect"
	"testing"
)

func TestPartOne(T *testing.T) {
}

func TestPartTwo(T *testing.T) {
}

func TestParseInput(T *testing.T) {
	input := "sample.txt"

	expected := [][]Octopus{
		{{5, false}, {4, false}, {8, false}, {3, false}, {1, false}, {4, false}, {3, false}, {2, false}, {2, false}, {3, false}},
		{{2, false}, {7, false}, {4, false}, {5, false}, {8, false}, {5, false}, {4, false}, {7, false}, {1, false}, {1, false}},
		{{5, false}, {2, false}, {6, false}, {4, false}, {5, false}, {5, false}, {6, false}, {1, false}, {7, false}, {3, false}},
		{{6, false}, {1, false}, {4, false}, {1, false}, {3, false}, {3, false}, {6, false}, {1, false}, {4, false}, {6, false}},
		{{6, false}, {3, false}, {5, false}, {7, false}, {3, false}, {8, false}, {5, false}, {4, false}, {7, false}, {8, false}},
		{{4, false}, {1, false}, {6, false}, {7, false}, {5, false}, {2, false}, {4, false}, {6, false}, {4, false}, {5, false}},
		{{2, false}, {1, false}, {7, false}, {6, false}, {8, false}, {4, false}, {1, false}, {7, false}, {2, false}, {1, false}},
		{{6, false}, {8, false}, {8, false}, {2, false}, {8, false}, {8, false}, {1, false}, {1, false}, {3, false}, {4, false}},
		{{4, false}, {8, false}, {4, false}, {6, false}, {8, false}, {4, false}, {8, false}, {5, false}, {5, false}, {4, false}},
		{{5, false}, {2, false}, {8, false}, {3, false}, {7, false}, {5, false}, {1, false}, {5, false}, {2, false}, {6, false}},
	}

	actual := ParseInput(input)

	if !reflect.DeepEqual(expected, actual) {
		T.Errorf("Error with parse input function\n Expected: %v\n Actual: %v\n", expected, actual)
	}

}

func TestProgressGridWithFlashes(T *testing.T) {

	input := [][]Octopus{
		{{1, false}, {1, false}, {1, false}, {1, false}, {1, false}},
		{{1, false}, {9, false}, {9, false}, {9, false}, {1, false}},
		{{1, false}, {9, false}, {1, false}, {9, false}, {1, false}},
		{{1, false}, {9, false}, {9, false}, {9, false}, {1, false}},
		{{1, false}, {1, false}, {1, false}, {1, false}, {1, false}},
	}

	expected := [][]Octopus{
		{{3, false}, {4, false}, {5, false}, {4, false}, {3, false}},
		{{4, false}, {0, false}, {0, false}, {0, false}, {4, false}},
		{{5, false}, {0, false}, {0, false}, {0, false}, {5, false}},
		{{4, false}, {0, false}, {0, false}, {0, false}, {4, false}},
		{{3, false}, {4, false}, {5, false}, {4, false}, {3, false}},
	}

	_, actual := ProgressGrid(input)

	if !reflect.DeepEqual(expected, actual) {
		T.Errorf("Error with progress grid (with flashes) function\n Expected: %v\n Actual: %v\n", expected, actual)
	}

}

func TestProgressGridWithoutFlashes(T *testing.T) {

	input := [][]Octopus{
		{{3, false}, {4, false}, {5, false}, {4, false}, {3, false}},
		{{4, false}, {0, false}, {0, false}, {0, false}, {4, false}},
		{{5, false}, {0, false}, {0, false}, {0, false}, {5, false}},
		{{4, false}, {0, false}, {0, false}, {0, false}, {4, false}},
		{{3, false}, {4, false}, {5, false}, {4, false}, {3, false}},
	}

	expected := [][]Octopus{
		{{4, false}, {5, false}, {6, false}, {5, false}, {4, false}},
		{{5, false}, {1, false}, {1, false}, {1, false}, {5, false}},
		{{6, false}, {1, false}, {1, false}, {1, false}, {6, false}},
		{{5, false}, {1, false}, {1, false}, {1, false}, {5, false}},
		{{4, false}, {5, false}, {6, false}, {5, false}, {4, false}},
	}

	_, actual := ProgressGrid(input)

	if !reflect.DeepEqual(expected, actual) {
		for i, row := range actual {
			for j, o := range row {
				if !reflect.DeepEqual(expected[i][j], o) {
					T.Errorf("%v is not same as %v", expected[i][j], o)
				}
			}
		}
		T.Errorf("Error with progress grid (with flashes) function\n Expected: %v\n Actual: %v\n", expected, actual)
	}

}
