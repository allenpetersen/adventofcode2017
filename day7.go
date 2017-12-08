package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func day7b(input string) string {
	return strconv.Itoa(findDiscBalance(input))
}

func findDiscBalance(input string) int {
	head := buildTree(input)

	return findBalanceError(head)
}

func findBalanceError(node *diskNode) int {
	node.totalWeight = node.weight
	weights := map[int]int{}
	for _, c := range node.children {
		correctWeight := findBalanceError(c)
		if correctWeight != 0 {
			return correctWeight
		}

		node.totalWeight += c.totalWeight
		weights[c.totalWeight]++
	}

	if len(weights) <= 1 {
		return 0
	}

	fmt.Printf("Bad weights on node %s %v\n", node.name, weights)

	goodWeight := 0
	badWeight := 0

	for k, v := range weights {
		if v == 1 {
			badWeight = k
		} else {
			goodWeight = k
		}
	}

	for _, c := range node.children {
		if c.totalWeight == badWeight {
			return c.weight + goodWeight - badWeight
		}
	}
	panic("Couldn't find bad weight")
}

func buildTree(input string) *diskNode {
	s := bufio.NewScanner(strings.NewReader(input))

	nodes := []*diskNode{}

	for s.Scan() {
		line := s.Text()
		di, err := parseDiskLine(line)
		if err != nil {
			panic(err)
		}
		dn := findNode(nodes, di.name)
		if dn == nil {
			dn = &diskNode{name: di.name, weight: di.weight}
			nodes = append(nodes, dn)
		} else {
			dn.weight = di.weight
		}

		for _, childName := range di.children {
			cn := findNode(nodes, childName)
			if cn == nil {
				cn = &diskNode{name: childName}
			} else {
				nodes = removeNode(nodes, cn.name)
			}
			dn.children = append(dn.children, cn)
		}
	}

	if len(nodes) != 1 {
		panic(fmt.Errorf("There wasn't a single head [%v]", nodes))
	}

	fmt.Printf("Head node %s\n", nodes[0].name)
	return nodes[0]
}

func findNode(nodes []*diskNode, name string) *diskNode {
	for _, n := range nodes {
		r := findNodeDepth(n, name)
		if r != nil {
			return r
		}
	}
	return nil
}

func findNodeDepth(node *diskNode, name string) *diskNode {
	if node == nil {
		return nil
	}

	if node.name == name {
		return node
	}

	for _, c := range node.children {
		r := findNodeDepth(c, name)
		if r != nil {
			return r
		}
	}
	return nil
}

func removeNode(nodes []*diskNode, name string) []*diskNode {
	var index int
	for i, n := range nodes {
		if n.name == name {
			index = i
			break
		}
	}

	nodes[len(nodes)-1], nodes[index] = nodes[index], nodes[len(nodes)-1]
	return nodes[:len(nodes)-1]
}

type diskNode struct {
	name        string
	weight      int
	totalWeight int
	children    []*diskNode
}

type diskInfo struct {
	name     string
	weight   int
	children []string
}

var patternDiskLine = regexp.MustCompile("^([[:alpha:]]+) \\(([[:digit:]]+)\\)( -> (.*))?$")

func parseDiskLine(line string) (diskInfo, error) {
	m := patternDiskLine.FindStringSubmatch(line)
	if m == nil || len(m) != 5 {
		return diskInfo{}, fmt.Errorf("Failed to parse line [%s]", line)
	}

	result := diskInfo{
		name:   m[1],
		weight: mustAtoi(m[2]),
	}
	if m[4] != "" {
		children := strings.Split(m[4], ",")
		for _, c := range children {
			result.children = append(result.children, strings.TrimSpace(c))
		}
	}

	return result, nil
}
