package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	var list = SplitInts(lines[1], " ")
	var _, freqMap = FrequenciesInt(list)

	var pairs = 0
	for _, value := range freqMap {
		if value%2 == 0 {
			pairs += value / 2
		} else {
			pairs += (value - 1) / 2
		}
	}

	Out(writer, fmt.Sprintf("%d", pairs))
}

func main() {
	//RunWithFile(solution, "PROBLEM/testcases/001.input")
	RunWithString(solution, "1\n2 3")
}

// https://www.hackerrank.com/challenges/URL_PROBLEM/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
