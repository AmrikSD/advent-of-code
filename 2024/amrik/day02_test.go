package amrik

import (
	"testing"
)

func TestDay02Part01(t *testing.T) {
	answer := Day02Part01("./inputs/day02/test.txt")
	expected := 2
	if answer != expected {
		t.Errorf("Got %d, expected %d", answer, expected)
	}
}

func TestDay02Part02(t *testing.T) {
	answer := Day02Part02("./inputs/day02/test.txt")
	expected := 4
	if answer != expected {
		t.Errorf("Got %d, expected %d", answer, expected)
	}
}
