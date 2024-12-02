package day2

import (
	"2024/utils"
	"math"
	"strconv"
	"strings"
)

func convert_report_to_number(string_report []string) []int {
	number_report := make([]int, len(string_report))

	for index, string_number := range string_report {
		number, err := strconv.Atoi(string_number)

		if err != nil {
			panic("Couldn't convert number in report")
		}

		number_report[index] = number
	}

	return number_report
}

func determine_safety(number_report []int) bool {
	index := 0
	length := len(number_report)

	direction := 0
	is_safe := true

	for is_safe && index != length-1 {
		diff := number_report[index+1] - number_report[index]

		// Set direction for report
		if direction == 0 {
			if diff < 0 {
				direction = -1
			} else if diff > 0 {
				direction = 1
			} else {
				is_safe = false
				break
			}
		}

		if (diff > 0 && direction > 0) || (diff < 0 && direction < 0) {
			if math.Abs(float64(diff)) > 3 {
				is_safe = false
				break
			}
		} else {
			is_safe = false
			break
		}

		index += 1
	}

	return is_safe
}

func Part1(path string) int {
	reports, err := utils.Strings(path)

	if err != nil {
		panic("Can't load file")
	}

	result := 0

	for _, report := range reports {
		string_report := strings.Split(report, " ")

		number_report := convert_report_to_number(string_report)

		if !determine_safety(number_report) {
			continue
		}

		result += 1
	}

	return result
}

// https://stackoverflow.com/a/57213476
func remove_at_index(s []int, index int) []int {
	ret := make([]int, 0, len(s)-1)
	ret = append(ret, s[:index]...)

	return append(ret, s[index+1:]...)
}

func Part2(path string) int {
	reports, err := utils.Strings(path)

	if err != nil {
		panic("Can't load file")
	}

	result := 0

	for _, report := range reports {
		string_report := strings.Split(report, " ")

		number_report := convert_report_to_number(string_report)

		is_safe := determine_safety(number_report)

		if is_safe {
			result += 1
		} else {
			is_new_report_safe := false
			index := 0

			for !is_new_report_safe && index != len(number_report) {
				new_report := remove_at_index(number_report, index)

				if determine_safety(new_report) {
					is_new_report_safe = true
					break
				}

				index += 1
			}

			if is_new_report_safe {
				result += 1
			}
		}
	}

	return result
}
