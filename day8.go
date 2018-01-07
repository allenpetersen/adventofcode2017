package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type instruction struct {
	reg       string
	action    string
	ammount   int
	a         string
	b         int
	condition string
}

func day8a(input string) string {
	regs := map[string]int{}

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		inst := parseInstructionLine(line)

		if testCondition(regs, inst) {
			switch inst.action {
			case "inc":
				regs[inst.reg] += inst.ammount
			case "dec":
				regs[inst.reg] -= inst.ammount
			}
		}
	}

	return strconv.Itoa(getMaxReg(regs))
}

func day8b(input string) string {
	regs := map[string]int{}

	highest := 0
	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		inst := parseInstructionLine(line)

		if testCondition(regs, inst) {
			switch inst.action {
			case "inc":
				regs[inst.reg] += inst.ammount
			case "dec":
				regs[inst.reg] -= inst.ammount
			}
		}
		h := getMaxReg(regs)
		if h > highest {
			highest = h
		}
	}

	return strconv.Itoa(highest)
}

func testCondition(regs map[string]int, inst instruction) bool {
	r := regs[inst.a]

	switch inst.condition {
	case ">":
		return r > inst.b
	case "<":
		return r < inst.b
	case ">=":
		return r >= inst.b
	case "<=":
		return r <= inst.b
	case "==":
		return r == inst.b
	case "!=":
		return r != inst.b
	default:
		panic(fmt.Errorf("Unknown condition %s", inst.condition))
	}
}

func parseInstructionLine(line string) instruction {
	parts := strings.Split(line, " ")

	if len(parts) != 7 {
		panic(fmt.Errorf("Failed to parse line %s", line))
	}

	return instruction{
		reg:       parts[0],
		action:    parts[1],
		ammount:   mustAtoi(parts[2]),
		a:         parts[4],
		condition: parts[5],
		b:         mustAtoi(parts[6]),
	}
}

func getMaxReg(regs map[string]int) int {
	first := true
	var max int

	for _, v := range regs {
		if first {
			max = v
			first = false
		} else {
			if v > max {
				max = v
			}
		}
	}
	return max
}
