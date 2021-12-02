package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	input := "input.txt"

	fmt.Println("part 1: ", partOne(input))
	fmt.Println("part 2: ", partTwo(input))

}

func partOne(input string) (answer int) {

	horizontal := 0
	depth := 0

	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		in := scanner.Text()
		direction := strings.Split(in, " ")[0]
		velocitystr := strings.Split(in, " ")[1]
		velocity, _ := strconv.Atoi(velocitystr)

		switch direction {
		case "forward":
			horizontal += velocity
		case "down":
			depth += velocity
		case "up":
			depth -= velocity
		}

	}

	return depth * horizontal

}

func partTwo(input string) (answer int) {

	horizontal := 0
	depth := 0
	aim := 0

	file, _ := os.Open(input)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		in := scanner.Text()
		direction := strings.Split(in, " ")[0]
		velocitystr := strings.Split(in, " ")[1]
		velocity, _ := strconv.Atoi(velocitystr)

		switch direction {
		case "forward":
			horizontal += velocity
			depth += (aim * velocity)
		case "down":
			aim += velocity
		case "up":
			aim -= velocity
		}

	}

	return depth * horizontal

}
