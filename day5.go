package main

import (
	"bufio"
	"strings"
)

func day5a(input string) interface{} {
	ins := []int{}
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		ins = append(ins, mustAtoi(s.Text()))
	}
	return runInstructionsA(ins)
}

func day5b(input string) interface{} {
	ins := []int{}
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		ins = append(ins, mustAtoi(s.Text()))
	}
	return runInstructionsB(ins)
}

func runInstructionsA(instructs []int) int {
	count := 0

	for i := 0; i < len(instructs); {
		count++
		t := instructs[i]
		instructs[i]++
		i += t
	}
	return count
}

func runInstructionsB(instructs []int) int {
	count := 0
	for i := 0; i < len(instructs); {
		count++
		t := instructs[i]
		if t >= 3 {
			instructs[i]--
		} else {
			instructs[i]++
		}
		i += t
	}
	return count
}
