package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {

	expected := [][]string{
		{"fdgacbe", "cefdb", "cefbgd", "gcbe"},
		{"fcgedb", "cgb", "dgebacf", "gc"},
		{"cg", "cg", "fdcagb", "cbg"},
		{"efabcd", "cedba", "gadfec", "cb"},
		{"gecf", "egdcabf", "bgf", "bfgea"},
		{"gebdcfa", "ecba", "ca", "fadegcb"},
		{"cefg", "dcbef", "fcge", "gbcadfe"},
		{"ed", "bcgafe", "cdgba", "cbgef"},
		{"gbdfcae", "bgc", "cg", "cgb"},
		{"fgae", "cfgab", "fg", "bagce"},
	}

	_, actual := parseInput("sample.txt")

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, Actual: %v", expected, actual)
	}

}

// 1,4,7,8 all use a unique number of segments (2,4,3,7)
func TestGetNumberOfDigitsWithUniqueSegments(t *testing.T) {
	input := [][]string{
		{"fdgacbe", "cefdb", "cefbgd", "gcbe"},
		{"fcgedb", "cgb", "dgebacf", "gc"},
		{"cg", "cg", "fdcagb", "cbg"},
		{"efabcd", "cedba", "gadfec", "cb"},
		{"gecf", "egdcabf", "bgf", "bfgea"},
		{"gebdcfa", "ecba", "ca", "fadegcb"},
		{"cefg", "dcbef", "fcge", "gbcadfe"},
		{"ed", "bcgafe", "cdgba", "cbgef"},
		{"gbdfcae", "bgc", "cg", "cgb"},
		{"fgae", "cfgab", "fg", "bagce"},
	}
	expected := 26

	actual := getNumberOfDigitsWithUniqueSegments(input)

	if expected != actual {
		t.Errorf("Expected: %v, Actual %v", expected, actual)
	}
}
