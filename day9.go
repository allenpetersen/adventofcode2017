package main

import (
	"strconv"
)

func day9a(input string) string {
	return strconv.Itoa(scoreGroups(input))
}

func day9b(input string) string {
	return strconv.Itoa(countGarbage(input))
}

func scoreGroups(input string) int {
	score := 0
	inGarbage := false
	excapeNextChar := false
	activeGroups := 0

	for _, b := range []byte(input) {
		if inGarbage {
			if excapeNextChar {
				excapeNextChar = false
				continue
			}
			if b == '>' {
				inGarbage = false
				continue
			}
			if b == '!' {
				excapeNextChar = true
				continue
			}
		} else {
			if b == '{' {
				activeGroups++
				continue
			}
			if b == '<' {
				inGarbage = true
				continue
			}
			if b == '}' {
				score += activeGroups
				activeGroups--
				continue
			}
		}
	}
	return score
}

func countGarbage(input string) int {
	count := 0
	inGarbage := false
	excapeNextChar := false

	for _, b := range []byte(input) {
		if inGarbage {
			if excapeNextChar {
				excapeNextChar = false
				continue
			}
			if b == '>' {
				inGarbage = false
				continue
			}
			if b == '!' {
				excapeNextChar = true
				continue
			}
			count++
		} else {
			if b == '<' {
				inGarbage = true
				continue
			}
		}
	}
	return count
}
