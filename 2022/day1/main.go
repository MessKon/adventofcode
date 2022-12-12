package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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
	var sums []int
	var sum, cur_int = 0, 0
	lines := strings.Split(string(rawBytes), "\n")
	for _, line := range lines {
		if line == "" {
			sums = append(sums, sum)
			sum = 0
			continue
		}
		cur_int, _ = strconv.Atoi(line)
		sum = sum + cur_int
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sums)))
	fmt.Println(sums[0] + sums[1] + sums[2])
}
