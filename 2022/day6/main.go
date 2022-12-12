package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"os"
)

// set to 4 for part 1
const package_len = 14

func isUnique(buf string) bool {
	var letters [123]int
	for i := 0; i < len(buf); i++ {
		letters[buf[i]]++
		if letters[buf[i]] == 2 {
			return false
		}
	}
	// fmt.Println(buf)
	return true

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

	line := strings.Split(string(rawBytes), "\n")[0]

	buffer := line[:package_len]
	if isUnique(buffer) {
		fmt.Println(package_len)
	}
	for i := 1; i < len(line)-package_len; i++ {
		buffer = line[i : i+package_len]
		if isUnique(buffer) {
			fmt.Println(i + package_len)
			break
		}

	}

}
