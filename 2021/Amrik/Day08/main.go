package main

import (
	"fmt"
	"math"
)

func main() {

	input := "input.txt"
	fmt.Println("part 1: ", partOne(input))
	fmt.Println("part 2: ", partTwo(input))

}
func partOne(input string) (answer int) {

	_, outputs := parseInput(input)

	return getNumberOfDigitsWithUniqueSegments(outputs)

}

func partTwo(input string) (answer int) {
	inputs, outputs := parseInput(input)

	sum := 0.0
	for i, v := range inputs {
		currSignals := MapSignalToUniqueDigits(v)

		//The order matters here, as we can only add Two after we know what Five and Three are..
		currSignals = addThreeToSignals(currSignals, v)
		currSignals = addFiveToSignals(currSignals, v)

		currSignals = addTwoToSignals(currSignals, v)
		currSignals = addNineToSignals(currSignals, v)
		currSignals = addSixToSignals(currSignals, v)
		currSignals = addZeroToSignals(currSignals, v)

		valueSignals := createMapFromSignalToValue(currSignals)
		fmt.Println(valueSignals)

		innerSum := 0.0
		for j, str := range outputs[i] {
			sortedStr := SortStringByCharacter(str)
			innerSum += math.Pow(10, float64(3-j)) * float64(valueSignals[sortedStr])
		}
		sum += innerSum
	}
	return int(sum)
}
