package main

import (
	"reflect"
	"testing"
)

func TestGetDistanceForCrabsToHorizonalSpaceP1(t *testing.T) {

	input := map[int]int{
		0:  1,
		1:  2,
		2:  3,
		4:  1,
		7:  1,
		14: 1,
		16: 1,
	}

	expected := 37
	actual := getDistanceForCrabsToHorizonalSpaceP1(input)

	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestGetDistanceForCrabsToHorizonalSpaceP2(t *testing.T) {

	input := map[int]int{
		0:  1,
		1:  2,
		2:  3,
		4:  1,
		7:  1,
		14: 1,
		16: 1,
	}

	expected := 37
	actual := getDistanceForCrabsToHorizonalSpaceP2(input)

	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}
}

func TestParseInput(t *testing.T) {
	expected := map[int]int{
		0:  1,
		1:  2,
		2:  3,
		4:  1,
		7:  1,
		14: 1,
		16: 1,
	}
	actual := parseInput("sample.txt")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual %v", expected, actual)
	}

}
