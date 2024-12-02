package utils

func isStrictlyIncreasing(input []int) bool {
	if len(input) < 2 {
		return true
	}
	for i := 0; i < len(input)-1; i++ {
		if input[i]-input[i+1] >= 0 {
			return false
		}
	}
	return true
}

func isStrictlyDecreasing(input []int) bool {
	if len(input) < 2 {
		return true
	}
	for i := 0; i < len(input)-1; i++ {
		if input[i]-input[i+1] <= 0 {
			return false
		}
	}
	return true
}

func IsStrictlyMonotonic(input []int) bool {
	if len(input) < 2 {
		return true
	}
	switch {
	case (input[0]-input[1] < 0):
		return isStrictlyIncreasing(input)
	case (input[0]-input[1] > 0):
		return isStrictlyDecreasing(input)
	case (input[0]-input[1] == 0):
		return false
	}
	return false
}
