package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"os"
)

func main() {
	fileIO, err := os.OpenFile("input.txt", os.O_RDWR, 0600)
	if err != nil {
		panic(err)
	}
	defer fileIO.Close()
	rawBytes, err := ioutil.ReadAll(fileIO)
	if err != nil {
		panic(err)
	}

	sum := 0
	var lstart, lend, rstart, rend int
	var splits, left, right []string
	lines := strings.Split(string(rawBytes), "\n")
	for _, line := range lines {

		splits = strings.Split(line, ",")
		left = strings.Split(splits[0], "-")
		right = strings.Split(splits[1], "-")

		lstart, _ = strconv.Atoi(left[0])
		lend, _ = strconv.Atoi(left[1])
		rstart, _ = strconv.Atoi(right[0])
		rend, _ = strconv.Atoi(right[1])

		// part 1
		// if (lstart >= rstart && lend <= rend) || (rstart >= lstart && rend <= lend) {
		// 	sum = sum + 1
		// }

		// part 2
		if !(lend < rstart || rend < lstart) {
			sum = sum + 1
		}
	}
	fmt.Println(sum)
}
