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
	answer = 0

	file, _ := os.Open(input)

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	prev, _ := strconv.Atoi(scanner.Text())
	for scanner.Scan() {
		currNum, _ := strconv.Atoi(scanner.Text())
		if currNum > prev {
			answer++
		}
		prev = currNum
	}

	return answer

}

func partTwo(input string) (answer int) {

	file, _ := os.Open(input)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var numbers []int

	for scanner.Scan() {
		currNum, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, currNum)
	}

	length := len(numbers)

	prevSum := 0
	for i, _ := range numbers {
		currSum := 0
		for j := i; j < i+3 && i+3 < length; j++ {
			currSum += numbers[j]
		}
		if currSum > prevSum {
			answer++
		}
		prevSum = currSum
	}

	return answer
}
