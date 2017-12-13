package main

import (
	"bufio"
	"strconv"
	"strings"
)

func day13a(input string) string {
	return strconv.Itoa(checkSeverity(input))
}

func day13b(input string) string {
	depths := map[int]int{}
	maxLayer := 0
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, ": ")
		layer := mustAtoi(parts[0])
		if layer > maxLayer {
			maxLayer = layer
		}
		depths[layer] = mustAtoi(parts[1])
	}

	maxLayer++

	for i := 0; i < 10000000; i++ {
		if checkSeverity2(depths, maxLayer, i) {
			return strconv.Itoa(i)
		}
	}
	panic("Failed to find soltion")
}

func checkSeverity(input string) int {
	depths := map[int]int{}
	maxLayer := 0
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		parts := strings.Split(line, ": ")
		layer := mustAtoi(parts[0])
		if layer > maxLayer {
			maxLayer = layer
		}
		depths[layer] = mustAtoi(parts[1])
	}

	maxLayer++
	layers := make([]int, maxLayer, maxLayer)
	goingUp := make([]bool, maxLayer, maxLayer)
	score := 0

	for i := 0; i < maxLayer; i++ {
		depth, ok := depths[i]
		if ok && layers[i] == 0 {
			score += i * depth
		}
		for j := range layers {
			layerDepth, ok := depths[j]
			if ok {
				if goingUp[j] {
					layers[j]--
					if layers[j] == 0 {
						goingUp[j] = false
					}
				} else {
					layers[j]++
					if layers[j] == layerDepth-1 {
						goingUp[j] = true
					}
				}
			}
		}
	}
	return score
}

func checkSeverity2(depths map[int]int, maxLayer, delay int) bool {
	for l, d := range depths {
		period := (d - 1) * 2
		if (delay+l)%period == 0 {
			return false
		}
	}
	return true
}
