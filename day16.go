package main

import (
	"fmt"
	"strings"
)

func day16a(input string) string {
	list := []byte("abcdefghijklmnop")

	moves := strings.Split(input, ",")

	for _, m := range moves {
		list = dance1(list, m)
	}

	return string(list)
}

func day16b(input string) string {

	list := []byte("abcdefghijklmnop")

	moves := strings.Split(input, ",")

	// pattern repeats every 36 iterations. 1,000,000,000 % 36 is 28

	for i := 0; i < 28; i++ {
		for _, m := range moves {
			list = dance1(list, m)
		}
		// find patern repeat
		// key := string(list)
		// if prev[key] {
		// 	fmt.Printf("Found loop at %d - %s\n", i, key)
		// 	break
		// }
		// prev[key] = true
	}

	return string(list)
}

func dance1(list []byte, move string) []byte {
	switch move[0] {
	case 's':
		return spin(list, mustAtoi(move[1:]))
	case 'x':
		sub := move[1:]
		parts := strings.Split(sub, "/")
		return swap(list, mustAtoi(parts[0]), mustAtoi(parts[1]))
	case 'p':
		sub := move[1:]
		parts := strings.Split(sub, "/")
		x, y := findPlaces(list, parts[0][0], parts[1][0])
		return swap(list, x, y)
	}
	panic("Bad things")
}

func spin(list []byte, size int) []byte {
	if size == 0 {
		return list
	}

	result := []byte{}
	result = append(result, list[len(list)-size:]...)
	result = append(result, list[:len(list)-size]...)

	return result
}

func swap(list []byte, x, y int) []byte {
	t := list[x]
	list[x] = list[y]
	list[y] = t
	return list
}

func findPlaces(list []byte, a, b byte) (int, int) {
	found := 0
	aPlace := 0
	bPlace := 0
	for i := range list {
		if found == 2 {
			return aPlace, bPlace
		}

		if list[i] == a {
			aPlace = i
			found++
		} else if list[i] == b {
			bPlace = i
			found++
		}
	}
	if found == 2 {
		return aPlace, bPlace
	}
	panic(fmt.Errorf("Could not find %c %c - %s", a, b, string(list)))
}
