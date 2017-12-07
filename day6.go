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
			fmt.Printf("c: %d, i: %d s: %s count: %d\n", current, items, state, count)
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
		fmt.Printf("c: %d, i: %d s: %s count: %d\n", current, items, state, count)
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
			fmt.Printf("c: %d, i: %d s: %s count: %d\n", current, items, state, count)
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
		fmt.Printf("c: %d, i: %d s: %s count: %d\n", current, items, state, count)
	}
}

// func redistBank(bank []int) {
// 	current := findLargestSlot(bank)
// 	items := bank[current]
// 	current++
// 	fmt.Printf("redistBank start c: %d, i: %d\n", current, items)
// 	for items > 0 {
// 		if items > 10 {
// 			return
// 		}
// 		fmt.Printf("c: %d, i: %d\n", current, items)
// 		if current > len(bank)-1 {
// 			current = 0
// 		}

// 		bank[current]++
// 		current++
// 		items--
// 	}
// 	fmt.Printf("redistBank done")
// }

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
