package main

import (
	"testing"
)

type testRow struct {
	input    string
	expected string
}

func runTest(t *testing.T, rows []testRow, fn func(string) string) {
	for _, r := range rows {
		result := fn(r.input)
		if r.expected != result {
			t.Errorf("failed!\nexpected [%s]\nactual [%s]\n", r.expected, result)
		}
	}
}

func TestDay1a(t *testing.T) {
	rows := []testRow{
		{"1122", "3"},
		{"1111", "4"},
		{"1234", "0"},
		{"91212129", "9"},
	}

	runTest(t, rows, day1a)
}

func TestDay1b(t *testing.T) {
	rows := []testRow{
		{"1212", "6"},
		{"1221", "0"},
		{"123425", "4"},
		{"123123", "12"},
		{"12131415", "4"},
	}

	runTest(t, rows, day1b)
}

func TestDay2a(t *testing.T) {
	rows := []testRow{
		{"", ""},
	}

	runTest(t, rows, day2a)
}
