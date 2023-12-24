package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	heights := SplitInts(lines[1], " ")
	freqMap := make(map[int]int, len(heights))

	maxVal := heights[0]
	for _, number := range heights {
		// counting freqMap
		freqMap[number] += 1
		// find maximum
		if number > maxVal {
			maxVal = number
		}
	}

	// option 02 using Frequencies method from library
	// freqMap, _, _, maxIndex := Frequencies(heights)
	// maxVal := heights[maxIndex]

	result := freqMap[maxVal]

	fmt.Fprintf(writer, "%d", result)
}

func main() {
	RunWithFile(solution, "birthday_cake_candles/testcases/001-in.txt")
	//RunWithString(solution, "4\n3 2 1 3")
}

// https://www.hackerrank.com/challenges/birthday-cake-candles/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
