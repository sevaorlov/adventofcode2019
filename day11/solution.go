package day11

import (
	"adventofcode2019/input"
	"adventofcode2019/intcode"
	"adventofcode2019/logger"
	"fmt"
	"log"
	"strconv"
)

type Direction struct {
	degrees int
}

func (d *Direction) turn(degrees int) {
	d.degrees += degrees
	if d.degrees < 0 {
		d.degrees += 360
	}
	if d.degrees >= 360 {
		d.degrees -= 360
	}
}

type Robot struct {
	x, y      int
	direction Direction
}

func (r *Robot) turnAndMove(degrees int) {
	r.direction.turn(degrees)
	switch r.direction.degrees {
	case 0:
		r.y--
	case 90:
		r.x++
	case 180:
		r.y++
	case 270:
		r.x--
	}
}

func (r *Robot) paint(area [][]string, color int) {
	if color == 0 {
		area[r.y][r.x] = "."
	} else if color == 1 {
		area[r.y][r.x] = "#"
	}
}

type Hull [][]string

func buildHull(size int) Hull {
	hull := make(Hull, size)
	for i, _ := range hull {
		hull[i] = make([]string, size)
	}
	return hull
}

func (h Hull) print() {
	for _, line := range h {
		for _, item := range line {
			if item == "" {
				fmt.Printf("-")
			}
			if item == "#" {
				fmt.Print("█")
			} else {
				fmt.Printf(item)
			}
		}
		fmt.Println()
	}
}

func Part1(filename string) string {
	instructions := input.ReadInt64CommonedArray(filename)

	a := make([]int64, len(instructions)*100)
	a = append(instructions, a...)

	hull := buildHull(200)
	robot := Robot{x: 100, y: 100}
	hull[robot.y][robot.x] = ""
	countPainted := paintHull(a, hull, robot)

	return strconv.Itoa(countPainted)
}

func Part2(filename string) string {
	instructions := input.ReadInt64CommonedArray(filename)

	a := make([]int64, len(instructions)*100)
	a = append(instructions, a...)

	hull := buildHull(100)
	robot := Robot{x: 30, y: 30}
	hull[robot.y][robot.x] = "#"

	paintHull(a, hull, robot)
	hull.print()

	return ""
}
func paintHull(code []int64, hull Hull, robot Robot) int {
	countPainted := 0
	colorCh := make(chan int)
	moveCh := make(chan int)
	halt := false

	inputFn := func() int64 {
		if hull[robot.y][robot.x] == "#" {
			return 1
		}
		return 0
	}

	inputCh := make(chan int64)
	go func() {
		inputCh <- inputFn()
	}()
	go instructionsForRobot(code, colorCh, moveCh, inputCh)

	for {
		select {
		case color, ok := <-colorCh:
			if !ok {
				halt = true
				break
			}
			if hull[robot.y][robot.x] == "" {
				countPainted++
			}
			inputCh <- inputFn()
			robot.paint(hull, color)
		case move, ok := <-moveCh:
			if !ok {
				halt = true
				break
			}

			switch move {
			case 0:
				robot.turnAndMove(-90)
			case 1:
				robot.turnAndMove(90)
			}

			//fmt.Println("x, y", robot.x, robot.y)
			inputCh <- inputFn()
		}

		if halt {
			break
		}
	}

	return countPainted
}

func instructionsForRobot(code []int64, color chan<- int, move chan<- int, inputCh <-chan int64) {
	index := 0
	colorInstruction := true
	relBase := 0

	var err error
	var output []int64
	var input int64

	for {
		input = <-inputCh
		//fmt.Println("solve", input, index)
		output, index, relBase, err = intcode.Solve(code, []int64{input}, index, relBase, true)
		if err != nil {
			log.Fatal("error from intcode", err.Error())
		}
		if index == -1 {
			break
		}
		if len(output) != 1 {
			log.Fatal("output is expected to have only one value", output)
		}

		if colorInstruction {
			color <- int(output[0])
			logger.Debug("print color", output[0])
		} else {
			move <- int(output[0])
			logger.Debug("move", output[0])
		}

		colorInstruction = !colorInstruction
	}

	close(color)
	close(move)
}
