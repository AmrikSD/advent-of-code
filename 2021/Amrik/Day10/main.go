package main

import (
	"fmt"
	"sort"
)

func main() {

	input := "input.txt"
	fmt.Println("part 1: ", partOne(input))
	fmt.Println("part 2: ", partTwo(input))

}
func partOne(input string) (answer int) {

	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	strings := ParseInput(input)

	badBrackets := []rune{}
	for _, str := range strings {
		badBracket, err := findBadBracket(str)
		if err == nil {

			badBrackets = append(badBrackets, badBracket)
		}
	}

	for _, v := range badBrackets {
		answer += scores[v]
	}

	return answer
}

func partTwo(input string) (answer int) {

	scoresMap := map[rune]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}

	strings := ParseInput(input)

	stacks := [][]rune{}
	for _, str := range strings {
		_, err := findBadBracket(str)
		if err != nil {
			stack := []rune{}
			for _, c := range str {
				switch c {
				case '(':
					stack = append(stack, ')')
				case '[':
					stack = append(stack, ']')
				case '{':
					stack = append(stack, '}')
				case '<':
					stack = append(stack, '>')
				default:
					if c == stack[len(stack)-1] {
						stack = stack[:len(stack)-1]
					}
				}
			}
			stacks = append(stacks, stack)
		}
	}

	scores := []int{}
	for _, v := range stacks {
		for i, j := 0, len(v)-1; i < j; i, j = i+1, j-1 {
			v[i], v[j] = v[j], v[i]
		}
		score := 0
		for _, char := range v {
			score *= 5
			score += scoresMap[char]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	return scores[(len(scores)-1)/2]

}
