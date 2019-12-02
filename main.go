package main

import (
	"adventofcode2019/day1"
	"adventofcode2019/day2"
	"flag"
	"fmt"
	"strconv"
)

var day = flag.Int("day", 0, "number of the day")
var test = flag.Bool("test", false, "if to use test input")
var loglevel = flag.String("loglevel", "info", "log level")
var part = flag.Int("part", 1, "part of the task")

var days = map[string][]func(string){
	"day1": {day1.Solve},
	"day2": {day2.Solve, day2.Solve2},
}

func main() {
	flag.Parse()

	solve, ok := days["day"+strconv.Itoa(*day)]
	if ok {
		solve[*part-1](filename())
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
