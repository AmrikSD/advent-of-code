package main

import (
	"bufio"
	"os"
	"reflect"
	"sort"
	"strings"
)

func parseInput(filePath string) (input [][]string, output [][]string) {
	file, _ := os.Open(filePath)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		inAndOut := strings.Split(text, " | ")
		lineOutputArr := strings.Split(inAndOut[1], " ")
		lineInputArr := strings.Split(inAndOut[0], " ")
		output = append(output, lineOutputArr)
		input = append(input, lineInputArr)
	}

	return input, output
}

// This esentially returns the amount of times a 1,4,7 or 8 are displayed.
func getNumberOfDigitsWithUniqueSegments(segments [][]string) int {

	interestingNumbers := []int{2, 3, 4, 7}

	sum := 0
	for _, v := range segments {
		for _, segment := range v {
			for _, num := range interestingNumbers {
				if len(segment) == num {
					sum++
					continue
				}
			}
		}
	}

	return sum

}

func MapSignalToUniqueDigits(signal []string) map[int]string {

	hm := make(map[int]string)

	// This gets 1,4,7,8 - still need 0,2,3,5,6,9
	for _, signal := range signal {
		switch len(signal) {
		case 2:
			hm[1] = signal
		case 4:
			hm[4] = signal
		case 3:
			hm[7] = signal
		case 7:
			hm[8] = signal
		}
	}

	return hm
}

//Three = len(n) == 5 && n contians two elements from One
func addThreeToSignals(startingSignals map[int]string, inputs []string) (signalsWithThree map[int]string) {

	signalsWithThree = startingSignals
	One := startingSignals[1]

	for _, n := range inputs {
		//len(n)==5
		if len(n) == 5 {
			// && n contains two elements from One
			total := 0
			for _, r := range n {
				for _, oneRune := range One {
					if r == oneRune {
						total++
					}
				}
			}
			if total == 2 {
				signalsWithThree[3] = n
			}
		}

	}
	return signalsWithThree
}

// Five = len(n) == 5 && n contains 3 elements from Four && n != Three
func addFiveToSignals(startingSignals map[int]string, inputs []string) (signalsWithFive map[int]string) {

	signalsWithFive = startingSignals
	Four := startingSignals[4]
	Three := startingSignals[3]

	for _, n := range inputs {
		//len(n)==5
		if len(n) == 5 {
			total := 0
			for _, r := range n {
				for _, oneRune := range Four {
					if r == oneRune {
						total++
					}
				}
			}
			if total == 3 && !reflect.DeepEqual(n, Three) {
				signalsWithFive[5] = n
			}
		}
	}
	return signalsWithFive
}

// Two = len(n) == 5 && n != Five && n != Three
func addTwoToSignals(startingSignals map[int]string, inputs []string) (signalsWithTwo map[int]string) {

	signalsWithTwo = startingSignals
	Three := startingSignals[3]
	Five := startingSignals[5]

	for _, n := range inputs {
		//len(n)==5
		if len(n) == 5 {
			// && n!= five && n!= three
			if !reflect.DeepEqual(n, Three) && !reflect.DeepEqual(n, Five) {
				signalsWithTwo[2] = n
			}
		}
	}
	return signalsWithTwo
}

// Nine = len(n) == 6 && n contains 4 elements from Four
func addNineToSignals(startingSignals map[int]string, inputs []string) (signalsWithNine map[int]string) {

	signalsWithNine = startingSignals
	Four := startingSignals[4]

	for _, n := range inputs {
		//len(n)==6
		if len(n) == 6 {
			// && n contains 4 elements from Four
			total := 0
			for _, r := range n {
				for _, oneRune := range Four {
					if r == oneRune {
						total++
					}
				}
			}
			if total == 4 {
				signalsWithNine[9] = n
			}
		}
	}
	return signalsWithNine
}

// Six = len(n) == 6 && n contains 5 elements from Five && n != Nine
func addSixToSignals(startingSignals map[int]string, inputs []string) (signalsWithSix map[int]string) {

	signalsWithSix = startingSignals
	Five := startingSignals[5]
	Nine := startingSignals[9]

	for _, n := range inputs {
		//len(n)==6
		if len(n) == 6 {
			// && n contains 5 elements from Five
			total := 0
			for _, r := range n {
				for _, oneRune := range Five {
					if r == oneRune {
						total++
					}
				}
			}
			if total == 5 && !reflect.DeepEqual(n, Nine) {
				signalsWithSix[6] = n
			}
		}
	}
	return signalsWithSix
}

// Zero = len(n) ==6 && n!= Six && n!=Nine
func addZeroToSignals(startingSignals map[int]string, inputs []string) (signalsWithZero map[int]string) {

	signalsWithZero = startingSignals
	Six := startingSignals[6]
	Nine := startingSignals[9]

	for _, n := range inputs {
		//len(n)==6
		if len(n) == 6 {
			// && n!=Six && n!=Nine
			if !reflect.DeepEqual(n, Six) && !reflect.DeepEqual(n, Nine) {
				signalsWithZero[0] = n
			}
		}
	}
	return signalsWithZero
}

func createMapFromSignalToValue(signals map[int]string) map[string]int {

	output := make(map[string]int)
	for k, v := range signals {
		output[SortStringByCharacter(v)] = k
	}
	return output

}

func SortStringByCharacter(s string) string {
	r := StringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}
