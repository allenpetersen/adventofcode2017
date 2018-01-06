package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type instructionD18 struct {
	action string
	reg    string
	value  string
}

func (inst instructionD18) getValue(reg map[string]int) int {
	// snd instruction from day18b
	if len(inst.value) == 0 {
		v := inst.reg
		if len(v) == 1 && v[0] >= 'a' && v[0] <= 'z' {
			return reg[v]
		}
		return mustAtoi(v)
	}

	// other values
	if len(inst.value) == 1 && inst.value[0] >= 'a' && inst.value[0] <= 'z' {
		return reg[inst.value]
	}
	return mustAtoi(inst.value)
}

func day18a(input string) string {
	regs := map[string]int{}

	program := []instructionD18{}

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		inst := parseD18InstructionLine(line)
		program = append(program, inst)
	}

	lastSound := 0
	i := 0
	for {
		inst := program[i]
		//fmt.Printf("i:%d %v\n", i, inst)
		switch inst.action {
		case "snd":
			lastSound = regs[inst.reg]
		case "set":
			// X = Y
			regs[inst.reg] = inst.getValue(regs)
		case "add":
			// X += Y
			regs[inst.reg] += inst.getValue(regs)
		case "mul":
			// X *= Y
			regs[inst.reg] *= inst.getValue(regs)
		case "mod":
			// X %= Y
			regs[inst.reg] %= inst.getValue(regs)
		case "rcv":
			if regs[inst.reg] != 0 {
				return strconv.Itoa(lastSound)
			}
			fmt.Printf("%s %d\n", inst, regs[inst.reg])

		case "jgz":
			if regs[inst.reg] > 0 {
				i += inst.getValue(regs)
				continue
			}
		default:
			panic(fmt.Errorf("Unknown command [%v]", inst))
		}
		i++
	}
}

func day18b(input string) string {
	p0toP1 := make(chan int, 100)
	p1toP0 := make(chan int, 100)
	p0 := newProgram(input, 0, p0toP1, p1toP0)
	p1 := newProgram(input, 1, p1toP0, p0toP1)

	p0Tick := p0.start()
	p1Tick := p1.start()

	timeout := time.Tick(1 * time.Second)
	p1Count := 0
	alive := false

	for {
		select {
		case <-timeout:
			if !alive {
				return strconv.Itoa(p1Count)
			}
			alive = false
		case <-p0Tick:
			alive = true
		case <-p1Tick:
			alive = true
			p1Count++
		}
	}
}

func parseD18InstructionLine(line string) instructionD18 {
	parts := strings.Split(line, " ")

	if len(parts) != 2 && len(parts) != 3 {
		panic(fmt.Errorf("Failed to parse line %s - l:%d", line, len(parts)))
	}

	inst := instructionD18{
		action: parts[0],
		reg:    parts[1],
	}

	if len(parts) == 3 {
		inst.value = parts[2]
	}

	return inst
}

type program struct {
	id           int
	regs         map[string]int
	instructions []instructionD18
	snd          chan int
	rcv          chan int
}

func newProgram(input string, id int, snd, rcv chan int) *program {
	insts := []instructionD18{}

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		inst := parseD18InstructionLine(line)
		insts = append(insts, inst)
	}

	return &program{
		id:           id,
		regs:         map[string]int{"p": id},
		instructions: insts,
		snd:          snd,
		rcv:          rcv,
	}
}

func (p *program) getValue(value string) int {
	if len(value) == 1 && value[0] >= 'a' && value[0] <= 'z' {
		return p.regs[value]
	}
	return mustAtoi(value)
}

func (p *program) start() chan struct{} {
	tick := make(chan struct{})

	go func() {
		i := 0
		for {
			inst := p.instructions[i]
			// fmt.Printf("p: %d i:%d %v %v\n", p.id, i, inst, p.regs)
			switch inst.action {
			case "snd":
				//fmt.Printf("p: %d i:%d snd %v\n", p.id, i, inst)
				p.snd <- p.getValue(inst.reg)
				tick <- struct{}{}
			case "set":
				// X = Y
				p.regs[inst.reg] = p.getValue(inst.value)
			case "add":
				// X += Y
				p.regs[inst.reg] += p.getValue(inst.value)
			case "mul":
				// X *= Y
				p.regs[inst.reg] *= p.getValue(inst.value)
			case "mod":
				// X %= Y
				p.regs[inst.reg] %= p.getValue(inst.value)
			case "rcv":
				//fmt.Printf("p: %d i:%d rcv %v\n", p.id, i, inst)
				p.regs[inst.reg] = <-p.rcv
			case "jgz":
				v := p.getValue(inst.reg)
				if v > 0 {
					i += p.getValue(inst.value)
					continue
				}
			default:
				panic(fmt.Errorf("Unknown command [%v]", inst))
			}
			i++
		}
	}()

	return tick
}
