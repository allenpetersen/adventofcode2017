package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/matryer/is"
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

func TestDay6a(t *testing.T) {
	rows := []struct {
		bank   []int
		result int
	}{
		{[]int{0, 2, 7, 0}, 5},
	}

	for _, r := range rows {
		result := balanceBank(r.bank)
		if r.result != result {
			t.Errorf("failed!\nexpected [%v]\nactual [%v]\n", r.result, result)
		}
	}
}

const day7TestData = `pbga (66)
xhth (57)
ebii (61)
havc (66)
ktlj (57)
fwft (72) -> ktlj, cntj, xhth
qoyq (66)
padx (45) -> pbga, havc, qoyq
tknk (41) -> ugml, padx, fwft
jptl (61)
ugml (68) -> gyxo, ebii, jptl
gyxo (61)
cntj (57)`

func TestDay7b(t *testing.T) {
	result := day7b(day7TestData)

	if result != "60" {
		t.Errorf("Failed to find Disc Balance\nexpected: %s\nactual: %s", "60", result)
	}
}

const day8TestData = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

func TestDay8a(t *testing.T) {
	result := day8a(day8TestData)

	if result != "1" {
		t.Errorf("Failed to find Disc Balance\nexpected: %s\nactual: %s", "1", result)
	}
}

func TestFlip(t *testing.T) {
	rows := []struct {
		input    []int
		length   int
		pos      int
		expected []int
	}{
		{[]int{0, 1, 2, 3, 4}, 3, 0, []int{2, 1, 0, 3, 4}},
		{[]int{2, 1, 0, 3, 4}, 4, 3, []int{4, 3, 0, 1, 2}},
		{[]int{4, 3, 0, 1, 2}, 1, 3, []int{4, 3, 0, 1, 2}},
		{[]int{4, 3, 0, 1, 2}, 5, 1, []int{3, 4, 2, 1, 0}},
	}

	for _, r := range rows {
		filpListA(r.input, r.length, r.pos)
		if !cmp.Equal(r.expected, r.input) {
			t.Errorf("Failed\nexpected: %v\n  actual: %v\n", r.expected, r.input)
		}
	}
}

func TestHashB(t *testing.T) {
	rows := []struct {
		input    string
		expected string
	}{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}

	for _, r := range rows {
		result := day10b(r.input)
		if result != r.expected {
			t.Errorf("Failed\nexpected: %v\n  actual: %v\n", r.expected, result)
		}
	}
}

func TestScoreGroups(t *testing.T) {
	rows := []struct {
		input string
		count int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<{},{},{{}}>}", 1},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}

	for _, r := range rows {
		result := scoreGroups(r.input)
		if result != r.count {
			t.Errorf("Failed %s\nexpected: %v\n  actual: %v\n", r.input, r.count, result)
		}
	}
}

func TestDay11a(t *testing.T) {
	rows := []struct {
		moves    string
		distance int
	}{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 0},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3},
	}

	for _, r := range rows {
		distance := hexgridDistance(r.moves)
		if distance != r.distance {
			t.Errorf("Failed %s\nexpected: %v\n  actual: %v\n", r.moves, r.distance, distance)
		}
	}
}

func TestDay13a(t *testing.T) {
	input := `0: 3
1: 2
4: 4
6: 4`

	result := checkSeverity(input)
	if result != 24 {
		t.Errorf("Failed\nexpected: %d\n  actual: %d\n", 24, result)
	}
}

func TestDay13b(t *testing.T) {
	input := `0: 3
1: 2
4: 4
6: 4`

	result := day13b(input)
	if result != "10" {
		t.Errorf("Failed\nexpected: %s\n  actual: %s\n", "10", result)
	}
}

func TestDay16a(t *testing.T) {
	is := is.New(t)
	rows := []struct {
		start string
		move  string
		end   string
	}{
		{"abcde", "s4", "bcdea"},
		{"abcdefghijklmnop", "s14", "cdefghijklmnopab"},
		// {[]byte{'a', 'b', 'c', 'd', 'e'}, "s3", []byte{'c', 'd', 'e', 'a', 'b'}},
		// {[]byte{'a', 'b', 'c', 'd', 'e'}, "s1", []byte{'e', 'a', 'b', 'c', 'd'}},
		// {[]byte{'e', 'a', 'b', 'c', 'd'}, "x3/4", []byte{'e', 'a', 'b', 'd', 'c'}},
		// {[]byte{'e', 'a', 'b', 'd', 'c'}, "pe/b", []byte{'b', 'a', 'e', 'd', 'c'}},
	}

	for _, r := range rows {
		result := dance1([]byte(r.start), r.move)
		is.Equal(string(result), r.end)
	}
}

func TestDay17a(t *testing.T) {
	is := is.New(t)
	is.Equal(day17a(3), 638)
}

func TestDay18b(t *testing.T) {
	input := `snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d`

	expected := "2"
	actual := day18b(input)
	if expected != actual {
		t.Errorf("Failed\nexpected: %s\nactual: %s", expected, actual)
	}
}
