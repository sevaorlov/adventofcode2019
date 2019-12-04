package main

import (
	"adventofcode2019/day1"
	"adventofcode2019/day2"
	"adventofcode2019/day3"
	"adventofcode2019/day4"
	"flag"
	"fmt"
	"strconv"
)

var day = flag.Int("day", 0, "number of the day")
var test = flag.Bool("test", false, "if to use test input")
var part = flag.Int("part", 1, "part of the task")

var days = map[string][]func(string){
	"day1": {day1.Solve},
	"day2": {day2.Solve, day2.Solve2},
	"day3": {day3.Solve, day3.Solve2},
	"day4": {day4.Part1, day4.Part2},
}

func main() {
	flag.Parse()

	funcs, ok := days["day"+strconv.Itoa(*day)]
	if ok {
		funcs[*part-1](filename())
	} else {
		fmt.Printf("solution for a day number %v is not found", *day)
	}
}

func filename() string {
	filename := "input.txt"
	if *test {
		filename = "test.txt"
	}

	return fmt.Sprintf("./day%v/%s", *day, filename)
}
