package main

func day1a(input string) interface{} {
	matches := []int{}

	for i := range input[:len(input)-1] {
		if input[i] == input[i+1] {
			matches = append(matches, mustBtoi(input[i]))
		}
	}

	if input[0] == input[len(input)-1] {
		matches = append(matches, mustBtoi(input[0]))
	}

	result := 0

	for _, i := range matches {
		result += i
	}
	return result
}

func day1b(input string) interface{} {
	offset := int(len(input) / 2)

	result := 0
	for i := range input[:offset] {
		target := i + offset
		if input[i] == input[target] {
			result += mustBtoi(input[i]) * 2
		}
	}

	return result
}
