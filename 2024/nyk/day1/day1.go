package day1

import (
	"2024/utils"
	"math"
	"sort"
	"strconv"
	"strings"
)

func parse_lists(values []string) [2][]int {
	number_lists := [2][]int{}
	number_lists[0] = make([]int, len(values))
	number_lists[1] = make([]int, len(values))

	for string_index, string_value := range values {
		split_values := strings.Split(string_value, "   ")

		for int_index, int_value := range split_values {
			value, err := strconv.Atoi(int_value)

			if err != nil {
				panic("Can't convert string to int")
			}

			number_lists[int_index][string_index] = value
		}
	}

	return number_lists
}

func Part1(path string) int {
	values, err := utils.Strings(path)

	if err != nil {
		panic("Can't load file")
	}

	number_lists := parse_lists(values)

	sort.Ints(number_lists[0])
	sort.Ints(number_lists[1])

	result := 0
	for string_index := range values {
		result += int(math.Abs(float64(number_lists[0][string_index] - number_lists[1][string_index])))
	}

	return result
}

func count_elements(arr []int, element_to_count int) int {
	result := 0

	for index := range arr {
		if arr[index] == element_to_count {
			result += 1
		}
	}

	return result
}

func Part2(path string) int {
	values, err := utils.Strings(path)

	if err != nil {
		panic("Can't load file")
	}

	number_lists := parse_lists(values)

	result := 0
	for _, value := range number_lists[0] {
		result += value * count_elements(number_lists[1], value)
	}

	return result
}
