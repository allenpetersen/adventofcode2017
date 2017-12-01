package main

import "testing"

func TestDay1a(t *testing.T) {
	rows := []struct {
		input    string
		expected string
	}{
		{"1122", "3"},
		{"1111", "4"},
		{"1234", "0"},
		{"91212129", "9"},
	}

	for _, r := range rows {
		result := day1a(r.input)
		if r.expected != result {
			t.Errorf("failed!\nexpected [%s]\nactual [%s]\n", r.expected, result)
		}
	}
}

func TestDay1b(t *testing.T) {
	rows := []struct {
		input    string
		expected string
	}{
		{"1212", "6"},
		{"1221", "0"},
		{"123425", "4"},
		{"123123", "12"},
		{"12131415", "4"},
	}

	for _, r := range rows {
		result := day1b(r.input)
		if r.expected != result {
			t.Errorf("failed! input [%s]\nexpected [%s]\nactual [%s]\n", r.input, r.expected, result)
		}
	}
}
