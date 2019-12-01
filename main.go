package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)
import "adventofcode2019/day1"

var day = flag.Int("day", 0, "number of the day")
var test = flag.Bool("test", false, "if to use test input")

var days = map[string]func(string) {
	"day1": day1.Solve,
}

func main() {
	flag.Parse()

	solve, ok := days["day" + strconv.Itoa(*day)]
	if ok {
		solve(filename())
	} else {
		log.Fatalf("solution for a day number %s is not found", *day)
	}
}

func filename() string {
	filename := "input.txt"
	if *test {
		filename = "test.txt"
	}

	return fmt.Sprintf("./day%v/%s", *day, filename)
}
