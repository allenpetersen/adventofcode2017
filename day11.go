package main

import (
	"strconv"
	"strings"
)

func day11a(input string) string {
	return strconv.Itoa(hexgridDistance(input))
}

func day11b(input string) string {
	return strconv.Itoa(hexgridDistanceB(input))
}

func hexgridDistance(input string) int {
	var x, y, z int
	moves := strings.Split(input, ",")
	for _, m := range moves {
		switch m {
		case "n":
			y++
			z--
		case "ne":
			x++
			z--
		case "se":
			x++
			y--
		case "s":
			y--
			z++
		case "sw":
			x--
			z++
		case "nw":
			x--
			y++
		}
	}

	return maxAbsInt(x, y, z)
}

func hexgridDistanceB(input string) int {
	var x, y, z int
	moves := strings.Split(input, ",")
	max := 0
	for _, m := range moves {
		switch m {
		case "n":
			y++
			z--
		case "ne":
			x++
			z--
		case "se":
			x++
			y--
		case "s":
			y--
			z++
		case "sw":
			x--
			z++
		case "nw":
			x--
			y++
		}
		current := maxAbsInt(x, y, z)
		if current > max {
			max = current
		}
	}

	return max
}

func maxAbsInt(values ...int) int {
	max := 0
	for _, i := range values {
		if i < 0 {
			i *= -1
		}
		if i > max {
			max = i
		}
	}
	return max
}
