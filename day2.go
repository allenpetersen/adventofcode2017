package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func day2a(input string) string {
	s := bufio.NewScanner(strings.NewReader(input))
	result := 0
	for s.Scan() {
		line := s.Text()
		nums := splitLine(line)
		lowest := nums[0]
		highest := lowest
		for _, i := range nums {
			if i > highest {
				highest = i
			}
			if i < lowest {
				lowest = i
			}
		}
		result += highest - lowest
	}

	return strconv.Itoa(result)
}

func day2b(input string) string {
	s := bufio.NewScanner(strings.NewReader(input))
	result := 0
	for s.Scan() {
		line := s.Text()
		nums := splitLine(line)
		fmt.Println(nums)
		for i, v1 := range nums[:len(nums)-1] {
			for _, v2 := range nums[i+1:] {
				if v1%v2 == 0 {
					result += v1 / v2
				} else if v2%v1 == 0 {
					result += v2 / v1
				}
			}
		}
	}

	return strconv.Itoa(result)
}
