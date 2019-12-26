package day24

import (
	"adventofcode2019/input"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func playMinute(grid [][]string) [][]string {
	newGrid := make([][]string, len(grid))

	for index, row := range grid {
		newGrid[index] = make([]string, len(row))
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			newGrid[i][j] = newTile(grid, i, j)
			//fmt.Println("put", newGrid[i][j])
		}
	}

	return newGrid
}

func newTile(grid [][]string, i, j int) string {
	bugsCount := 0

	if isBug(grid, i-1, j) {
		bugsCount++
	}

	if isBug(grid, i+1, j) {
		bugsCount++
	}

	if isBug(grid, i, j-1) {
		bugsCount++
	}

	if isBug(grid, i, j+1) {
		bugsCount++
	}

	//fmt.Println(i, j, grid[i][j], bugsCount)
	if grid[i][j] == "#" && bugsCount != 1 {
		return "."
	}

	if grid[i][j] == "." && (bugsCount == 1 || bugsCount == 2) {
		return "#"
	}

	return grid[i][j]
}

func isBug(grid [][]string, i, j int) bool {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid) {
		return false
	}
	return grid[i][j] == "#"
}

func printGrid(grid [][]string) {
	fmt.Println("_")
	for _, row := range grid {
		fmt.Println(row)
	}
}

func biodiversity(grid [][]string) int64 {
	var rate int64

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == "#" {
				pow := i * 5 + j
				rate += int64(math.Pow(float64(2), float64(pow)))
				//fmt.Println(i, j, rate)
			}
		}
	}

	return rate
}

func gridString(grid [][]string) string {
	str := ""

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			str += grid[i][j]
		}
	}

	return str
}

func Part1(filename string) string {
	grid := gridFromFile(filename)

	//printGrid(grid)

	cache := make(map[string]bool)
	cache[gridString(grid)] = true

	for {
		grid = playMinute(grid)

		if cache[gridString(grid)] {
			break
		}

		cache[gridString(grid)] = true
	}

	return strconv.FormatInt(biodiversity(grid), 10)
}

func Part2(filename string) string {
	return ""
}

func gridFromFile(filename string) [][]string {
	var lines []string

	input.ReadFile(filename, func(line string) {
		lines = append(lines, line)
	})

	grid := make([][]string, len(lines))

	for index, line := range lines {
		grid[index] = strings.Split(line, "")
	}

	return grid
}
