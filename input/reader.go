package input

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	MainFileName = "input.txt"
	TestFileName = "test.txt"
)

func FilePath(day int, test bool) string {
	filename := MainFileName
	if test {
		filename = TestFileName
	}

	return fmt.Sprintf("./day%v/%s", day, filename)
}

func ReadFile(filepath string, callback func(line string)) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed to read the line. Error: %s\n", err.Error())
		}

		callback(strings.TrimSuffix(line, "\n"))
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

func ReadInt64CommonedArray(filename string) []int64 {
	var a []int64

	line := ReadSingleLine(filename)
	strArr := strings.Split(line, ",")

	for _, item := range strArr {
		v, err := strconv.ParseInt(item, 10, 64)
		if err != nil {
			log.Fatalf("failed to parse a line item to int. Item: %v. Error: %s\n", item, err.Error())
		}
		a = append(a, v)
	}

	return a
}

func ReadSingleLine(filename string) string {
	var read bool
	var singleLine string

	ReadFile(filename, func(line string) {
		if read {
			log.Fatalf("more than 1 line in the input file\n")
		}

		singleLine = line

		read = true
	})

	return singleLine
}

func ReadIntSpacedArray(filename string) []int {
	line := ReadSingleLine(filename)

	var a []int

	strA := strings.Split(line, "")

	for _, strInt := range strA {
		i, err := strconv.Atoi(strInt)
		if err != nil {
			log.Fatalf("couldnt parse string to int %v. Error: %s \n", strInt, err.Error())
		}
		a = append(a, i)
	}

	return a
}

func ReadIntArray(filename string, separator string) []int {
	line := ReadSingleLine(filename)

	var elements []int

	for _, item := range strings.Split(line, separator) {
		i, err := strconv.Atoi(item)
		if err != nil {
			log.Fatalf("couldnt parse string to int %v. Error: %s \n", item, err.Error())
		}
		elements = append(elements, i)
	}

	return elements
}
