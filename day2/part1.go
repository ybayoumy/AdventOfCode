package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var result int64

	ranges := strings.Split(string(data), ",")
	for _, idRange := range ranges {
		ids := strings.Split(idRange, "-")

		firstId, err := strconv.Atoi(ids[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		lastId, err := strconv.Atoi(ids[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		for id := firstId; id <= lastId; id++ {
			if !isValidId(id) {
				result += int64(id)
			}
		}
	}

	fmt.Println("Result:", result)
}

func isValidId(id int) bool {
	idStr := strconv.Itoa(id)

	idLen := utf8.RuneCountInString(idStr)
	if idLen%2 != 0 {
		return true
	}

	if idStr[:idLen/2] == idStr[idLen/2:] {
		return false
	}
	return true
}
