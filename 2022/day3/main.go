package main

import (
	"fmt"
	"io/ioutil"
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

	sum, count := 0, 0
	var chars string

	lines := strings.Split(string(rawBytes), "\n")
	for _, line := range lines {
		// part 1
		// left = line[:(len(line) / 2)]
		// right = line[len(line)/2:]

		// for _, ch := range left {

		// 	if strings.Contains(right, string(ch)) {
		// 		if ch >= 97 && ch <= 122 {
		// 			sum = sum + int(ch) - 96
		// 		}
		// 		if ch >= 65 && ch <= 90 {
		// 			sum = sum + int(ch) - 38
		// 		}
		// 		break
		// 	}
		// }

		// part 2
		switch count {
		case 0:
			chars = line
		case 1, 2:
			for _, ch := range chars {
				if !strings.Contains(line, string(ch)) {
					chars = strings.ReplaceAll(chars, string(ch), "")
				}
			}
		}

		count = count + 1
		if count == 3 {
			if chars[0] >= 97 && chars[0] <= 122 {
				sum = sum + int(chars[0]) - 96
			}
			if chars[0] >= 65 && chars[0] <= 90 {
				sum = sum + int(chars[0]) - 38
			}

			count = 0
		}

	}
	fmt.Println(sum)

}
