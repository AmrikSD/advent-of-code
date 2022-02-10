package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	input := "input.txt"
	fmt.Println("part 1: ", partOne(input))
	fmt.Println("part 2: ", partTwo(input))

}

func partOne(input string) (answer int) {

	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)

	var length int
	used := make(map[int]int)

	for scanner.Scan() {
		length++
		number := scanner.Text()

		for pos, char := range number {
			used[pos] += int(char - '0')
		}

	}

	gamma := ""
	epsilon := ""

	for i := 0; i < len(used); i++ {
		if used[i] >= length/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}

	gammaInt, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonInt, _ := strconv.ParseInt(epsilon, 2, 64)
	return int(gammaInt * epsilonInt)

}

func partTwo(input string) (answer int) {

	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)

	allInput := []string{}

	for scanner.Scan() {
		in := scanner.Text()
		allInput = append(allInput, in)
	}

	// Find Oxygen
	oxygen, _ := strconv.ParseInt(findOxygen(allInput), 2, 64)

	// Find Co2
	CO2, _ := strconv.ParseInt(findCO2(allInput), 2, 64)
	return int(oxygen) * int(CO2)

}

func findOxygen(input []string) string {

	// Consider just the first bit.

	// Keep only nunmbers selected by the bit criteria
	// Discard numbers which aren't selected by the criteria.

	// if there is only one number left, return this number.
	// Otherwise, repeat this process, considering the next bit to the right.

	i := 0
	for len(input) > 1 {
		input = filterByOxygenCriteria(input, i)
		i++
	}
	return input[0]
}

func filterByOxygenCriteria(input []string, pos int) (out []string) {
	// To find oxygen generator rating, determine the most common value (0 or 1) in the current bit position,

	sum := 0
	for _, v := range input {
		if v[pos]-'0' == 1 {
			sum++
		}
	}
	bit := 0
	if sum*2 >= len(input) {
		bit = 1
	}

	// and keep only numbers with that bit in that position.
	// If 0 and 1 are equally common, keep values with a 1 in the position being considered.

	for _, v := range input {
		if int(v[pos]-'0') == bit {
			out = append(out, v)
		}
	}
	return out
}

func findCO2(input []string) string {

	// Consider just the first bit.

	// Keep only nunmbers selected by the bit criteria
	// Discard numbers which aren't selected by the criteria.

	// if there is only one number left, return this number.
	// Otherwise, repeat this process, considering the next bit to the right.

	i := 0
	for len(input) > 1 {
		input = filterByCO2Criteria(input, i)
		i++
	}
	return input[0]
}

func filterByCO2Criteria(input []string, pos int) (out []string) {
	// To find oxygen generator rating, determine the most common value (0 or 1) in the current bit position,

	sum := 0
	for _, v := range input {
		if v[pos]-'0' == 1 {
			sum++
		}
	}
	bit := 0
	if sum*2 >= len(input) {
		bit = 1
	}
	//conver to unpoopular bit
	bit = (bit - 1) * -1

	// and keep only numbers with that bit in that position.
	// If 0 and 1 are equally common, keep values with a 1 in the position being considered.

	for _, v := range input {
		if int(v[pos]-'0') == bit {
			out = append(out, v)
		}
	}
	return out
}
