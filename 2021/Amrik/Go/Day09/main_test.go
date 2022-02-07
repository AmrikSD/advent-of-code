package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestPartOne(T *testing.T) {

	expected := 15
	actual := partOne("sample.txt")
	if expected != actual {
		T.Errorf("expected: %v, actual: %v", expected, actual)
	}

}

func TestFindLowPoint(T *testing.T) {
	input := [][]int{
		{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
		{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
		{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
		{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
		{9, 8, 9, 9, 9, 6, 5, 6, 5, 8},
	}

	_, actual := FindLowPoints(input)

	expected := []int{1, 0, 5, 5}

	sort.Ints(expected)
	sort.Ints(actual)

	if !reflect.DeepEqual(expected, actual) {
		T.Errorf("expected: %v, actual: %v", expected, actual)
	}

}
