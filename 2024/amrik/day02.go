package amrik

import (
	"2024/utils"
	"strconv"
	"strings"
)

// Must diff by at least N, at most M, inclusive.
func MinMaxDiffByNM(arr []int, n int, m int) bool {

	if len(arr) < 2 {
		return true
	}
	for i := 0; i < len(arr)-1; i++ {
		delta := arr[i] - arr[i+1]
		if delta < 0 {
			delta = -delta
		}
		if delta < n || delta > m {
			return false
		}
	}
	return true
}

func Day02Part01(path string) int {
	strs, err := utils.Strings(path)
	if err != nil {
		panic(err.Error)
	}

	var numbersArr [][]int
	for _, str := range strs {
		arr := strings.Split(str, " ")
		var numbers []int
		for _, n := range arr {
			n, err := strconv.Atoi(n)
			if err != nil {
				panic(err.Error)
			}
			numbers = append(numbers, n)
		}
		numbersArr = append(numbersArr, numbers)
	}

	var monotonic [][]int
	for _, v := range numbersArr {
		if utils.IsStrictlyMonotonic(v) {
			monotonic = append(monotonic, v)
		}
	}

	var monotonicAndMinMaxDiff [][]int
	for _, v := range monotonic {
		if MinMaxDiffByNM(v, 1, 3) {
			monotonicAndMinMaxDiff = append(monotonicAndMinMaxDiff, v)
		}
	}

	return len(monotonicAndMinMaxDiff)

}

func GenerateVariants(input []int) [][]int {
	// Create a slice to hold the results, starting with the original slice
	var result [][]int

	// Generate all variations by removing one element at a time
	for i := range input {
		// Create a new slice with one element removed
		variant := append([]int{}, input[:i]...)
		variant = append(variant, input[i+1:]...)
		result = append(result, variant)
	}

	return result
}

func Day02Part02(path string) int {
	strs, err := utils.Strings(path)
	if err != nil {
		panic(err.Error)
	}

	var numbersArr [][]int
	for _, str := range strs {
		arr := strings.Split(str, " ")
		var numbers []int
		for _, n := range arr {
			n, err := strconv.Atoi(n)
			if err != nil {
				panic(err.Error)
			}
			numbers = append(numbers, n)
		}
		numbersArr = append(numbersArr, numbers)
	}

	var takeTwo [][]int
	var monotonic [][]int
	for _, v := range numbersArr {
		if utils.IsStrictlyMonotonic(v) {
			monotonic = append(monotonic, v)
		} else {
			takeTwo = append(takeTwo, v)
		}
	}

	var monotonicAndMinMaxDiff [][]int
	for _, v := range monotonic {
		if MinMaxDiffByNM(v, 1, 3) {
			monotonicAndMinMaxDiff = append(monotonicAndMinMaxDiff, v)
		} else {
			takeTwo = append(takeTwo, v)
		}
	}

	sum := 0
	for _, v := range takeTwo {
		arrs := GenerateVariants(v)
		for _, arr := range arrs {
			if !utils.IsStrictlyMonotonic(arr) {
				continue
			}
			if !MinMaxDiffByNM(arr, 1, 3) {
				continue
			}
			sum++
			break
		}
	}

	return len(monotonicAndMinMaxDiff) + sum
}

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
