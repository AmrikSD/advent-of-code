package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, _ := os.Open("input.txt")

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	prev, _ := strconv.Atoi(scanner.Text())
	count := 0
	for scanner.Scan() {
		currNum, _ := strconv.Atoi(scanner.Text())
		if currNum > prev {
			count++
		}
		prev = currNum
	}

	fmt.Println(count)

}
