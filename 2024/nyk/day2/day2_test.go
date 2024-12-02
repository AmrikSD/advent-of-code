package day2

import (
	"testing"
)

func TestPart1(t *testing.T) {
	answer := Part1("./test.txt")
	expected := 2

	if answer != expected {
		t.Errorf("Got %d, expected %d", answer, expected)
	}
}

func TestPart2(t *testing.T) {
	answer := Part2("./test.txt")
	expected := 4

	if answer != expected {
		t.Errorf("Got %d, expected %d", answer, expected)
	}
}
