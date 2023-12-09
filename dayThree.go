package main

import (
	"strconv"
)

func isNotPeriod(rowIndex, col int, rows []string) bool {

	if rowIndex < 0 || rowIndex >= len(rows) || col < 0 || col >= len(rows[rowIndex]) {
		return false
	}
	return string(rows[rowIndex][col]) != "."
}

func dayThreePartOne() int {
	fileLines := readFileLines("dayThreeInput.txt")
	partNumsSum := 0
	for rowId, row := range fileLines {
		for colId := 0; colId < len(row); colId++ {
			if !isDigit(string(row[colId])) {
				continue
			}

			digitStartIndex := colId
			digitEndIndex := colId
			numStr := string(row[digitStartIndex])

			for j := digitStartIndex + 1; j < len(row); j++ {
				if !isDigit(string(row[j])) {
					break
				}
				numStr += string(row[j])
				digitEndIndex = j
			}

			colId = digitEndIndex

			// left and right check
			if isNotPeriod(rowId, digitStartIndex-1, fileLines) || isNotPeriod(rowId, digitEndIndex+1, fileLines) {
				num, _ := strconv.Atoi(numStr)
				partNumsSum += num
				continue
			}

			// top-down check.
			for col := digitStartIndex - 1; col <= digitEndIndex+1; col++ {
				if isNotPeriod(rowId-1, col, fileLines) || isNotPeriod(rowId+1, col, fileLines) {
					num, _ := strconv.Atoi(numStr)
					partNumsSum += num
					break
				}
			}
		}

	}
	return partNumsSum
}

func isNumAdjacentToStar(rowIndex, col int, rows []string) bool {

	if rowIndex < 0 || rowIndex >= len(rows) || col < 0 || col >= len(rows[rowIndex]) {
		return false
	}
	return string(rows[rowIndex][col]) == "*"
}

type Position struct {
	Row    int
	Column int
}

type Gears struct {
	Count   int
	Numbers []int
}

func dayThreePartTwo() int {
	fileLines := readFileLines("dayThreeInput.txt")
	gears := make(map[Position]Gears)
	// gearNumsCount := make(map[Position]int)
	// gearNums := make(map[Position][]int)
	gearRatiosSum := 0
	for rowId, row := range fileLines {
		for colId := 0; colId < len(row); colId++ {
			if !isDigit(string(row[colId])) {
				continue
			}

			digitStartIndex := colId
			digitEndIndex := colId
			numStr := string(row[digitStartIndex])

			for j := digitStartIndex + 1; j < len(row); j++ {
				if !isDigit(string(row[j])) {
					break
				}
				numStr += string(row[j])
				digitEndIndex = j
			}

			colId = digitEndIndex
			num, _ := strconv.Atoi(numStr)

			//left.
			if isNumAdjacentToStar(rowId, digitStartIndex-1, fileLines) {
				position := Position{Row: rowId, Column: digitStartIndex - 1}
				gearPosition := gears[position]
				gearPosition.Count++
				gearPosition.Numbers = append(gearPosition.Numbers, num)
				gears[position] = gearPosition
				// gearNumsCount[position]++
				// gearNums[position] = append(gearNums[position], num)
				continue
			}

			// right.
			if isNumAdjacentToStar(rowId, digitEndIndex+1, fileLines) {
				position := Position{Row: rowId, Column: digitEndIndex + 1}
				gearPosition := gears[position]
				gearPosition.Count++
				gearPosition.Numbers = append(gearPosition.Numbers, num)
				gears[position] = gearPosition
				continue
			}

			// top-down.
			// TODO: Move the code blocks inside each if condns outside.
			for col := digitStartIndex - 1; col <= digitEndIndex+1; col++ {
				if isNumAdjacentToStar(rowId-1, col, fileLines) {
					position := Position{Row: rowId - 1, Column: col}
					gearPosition := gears[position]
					gearPosition.Count++
					gearPosition.Numbers = append(gearPosition.Numbers, num)
					gears[position] = gearPosition
					// gearNumsCount[position]++
					// gearNums[position] = append(gearNums[position], num)
					break
				} else if isNumAdjacentToStar(rowId+1, col, fileLines) {
					position := Position{Row: rowId + 1, Column: col}
					gearPosition := gears[position]
					gearPosition.Count++
					gearPosition.Numbers = append(gearPosition.Numbers, num)
					gears[position] = gearPosition
					// gearNumsCount[position]++
					// gearNums[position] = append(gearNums[position], num)
					break
				}

			}
		}

	}
	for _, gearData := range gears {
		if gearData.Count != 2 {
			continue
		}
		// nums := gearNums[position]
		gearRatiosSum += gearData.Numbers[0] * gearData.Numbers[1]
	}
	return gearRatiosSum
}
