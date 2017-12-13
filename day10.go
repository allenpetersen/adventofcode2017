package main

import (
	"encoding/hex"
	"strconv"
	"strings"
)

func day10a(input string) string {
	list := make([]int, 256, 256)

	for i := 0; i < 256; i++ {
		list[i] = i
	}

	parts := strings.Split(input, ",")
	lengths := []int{}
	for _, p := range parts {
		lengths = append(lengths, mustAtoi(p))
	}
	pos := 0
	for i, l := range lengths {
		filpListA(list, l, pos)
		pos += l + i
	}

	return strconv.Itoa(list[0] * list[1])
}

func day10b(input string) string {
	list := make([]int, 256, 256)

	for i := 0; i < 256; i++ {
		list[i] = i
	}

	lengths := []int{}
	for _, b := range []byte(input) {
		lengths = append(lengths, int(b))
	}
	for _, b := range []int{17, 31, 73, 47, 23} {
		lengths = append(lengths, b)
	}

	pos := 0
	skipsize := 0
	for round := 0; round < 64; round++ {
		for _, l := range lengths {
			filpListB(list, l, pos)
			pos += l + skipsize
			skipsize++
		}
	}

	hash := make([]byte, 16, 16)

	for i := range hash {
		var t byte
		offset := i * 16
		for j := offset; j < offset+16; j++ {
			t ^= byte(list[j])
		}
		hash[i] = t
	}

	return hex.EncodeToString(hash)
}

func filpListA(list []int, length, pos int) {
	if length == 1 {
		return
	}
	for i := 0; i < length/2; i++ {
		start := (pos + i) % len(list)
		last := (pos + length - i - 1) % len(list)
		t := list[start]
		list[start] = list[last]
		list[last] = t
	}
}

func filpListB(list []int, length, pos int) {
	for i := 0; i < length/2; i++ {
		start := (pos + i) % len(list)
		last := (pos + length - i - 1) % len(list)
		t := list[start]
		list[start] = list[last]
		list[last] = t
	}
}
