package amrik

import (
	"testing"
)

func TestDay01Part01(t *testing.T) {
	answer := Day01Part01("./inputs/day01/test.txt")
	expected := 11
	if answer != expected {
		t.Errorf("Got %d, expected %d", answer, expected)
	}
}

func TestDay01Part02(t *testing.T) {
	answer := Day01Part02("./inputs/day01/test.txt")
	expected := 31
	if answer != expected {
		t.Errorf("Got %d, expected %d", answer, expected)
	}
}
