package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func convertStringArrayToInt(stringArray []string) []int {
	var intArray []int
	for _, element := range stringArray {
		if element != "" {
			elementInt, _ := strconv.Atoi(element)
			intArray = append(intArray, elementInt)
		}
	}
	return intArray
}

// iterative binary search
func binarySearch(array []int, element int) bool {
	found := false
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] == element {
			found = true
			break
		}
		if array[mid] > element {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return found
}

// recursive binary search.
// func binarySearch(array []int, target int, low int, high int) bool {
// 	//specify condition to end the recursion
// 	if high < low {
// 		return false
// 	}
// 	// Define our middle index
// 	mid := int((low + high) / 2)
// 	if array[mid] > target {
// 		return binarySearch(array, target, low, mid)
// 	} else if array[mid] < target {
// 		return binarySearch(array, target, mid+1, high)
// 	} else {
// 		return true
// 	}
// }

func dayFourPartOne() int {
	cards := readFileLines("dayFourInput.txt")
	matchingNumbersSum := 0
	for _, line := range cards {
		_, scratchCardsString, _ := strings.Cut(line, ": ")
		scratchCards := strings.Split(scratchCardsString, " | ")
		winningNumberStrings, existingNumberStrings := strings.Split(scratchCards[0], " "), strings.Split(scratchCards[1], " ")
		winningNumbers := convertStringArrayToInt(winningNumberStrings)
		existingNumbers := convertStringArrayToInt(existingNumberStrings)
		sort.Ints(existingNumbers)
		matchingNumbersCount := 0
		for _, winningNumber := range winningNumbers {

			if binarySearch(existingNumbers, winningNumber) {
				matchingNumbersCount++
			}
		}
		if matchingNumbersCount != 0 {
			matchingNumbersSum += int(math.Pow(2, float64(matchingNumbersCount-1)))
		}
	}

	return matchingNumbersSum
}

func dayFourPartTwo() int {
	cards := readFileLines("dayFourInput.txt")
	cardsCount := make(map[int]int)
	for i := 0; i < len(cards); i++ {
		cardsCount[i] = 1
	}
	for cardIndex, line := range cards {
		_, scratchCardsString, _ := strings.Cut(line, ": ")
		scratchCards := strings.Split(scratchCardsString, " | ")
		winningNumberStrings, existingNumberStrings := strings.Split(scratchCards[0], " "), strings.Split(scratchCards[1], " ")
		winningNumbers := convertStringArrayToInt(winningNumberStrings)
		existingNumbers := convertStringArrayToInt(existingNumberStrings)
		sort.Ints(existingNumbers)
		matchingNumbersCount := 0
		for _, winningNumber := range winningNumbers {
			if binarySearch(existingNumbers, winningNumber) {
				matchingNumbersCount++
			}
		}
		for n := 0; n < cardsCount[cardIndex]; n++ {
			for cardNum := cardIndex + 1; cardNum <= cardIndex+matchingNumbersCount; cardNum++ {
				if cardNum >= len(cards) {
					break
				}
				cardsCount[cardNum]++
			}
		}

	}
	totalCards := 0
	for _, count := range cardsCount {
		totalCards += count
	}
	return totalCards
}
