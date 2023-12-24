package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	arr := SplitInts(lines[1], " ")

	var accumulator = make([]int, len(arr))
	for i := range arr {
		accumulator[arr[i]]++
	}

	var maxValue = arr[0]
	for i := 1; i < len(accumulator); i++ {
		if accumulator[i] > maxValue {
			maxValue = accumulator[i]
		}
	}

	var number = 0
	for i := 1; i < len(accumulator); i++ {
		if accumulator[i] == maxValue {
			number = i
			break
		}
	}

	fmt.Fprintf(writer, "%d", number)
}

func main() {
	//RunWithFile(solution, "migratory_birds/testcases/001-in.txt")
	RunWithString(solution, "11\n1 2 3 4 5 4 3 2 1 3 4")
}

// https://www.hackerrank.com/challenges/migratory-birds/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
