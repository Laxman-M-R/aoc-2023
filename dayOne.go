package main

import (
	"strconv"
	"strings"
)

func dayOnePartOne() int {
	fileLines := readFileLines("dayOneInput.txt")
	totalSum := 0

	for _, line := range fileLines {
		left, right := 0, len(line)-1
		var firstNum, secondNum int
		for {
			if isDigit(string(line[left])) {
				firstNum, _ = strconv.Atoi(string(line[left]))
				break
			}
			left++
		}
		for {
			if isDigit(string(line[right])) {
				secondNum, _ = strconv.Atoi(string(line[right]))
				break
			}
			right--
		}

		totalSum += (firstNum*10 + secondNum)
	}

	return totalSum

}

var numbersWordsMap map[string]int = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func wordStartsWithNumber(s string) int {
	for name, number := range numbersWordsMap {
		if strings.HasPrefix(s, name) {
			return number
		}
	}
	return -1
}

func wordEndsWithNumber(s string) int {
	for name, number := range numbersWordsMap {
		if strings.HasSuffix(s, name) {
			return number
		}
	}
	return -1
}

func dayOnePartTwo() int {
	fileLines := readFileLines("dayOneInput.txt")
	totalSum := 0

	for _, line := range fileLines {
		left, right := 0, len(line)-1
		var firstNum, secondNum int

		for {
			if isDigit(string(line[left])) {
				firstNum, _ = strconv.Atoi(string(line[left]))
				break
			}
			if wordStartNumber := wordStartsWithNumber(line[left:]); wordStartNumber != -1 {
				firstNum = wordStartNumber
				break
			}
			left++
		}
		for {
			if isDigit(string(line[right])) {
				secondNum, _ = strconv.Atoi(string(line[right]))
				break
			}

			if wordEndNumber := wordEndsWithNumber(line[:right+1]); wordEndNumber != -1 {
				secondNum = wordEndNumber
				break
			}
			right--
		}
		totalSum += (firstNum*10 + secondNum)
	}
	return totalSum
}
