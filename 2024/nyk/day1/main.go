package nyk

import (
	"2024/utils"
	"fmt"
)

func Part1(path string) int {
	values, err := utils.Strings(path)

	if err != nil {
		panic("Can't load file")
	}

	for _, value := range values {
		fmt.Println(value)
	}

	return 0
}

func Part2(path string) int {
	return 0
}
