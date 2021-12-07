package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestParseInput(t *testing.T) {

	input := "sample.txt"
	expected := school{
		3: 2,
		4: 1,
		1: 1,
		2: 1,
	}

	ExpectedKeys := make([]int, 0, len(expected))
	for k := range expected {
		ExpectedKeys = append(ExpectedKeys, k)
	}

	actual := parseInput(input)
	ActualKeys := make([]int, 0, len(actual))
	for k := range actual {
		ActualKeys = append(ActualKeys, k)
	}

	sort.Ints(ActualKeys)
	sort.Ints(ExpectedKeys)
	if !reflect.DeepEqual(ExpectedKeys, ActualKeys) {
		t.Errorf("Expected keys: %v, Actual keys: %v", ExpectedKeys, ActualKeys)
	}

	// we know if we get here the keys are the same, so we can check the value are too
	for _, v := range ActualKeys {
		if actual[v] != expected[v] {
			t.Errorf("%v value is mismatched, actual is %v and expected is %v", v, actual[v], expected[v])
		}
	}

}

func TestGetNextDaySchool(t *testing.T) {

	input := school{
		3: 2,
		4: 1,
		1: 1,
		2: 1,
	}

	expected := school{
		2: 2,
		3: 1,
		0: 1,
		1: 1,
	}
	ExpectedKeys := make([]int, 0, len(expected))
	for k := range expected {
		ExpectedKeys = append(ExpectedKeys, k)
	}

	actual := getNextDay(input)
	ActualKeys := make([]int, 0, len(actual))
	for k := range actual {
		ActualKeys = append(ActualKeys, k)
	}

	sort.Ints(ActualKeys)
	sort.Ints(ExpectedKeys)
	if !reflect.DeepEqual(ExpectedKeys, ActualKeys) {
		t.Errorf("Expected keys: %v, Actual keys: %v", ExpectedKeys, ActualKeys)
	}

	// we know if we get here the keys are the same, so we can check the value are too
	for _, v := range ActualKeys {
		if actual[v] != expected[v] {
			t.Errorf("%v value is mismatched, actual is %v and expected is %v", v, actual[v], expected[v])
		}
	}

}

func TestGetNextDaySchoolWithNewBorns(t *testing.T) {

	input := school{
		2: 2,
		3: 1,
		0: 1,
		1: 1,
	}

	expected := school{
		0: 1,
		1: 2,
		2: 1,
		6: 1,
		8: 1,
	}

	ExpectedKeys := make([]int, 0, len(expected))
	for k := range expected {
		ExpectedKeys = append(ExpectedKeys, k)
	}

	actual := getNextDay(input)
	ActualKeys := make([]int, 0, len(actual))
	for k := range actual {
		ActualKeys = append(ActualKeys, k)
	}

	sort.Ints(ActualKeys)
	sort.Ints(ExpectedKeys)

	if !reflect.DeepEqual(ExpectedKeys, ActualKeys) {
		t.Errorf("Expected keys: %v, Actual keys: %v", ExpectedKeys, ActualKeys)
	}

	// we know if we get here the keys are the same, so we can check the value are too
	for _, v := range ActualKeys {
		if actual[v] != expected[v] {
			t.Errorf("%v value is mismatched, actual is %v and expected is %v", v, actual[v], expected[v])
		}
	}

}

func TestPartOne(t *testing.T) {

	expected := 26
	actual := partOne("sample.txt", 18)

	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

	expected = 5934
	actual = partOne("sample.txt", 80)

	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

}

func TestPartTwo(t *testing.T) {

	expected := 26984457539
	actual := partOne("sample.txt", 256)

	if expected != actual {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

}
