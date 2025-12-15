package day09

import (
	"aoc2025/internal/must"
	"math"
	"strings"
)

type Day09 struct {
	nodes map[string]map[string]int
}

func NewDay09(data []string) *Day09 {
	nodes := make(map[string]map[string]int)

	for _, line := range data {
		fields := strings.Fields(line)
		start := fields[0]
		end := fields[2]
		cost := must.ParseInt(fields[4])
		if _, exists := nodes[start]; !exists {
			nodes[start] = make(map[string]int)
		}
		if _, exists := nodes[end]; !exists {
			nodes[end] = make(map[string]int)
		}
		nodes[start][end] = cost
		nodes[end][start] = cost
	}
	return &Day09{nodes: nodes}
}

func (d *Day09) Part1() (int, error) {
	cost := findRoutes("", d.nodes, make(map[string]bool), 0)

	return cost, nil
}

func (d *Day09) Part2() (int, error) {
	cost := findRoutes2("", d.nodes, make(map[string]bool), 0)

	return cost, nil
}

func findRoutes(from string, nodes map[string]map[string]int, visited map[string]bool, cost int) int {
	if len(visited) == len(nodes) {
		return cost
	}
	minCost := math.MaxInt
	for node := range nodes {
		if !visited[node] {
			var moveCost int
			if from != "" {
				moveCost = nodes[from][node]
			}

			visited[node] = true
			bestPath := findRoutes(node, nodes, visited, cost+moveCost)
			minCost = min(minCost, bestPath)
			delete(visited, node)
		}
	}
	return minCost
}

// findRoutes2 is dupl of findRoutes. Could improve by allowing a best cost func or similar
func findRoutes2(from string, nodes map[string]map[string]int, visited map[string]bool, cost int) int {
	if len(visited) == len(nodes) {
		return cost
	}
	maxCost := math.MinInt
	for node := range nodes {
		if !visited[node] {
			var moveCost int
			if from != "" {
				moveCost = nodes[from][node]
			}

			visited[node] = true
			bestPath := findRoutes2(node, nodes, visited, cost+moveCost)
			maxCost = max(maxCost, bestPath)
			delete(visited, node)
		}
	}
	return maxCost
}
