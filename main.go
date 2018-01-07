package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Printf("Advent of Code 2017\n\n")
	start := time.Now()

	runDay("day 1a", "day1.txt", day1a)
	runDay("day 1b", "day1.txt", day1b)
	runDay("day 2a", "day2.txt", day2a)
	runDay("day 2b", "day2.txt", day2b)
	fmt.Println(buildGrid3a(265149))
	fmt.Println(buildGrid3b(265149, 265149))
	runDay("day 4a", "day4.txt", day4a)
	runDay("day 4b", "day4.txt", day4b)
	runDay("day 5a", "day5.txt", day5a)
	runDay("day 5b", "day5.txt", day5b)
	runDay("day 6a", "day6.txt", day6a)
	runDay("day 6b", "day6.txt", day6b)
	runDay("day 7b", "day7.txt", day7b)
	runDay("day 8a", "day8.txt", day8a)
	runDay("day 8b", "day8.txt", day8b)
	runDay("day 9a", "day9.txt", day9a)
	runDay("day 9b", "day9.txt", day9b)
	runDay("day 10a", "day10.txt", day10a)
	runDay("day 10b", "day10.txt", day10b)
	runDay("day 11a", "day11.txt", day11a)
	runDay("day 11b", "day11.txt", day11b)
	runDay("day 12a", "day12.txt", day12a)
	runDay("day 12b", "day12.txt", day12b)
	runDay("day 13a", "day13.txt", day13a)
	runDay("day 13b", "day13.txt", day13b)
	fmt.Printf("day 14a %d\n", day14a("hfdlxzhv"))
	fmt.Printf("day 14b %d\n", day14b("hfdlxzhv"))
	runDay("day 16a", "day16.txt", day16a)
	runDay("day 16b", "day16.txt", day16b)
	fmt.Printf("Day 17a %d\n", day17a(329))
	//	fmt.Printf("Day 17b %d\n", day17b(329))

	runDay("day 18a", "day18.txt", day18a)
	runDay("day 18b", "day18.txt", day18b)

	//	modTest()

	fmt.Printf("\nDone in %s\n", time.Since(start))
}

func modTest() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i:%d 1:%d 2:%d 3:%d 4:%d 5:%d 6:%d 7:%d\n", i, i%1, i%2, i%3, i%4, i%5, i%6, i%7)
	}
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
	b, err := ioutil.ReadFile(path.Join("data", name))
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
