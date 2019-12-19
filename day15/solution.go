package day15

import (
	"adventofcode2019/input"
	"adventofcode2019/intcode"
	"fmt"
	"log"
	"strconv"
)

const (
	north = 1
	south = 2
	west  = 3
	east  = 4
)

type Point struct {
	x, y, dist int
}

type Area [][]string

func buildArea(size int) Area {
	area := make(Area, size)
	for i, _ := range area {
		area[i] = make([]string, size)
	}
	return area
}

func (a Area) print() {
	for _, line := range a {
		var lineStr string
		notEmptyCount := 0
		for _, item := range line {
			if item == "" {
				lineStr += "#"
			} else {
				notEmptyCount++
				lineStr += item
			}
		}
		if notEmptyCount > 0 {
			fmt.Printf("%s\r", lineStr)
		}
	}
}

type Robot struct {
	program intcode.OneStepProgram
}

func (r *Robot) explore(x, y int, area Area) {
	for _, direction := range []int{north, south, west, east} {
		x1, y1 := newCoord(x, y, direction)
		if x1 < 0 || y1 < 0 {
			break
		}
		if area[y1][x1] != "" {
			continue
		}
		res := r.program.Run(int64(direction))
		if r.program.Finished {
			return
		}

		switch res {
		case 0:
			area[y1][x1] = "#"
		case 1:
			area[y1][x1] = "."
			r.explore(x1, y1, area)
			r.program.Run(int64(reverseDirection(direction)))
		case 2:
			area[y1][x1] = "O"
			r.explore(x1, y1, area)
			r.program.Run(int64(reverseDirection(direction)))
		}
	}
}

func (r *Robot) findPathToMask(x, y int, area Area) []int {
	var paths []int

	visited := make([][]bool, len(area))
	for i, _ := range visited {
		visited[i] = make([]bool, len(area))
	}

	queueCount := 1

	searchCh := make(chan Point)
	go func() {
		searchCh <- Point{x, y, 0}
	}()

	for point := range searchCh {
		queueCount--
		for _, direction := range []int{north, south, west, east} {

			x1, y1 := newCoord(point.x, point.y, direction)
			//fmt.Println("go to ", x1, y1)

			if visited[y1][x1] || x1 < 0 || y1 < 0 || x1 >= len(area) || y1 >= len(area) {
				//fmt.Println("cont")
				continue
			}

			if area[y1][x1] == "#" {
				//fmt.Println("wall")
				continue
			}

			if area[y1][x1] == "O" {
				paths = append(paths, point.dist+1)
				break
			}

			visited[y1][x1] = true
			queueCount++

			go func(x, y, dist int) {
				searchCh <- Point{x, y, dist}
			}(x1, y1, point.dist+1)
		}

		if queueCount == 0 {
			close(searchCh)
		}
	}

	return paths
}

func newCoord(x, y int, direction int) (int, int) {
	switch direction {
	case north:
		return x, y - 1
	case south:
		return x, y + 1
	case west:
		return x - 1, y
	case east:
		return x + 1, y
	}

	log.Fatal("wrong direction passed", direction)
	return -1, -1
}

func reverseDirection(direction int) int {
	switch direction {
	case 1:
		return 2
	case 2:
		return 1
	case 3:
		return 4
	case 4:
		return 3
	}

	log.Fatal("wrong direction passed to reverseDirection", direction)
	return -1
}

func inputInstructions(filename string) []int64 {
	instructions := input.ReadInt64CommonedArray(filename)

	code := make([]int64, len(instructions)*100)
	code = append(instructions, code...)

	return code
}

func min(values []int) int {
	min := values[0]

	for _, item := range values {
		if item < min {
			min = item
		}
	}

	return min
}

func Part1(filename string) string {
	code := inputInstructions(filename)
	program := intcode.NewOneStepProgram(code)

	robotX := 50
	robotY := 50

	area := buildArea(200)
	area[robotY][robotX] = "D"

	robot := Robot{program: program}
	robot.explore(robotX, robotY, area)

	area.print()

	paths := robot.findPathToMask(robotX, robotY, area)
	return strconv.Itoa(min(paths))
}

func Part2(filename string) string {
	return ""
}
