package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func day12a(input string) string {
	g := buildGraph(input)

	return strconv.Itoa(g.findZeroCount())
}

func day12b(input string) string {
	g := buildGraph(input)

	return strconv.Itoa(g.findGroupCount())
}

func buildGraph(input string) graph {
	g := graph{[]*graphNode{}}

	s := bufio.NewScanner(strings.NewReader(input))
	for s.Scan() {
		line := s.Text()
		con := parsePlumberLine(line)
		g.addNode(con)
	}
	return g
}

type graph struct {
	nodes []*graphNode
}

type graphNode struct {
	id       int
	adjacent []*graphNode
}

type connections struct {
	from int
	to   []int
}

func parsePlumberLine(line string) connections {
	parts := strings.Split(line, " <-> ")
	if len(parts) != 2 {
		panic(fmt.Errorf("Bad line [%s] [%#v]", line, parts))
	}

	to := []int{}

	for _, p := range strings.Split(parts[1], ",") {
		to = append(to, mustAtoi(strings.TrimSpace(p)))
	}

	return connections{from: mustAtoi(parts[0]), to: to}
}

func (g *graph) addNode(con connections) *graphNode {
	node := g.findGraphNode(con.from)
	if node == nil {
		node = &graphNode{id: con.from, adjacent: []*graphNode{}}
		g.nodes = append(g.nodes, node)
	}
	for _, to := range con.to {
		child := g.addNode(connections{from: to})
		node.adjacent = append(node.adjacent, child)
	}

	return node
}

func (g *graph) findGraphNode(id int) *graphNode {
	visited := map[int]bool{}

	for _, n := range g.nodes {
		r := traverseGraph(id, n, visited)
		if r != nil {
			return r
		}
	}

	return nil
}

func (g *graph) findZeroCount() int {

	count := 0
	for _, n := range g.nodes {
		visited := map[int]bool{}
		r := traverseGraph(0, n, visited)
		if r != nil {
			count++
		}
	}

	return count
}

func traverseGraph(id int, node *graphNode, visited map[int]bool) *graphNode {
	if node.id == id {
		return node
	}

	visited[node.id] = true

	for _, n := range node.adjacent {
		if visited[n.id] {
			continue
		}
		r := traverseGraph(id, n, visited)
		if r != nil {
			return r
		}
	}
	return nil
}

func (g *graph) findGroupCount() int {
	groups := map[int]bool{}

	for _, n := range g.nodes {
		visited := map[int]bool{}
		low := findLowestNode(0, n, visited, 3000)
		if !groups[low] {
			groups[low] = true
		}
	}

	return len(groups)
}

func findLowestNode(id int, node *graphNode, visited map[int]bool, lowest int) int {
	if node.id == 0 {
		return 0
	}

	visited[node.id] = true
	currentLow := lowest
	if node.id < currentLow {
		currentLow = node.id
	}

	for _, n := range node.adjacent {
		if visited[n.id] {
			continue
		}
		low := findLowestNode(id, n, visited, currentLow)
		if low == 0 {
			return 0
		}
		if low < currentLow {
			currentLow = low
		}
	}
	return currentLow
}
