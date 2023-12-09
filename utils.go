package main

import (
	"bufio"
	"fmt"
	"os"
)

func readFileLines(filename string) []string {
	f, err := os.Open(filename)

	if err != nil {
		fmt.Println("err", err)
		panic(err)
	}

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	f.Close()

	return fileLines
}

func isDigit(c string) bool {
	return c >= "0" && c <= "9"
}
