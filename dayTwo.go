package main

import (
	"strconv"
	"strings"
)

func getNumColorBalls(s string) map[string]int {
	colorNumsString := strings.Split(s, ", ")
	colorBallsNum := make(map[string]int)
	for _, colorNum := range colorNumsString {
		numString, color, _ := strings.Cut(colorNum, " ")
		num, _ := strconv.Atoi(numString)
		colorBallsNum[color] = num
	}
	return colorBallsNum
}

func dayTwoPartOne(numRedBalls, numGreenBalls, numBlueBalls int) int {
	fileLines := readFileLines("dayTwoInput.txt")
	gameIdsTotal := 0
	for _, line := range fileLines {
		gameId, gameSetsString, _ := strings.Cut(line, ": ")
		gameSets := strings.Split(gameSetsString, "; ")
		isPossible := true

	out:
		for _, gameSet := range gameSets {
			colorBallsNum := getNumColorBalls(gameSet)
			var totalNumOfBallsInGameSet int
			for color, num := range colorBallsNum {
				if color == "red" && num > numRedBalls || color == "green" && num > numGreenBalls || color == "blue" && num > numBlueBalls {
					isPossible = false
					break out
				}
				totalNumOfBallsInGameSet += num
			}
			if totalNumOfBallsInGameSet > (numRedBalls + numGreenBalls + numBlueBalls) {
				isPossible = false
				break out
			}

		}
		if isPossible {
			_, idstr, _ := strings.Cut(gameId, " ")
			id, _ := strconv.Atoi(idstr)
			gameIdsTotal += id
		}

	}
	return gameIdsTotal
}

func dayTwoPartTwo() int {
	fileLines := readFileLines("dayTwoInput.txt")
	powersSum := 0
	for _, line := range fileLines {
		_, gameSetsString, _ := strings.Cut(line, ": ")
		gameSets := strings.Split(gameSetsString, "; ")
		maxRedBalls := 0
		maxGreenBalls := 0
		maxBlueBalls := 0

		for _, gameSet := range gameSets {
			colorBallsNum := getNumColorBalls(gameSet)

			for color, num := range colorBallsNum {
				if color == "red" {
					if num > maxRedBalls {
						maxRedBalls = num
					}
				}
				if color == "green" {
					if num > maxGreenBalls {
						maxGreenBalls = num
					}
				}
				if color == "blue" {
					if num > maxBlueBalls {
						maxBlueBalls = num
					}
				}
			}
		}
		powersSum += (maxRedBalls * maxGreenBalls * maxBlueBalls)
	}
	return powersSum
}
