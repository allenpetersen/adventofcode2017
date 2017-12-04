package main

import (
	"bufio"
	"sort"
	"strconv"
	"strings"
)

func day4a(input string) string {
	count := 0
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		if checkPass(s.Text()) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func day4b(input string) string {
	count := 0
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		if checkPass2(s.Text()) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func checkPass(pass string) bool {
	m := map[string]struct{}{}

	words := strings.Split(pass, " ")
	for _, w := range words {
		_, ok := m[w]
		if ok {
			return false
		}
		m[w] = struct{}{}
	}

	return true
}

func checkPass2(pass string) bool {
	m := map[string]struct{}{}

	words := strings.Split(pass, " ")
	for _, w := range words {
		sorted := sortPass(w)
		_, ok := m[sorted]
		if ok {
			return false
		}
		m[sorted] = struct{}{}
	}

	return true
}

func sortPass(input string) string {
	b := []byte(input)
	sort.Slice(b, func(i, j int) bool { return b[i] > b[j] })
	return string(b)
}
