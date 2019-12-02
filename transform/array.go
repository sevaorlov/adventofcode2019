package transform

import (
	"log"
	"strconv"
	"strings"
)

func Int64ArrayFromLine(line string) []int64 {
	return stringsToInt64Array(strings.Split(line, ","))
}

func stringsToInt64Array(a []string) []int64 {
	var err error

	b := make([]int64, len(a))
	for index, item := range a {
		b[index], err = strconv.ParseInt(item, 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	return b
}
