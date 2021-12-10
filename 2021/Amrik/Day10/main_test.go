package main

import (
	"reflect"
	"testing"
)

func TestPartOne(T *testing.T) {

	actual := partOne("sample.txt")
	expected := 26397
	if expected != actual {
		T.Errorf("expected: %v, actual: %v", expected, actual)
	}

}

func TestPartTwo(T *testing.T) {

	actual := partTwo("sample.txt")
	expected := 288957
	if expected != actual {
		T.Errorf("expected: %v, actual: %v", expected, actual)
	}

}
func TestParseInput(T *testing.T) {

	expected := []string{
		"[({(<(())[]>[[{[]{<()<>>",
		"[(()[<>])]({[<{<<[]>>(",
		"{([(<{}[<>[]}>{[]{[(<()>",
		"(((({<>}<{<{<>}{[]{[]{}",
		"[[<[([]))<([[{}[[()]]]",
		"[{[{({}]{}}([{[{{{}}([]",
		"{<[[]]>}<{[{[{[]{()[[[]",
		"[<(<(<(<{}))><([]([]()",
		"<{([([[(<>()){}]>(<<{{",
		"<{([{{}}[<[[[<>{}]]]>[]]",
	}

	actual := ParseInput("sample.txt")

	if !reflect.DeepEqual(expected, actual) {
		T.Errorf("expected: %v actual :%v", expected, actual)
	}

}
