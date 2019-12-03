package input

import (
	"bufio"
	"io"
	"log"
	"os"
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
			log.Fatal(err.Error())
		}

		callback(string(line))
	}
}

func ReadArrayFromFile(filename string, size int) [][]string {
	data := make([][]string, size)
	var n int

	ReadFile(filename, func(line string) {
		data[n] = strings.Split(line, ",")
		n++
	})

	return data
}
