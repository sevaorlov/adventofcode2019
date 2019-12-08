package input

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadFile(filename string, callback func(line string)) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed to read the line. Error: %s\n", err.Error())
		}

		callback(string(line))
	}
}

func ReadStringArraysFromFile(filename string, size int) [][]string {
	data := make([][]string, size)
	var n int

	ReadFile(filename, func(line string) {
		data[n] = strings.Split(line, ",")
		n++
	})

	return data
}

func ReadIntArrayFromFile(filename string) []int64 {
	var a []int64
	var read bool

	ReadFile(filename, func(line string) {
		if read {
			log.Fatalf("more than 1 line in the input file\n")
		}

		strArr := strings.Split(line, ",")

		for _, item := range strArr {
			v, err := strconv.ParseInt(item, 10, 64)
			if err != nil {
				log.Fatalf("failed to parse a line to int. Error: %s\n", err.Error())
			}
			a = append(a, v)
		}
		read = true
	})

	return a
}
