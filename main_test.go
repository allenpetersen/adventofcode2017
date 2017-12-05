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
		{"5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8\n", "18"},
	}

	runTest(t, rows, day2a)
}

func TestDay2b(t *testing.T) {
	rows := []testRow{
		{"5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5\n", "9"},
	}

	runTest(t, rows, day2b)
}

func TestBuildGrid3a(t *testing.T) {
	rows := []struct {
		goal int
		x    int
		y    int
	}{
		{1, 0, 0},
		{2, 1, 0},
		{3, 1, 1},
	}

	for _, r := range rows {
		x, y := buildGrid3a(r.goal)
		if r.x != x && r.y != y {
			t.Errorf("failed! for %d\nexpected [%d,%d]\nactual [%d,%d]\n", r.goal, r.x, r.y, x, y)
		}
	}
}

func TestBuildGrid3b(t *testing.T) {
	rows := []struct {
		goal   int
		result int
	}{
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 4},
		{5, 5},
		{6, 10},
		{23, 806},
	}

	for _, r := range rows {
		result := buildGrid3b(r.goal, 265149)
		if r.result != result {
			t.Errorf("failed! for %d\nexpected [%d]\nactual [%d]\n", r.goal, r.result, result)
		}
	}
}

func TestDay4a(t *testing.T) {
	rows := []struct {
		pass   string
		result bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	}

	for _, r := range rows {
		result := checkPass(r.pass)
		if r.result != result {
			t.Errorf("failed! for %s\nexpected [%v]\nactual [%v]\n", r.pass, r.result, result)
		}
	}
}

func TestDay4b(t *testing.T) {
	rows := []struct {
		pass   string
		result bool
	}{
		{"abcde fghij", true},
		{"abcde xyz ecdab", false},
		{"a ab abc abd abf abj", true},
		{"iiii oiii ooii oooi oooo", true},
		{"oiii ioii iioi iiio", false},
	}

	for _, r := range rows {
		result := checkPass2(r.pass)
		if r.result != result {
			t.Errorf("failed! for %s\nexpected [%v]\nactual [%v]\n", r.pass, r.result, result)
		}
	}
}

func TestDay5a(t *testing.T) {
	rows := []struct {
		instructions []int
		result       int
	}{
		{[]int{0, 3, 0, 1, -3}, 5},
	}

	for _, r := range rows {
		result := runInstructionsA(r.instructions)
		if r.result != result {
			t.Errorf("failed!\nexpected [%v]\nactual [%v]\n", r.result, result)
		}
	}
}

func TestDay5b(t *testing.T) {
	rows := []struct {
		instructions []int
		result       int
	}{
		{[]int{0, 3, 0, 1, -3}, 10},
	}

	for _, r := range rows {
		result := runInstructionsB(r.instructions)
		if r.result != result {
			t.Errorf("failed!\nexpected [%v]\nactual [%v]\n", r.result, result)
		}
	}
}
