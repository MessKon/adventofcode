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

	sum := 0
	var his, mine string
	lines := strings.Split(string(rawBytes), "\n")
	for _, line := range lines {
		his = string(rune(line[0]))
		mine = string(rune(line[2]))
		// part 1
		// switch mine {
		// case "X":
		// 	sum = sum + 1
		// 	switch his {
		// 	case "A":
		// 		sum = sum + 3
		// 	case "C":
		// 		sum = sum + 6
		// 	}
		// case "Y":
		// 	sum = sum + 2
		// 	switch his {
		// 	case "B":
		// 		sum = sum + 3
		// 	case "A":
		// 		sum = sum + 6
		// 	}
		// case "Z":
		// 	sum = sum + 3
		// 	switch his {
		// 	case "C":
		// 		sum = sum + 3
		// 	case "B":
		// 		sum = sum + 6
		// 	}
		// }

		// part 2
		switch mine {
		case "X":
			switch his {
			case "A":
				sum = sum + 3
			case "B":
				sum = sum + 1
			case "C":
				sum = sum + 2
			}
		case "Y":
			sum = sum + 3
			switch his {
			case "A":
				sum = sum + 1
			case "B":
				sum = sum + 2
			case "C":
				sum = sum + 3
			}
		case "Z":
			sum = sum + 6
			switch his {
			case "A":
				sum = sum + 2
			case "B":
				sum = sum + 3
			case "C":
				sum = sum + 1
			}
		}
	}
	fmt.Println(sum)
}
