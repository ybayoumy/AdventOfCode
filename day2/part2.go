package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

func main() {
	data, err := os.ReadFile("input.txt")
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
			if !isValidId2(id) {
				result += int64(id)
			}
		}
	}

	fmt.Println("Result:", result)
}

func isValidId2(id int) bool {
	idStr := strconv.Itoa(id)

	idLen := utf8.RuneCountInString(idStr)
	for i := 2; i <= idLen; i++ {
		parts := getEqualSplits(idStr, i)
		if len(parts) == 0 {
			continue
		}

		if allEquals(parts) {
			return false
		}
	}

	return true
}

func allEquals(s []string) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] != s[i] {
			return false
		}
	}
	return true
}

func getEqualSplits(s string, n int) []string {
	var result []string

	sLen := utf8.RuneCountInString(s)
	if sLen%n != 0 {
		return result // empty
	}

	jump := sLen / n
	for i := 0; i < sLen; i += jump {
		result = append(result, s[i:i+jump])
	}

	return result
}
