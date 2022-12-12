package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"os"
)

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

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

	var (
		crates            [9][]string
		splits, to_append []string
		amount, from, to  int
	)
	moves := false

	lines := strings.Split(string(rawBytes), "\n")
	for _, line := range lines {
		if line == "" {
			moves = true
			continue
		}
		if !moves {
			// reading crates
			if string(line[1]) == "1" {
				continue
			}
			for i := 0; i < 9; i++ {
				if string(line[1+i*4]) != " " {
					crates[i] = append([]string{string(line[1+i*4])}, crates[i]...)
				}
			}

		} else {
			// parsing moves
			splits = strings.Split(line, " ")
			amount, _ = strconv.Atoi(splits[1])
			from, _ = strconv.Atoi(splits[3])
			to, _ = strconv.Atoi(splits[5])

			to_append = crates[from-1][len(crates[from-1])-amount:]
			// reverse needed for part 1
			//reverse(to_append)
			crates[to-1] = append(crates[to-1], to_append...)
			crates[from-1] = crates[from-1][:len(crates[from-1])-amount]

		}

	}
	for i := 0; i < 9; i++ {
		fmt.Print(crates[i][len(crates[i])-1])
	}
	fmt.Println()
}
