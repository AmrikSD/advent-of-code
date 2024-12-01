package amrik

import (
	"2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Day01Part01(path string) int {

	strs, err := utils.Strings(path)
	if err != nil {
		panic(err.Error)
	}

	var leftArr []int
	var rightArr []int

	for _, str := range strs {
		arr := strings.Split(str, "   ")

		left, err := strconv.Atoi(arr[0])
		if err != nil {
			panic(fmt.Sprintf("\"%s\" was not a number?", arr[0]))
		}
		right, err := strconv.Atoi(arr[1])
		if err != nil {
			panic(fmt.Sprintf("\"%s\" was not a number?", arr[1]))
		}
		leftArr = append(leftArr, left)
		rightArr = append(rightArr, right)
	}

	slices.Sort(leftArr)
	slices.Sort(rightArr)

	if len(leftArr) != len(rightArr) {
		panic("left array not equal to right")
	}

	sum := 0
	for i, _ := range leftArr {
		diff := leftArr[i] - rightArr[i]
		if diff < 0 {
			diff = diff * -1
		}
		sum += diff
	}

	return sum
}

func Day01Part02(path string) int {

	strs, err := utils.Strings(path)
	if err != nil {
		panic(err.Error)
	}

	var leftArr []int
	rightMap := make(map[int]int)

	for _, str := range strs {
		arr := strings.Split(str, "   ")

		left, err := strconv.Atoi(arr[0])
		if err != nil {
			panic(fmt.Sprintf("\"%s\" was not a number?", arr[0]))
		}
		right, err := strconv.Atoi(arr[1])
		if err != nil {
			panic(fmt.Sprintf("\"%s\" was not a number?", arr[1]))
		}
		leftArr = append(leftArr, left)

		rightMap[right] += 1
	}

	sum := 0
	for _, i := range leftArr {
		sum += i * rightMap[i]
	}

	return sum
}
