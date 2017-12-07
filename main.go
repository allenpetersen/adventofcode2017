package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Printf("Advent of Code 2017\n\n")
	start := time.Now()

	// runDay("day 1a", "day1.txt", day1a)
	// runDay("day 1b", "day1.txt", day1b)
	// runDay("day 2a", "day2.txt", day2a)
	// runDay("day 2b", "day2.txt", day2b)
	// fmt.Println(buildGrid3a(265149))
	// fmt.Println(buildGrid3b(265149, 265149))
	// runDay("day 4a", "day4.txt", day4a)
	// runDay("day 4b", "day4.txt", day4b)
	// runDay("day 5a", "day5.txt", day5a)
	// runDay("day 5b", "day5.txt", day5b)
	// runDay("day 6a", "day6.txt", day6a)
	runDay("day 6b", "day6.txt", day6b)

	fmt.Printf("\nDone in %s\n", time.Since(start))
}

func runDay(name, filename string, fn func(string) string) {
	start := time.Now()
	result := fn(readInputFile(filename))
	fmt.Printf("%s: %v in %v\n", name, result, time.Since(start))
}

func splitLine(line string) []int {
	parts := strings.Split(line, "\t")
	result := []int{}
	for _, p := range parts {
		result = append(result, mustAtoi(p))
	}
	return result
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

func mustAtoi(s string) int {
	result, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return result
}
