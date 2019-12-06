package day6

import (
	"adventofcode2019/input"
	"adventofcode2019/logger"
	"strings"
)

func Part1(filename string) {
	graph := make(map[string][]string)

	input.ReadFile(filename, func(line string) {
		s := strings.Split(line, ")")
		graph[s[0]] = append(graph[s[0]], s[1])
	})

	logger.Debug(graph)

	count := 0
	for key, _ := range graph {
		count += countConnections(graph, key)
	}
	logger.Info(count)
}

func Part2(filename string) {
	graph := make(map[string][]string)

	input.ReadFile(filename, func(line string) {
		s := strings.Split(line, ")")
		graph[s[0]] = append(graph[s[0]], s[1])
		graph[s[1]] = append(graph[s[1]], s[0])
	})

	logger.Debug(graph)

	min := findMin(findPaths(graph, make(map[string]bool), "SAN", "YOU"))

	// subtract SAN and YOU nodes paths
	min -= 2

	logger.Info(min)
}

func findPaths(graph map[string][]string, visited map[string]bool, from string, to string) []int {
	//fmt.Println("findPaths", from, to)

	if len(graph[from]) == 0 || visited[from] {
		return []int{}
	}
	visited[from] = true

	var pathLengths []int

	for _, item := range graph[from] {
		if item == to {
			pathLengths = append(pathLengths, 1)
			break
		}

		for _, v := range findPaths(graph, visited, item, to) {
			pathLengths = append(pathLengths, 1 + v)
		}
	}

	return pathLengths
}

func countConnections(graph map[string][]string, object string) int {
	if len(graph[object]) == 0 {
		return 0
	}

	count := 0
	for _, item := range graph[object] {
		count += 1 + countConnections(graph, item)
	}

	//fmt.Println("countConnections", object, count)

	return count
}

func findMin(vals []int) int {
	min := vals[0]

	for _, item := range vals {
		if item < min {
			min = item
		}
	}

	return min
}
