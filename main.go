package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	fmt.Printf("Advent of Code 2017\n\n")

	//fmt.Println(day1a(readInputFile("day1.txt")))
	fmt.Println(day1b(readInputFile("day1.txt")))
}

func day1a(input string) string {
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
	return strconv.Itoa(result)
}

func day1b(input string) string {
	offset := int(len(input) / 2)

	result := 0
	for i := range input[:offset] {
		target := i + offset
		if input[i] == input[target] {
			result += mustBtoi(input[i]) * 2
		}
	}

	return strconv.Itoa(result)
}

func readInputFile(name string) string {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func mustBtoi(b byte) int {
	result, err := strconv.Atoi(string(b))
	if err != nil {
		panic(err)
	}
	return result
}
