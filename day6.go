package main

import (
	"fmt"
	"strconv"
)

func day6a(input string) string {
	return strconv.Itoa(balanceBank(splitLine(input)))
}

func day6b(input string) string {
	return strconv.Itoa(balanceBank2(splitLine(input)))
}

func balanceBank(bank []int) int {
	states := map[string]struct{}{}

	state := fmt.Sprintf("%+v", bank)
	states[state] = struct{}{}

	count := 1
	current := findLargestSlot(bank)
	items := bank[current]
	bank[current] = 0
	current++

	for {
		state = fmt.Sprintf("%+v", bank)
		_, ok := states[state]
		if ok {
			return count
		}

		if current > len(bank)-1 {
			current = 0
		}

		items--

		if items == -1 {
			states[state] = struct{}{}
			current = findLargestSlot(bank)
			items = bank[current]
			bank[current] = 0
			current++
			count++
		} else {
			bank[current]++
			current++
		}
	}
}

func balanceBank2(bank []int) int {
	states := map[string]int{}

	state := fmt.Sprintf("%+v", bank)
	states[state] = 0

	count := 1
	current := findLargestSlot(bank)
	items := bank[current]
	bank[current] = 0
	current++

	for {
		state = fmt.Sprintf("%+v", bank)
		loop, ok := states[state]
		if ok {
			return count - loop
		}

		if current > len(bank)-1 {
			current = 0
		}

		items--

		if items == -1 {
			states[state] = count
			current = findLargestSlot(bank)
			items = bank[current]
			bank[current] = 0
			current++
			count++
		} else {
			bank[current]++
			current++
		}
	}
}

func findLargestSlot(bank []int) int {
	max := bank[0]
	index := 0
	for i, s := range bank[1:] {
		if s > max {
			max = s
			index = i + 1
		}
	}
	return index
}
