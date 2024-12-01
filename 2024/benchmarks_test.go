package main

import (
	amrik "2024/amrik"
	nyk "2024/nyk/day1"
	"testing"
)

// Day 1 Part 1
func BenchmarkAmrikDay1Part1(b *testing.B) {
	amrik.Day01Part01("./amrik/inputs/day01/input.txt")
}
func BenchmarkNykDay1Part1(b *testing.B) {
	nyk.Part1("./nyk/day1/input.txt")
}

// Day 1 Part 2
func BenchmarkAmrikDay1Part2(b *testing.B) {
	amrik.Day01Part02("./amrik/inputs/day01/input.txt")
}
func BenchmarkNykDay1Part2(b *testing.B) {
	nyk.Part2("./nyk/day1/input.txt")
}

// Day 2 Part 1
