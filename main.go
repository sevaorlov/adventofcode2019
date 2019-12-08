package main

import (
	"adventofcode2019/day1"
	"adventofcode2019/day2"
	"adventofcode2019/day3"
	"adventofcode2019/day4"
	"adventofcode2019/day5"
	"adventofcode2019/day6"
	"adventofcode2019/day7"
	"adventofcode2019/day8"
	"adventofcode2019/logger"
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
	"day3": {day3.Solve, day3.Solve2},
	"day4": {day4.Part1, day4.Part2},
	"day5": {day5.Part1, day5.Part2},
	"day6": {day6.Part1, day6.Part2},
	"day7": {day7.Part1, day7.Part2},
	"day8": {day8.Part1, day8.Part2},
}

func main() {
	flag.Parse()

	logger.Init(*loglevel)

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
