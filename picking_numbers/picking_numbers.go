package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	list := SplitInts(lines[1], " ")
	freqMap, sList, _, _ := Freq(list, false)

	// solution
	var counter int
	maxCounter := 0
	for i := 0; i < len(sList)-1; i++ {
		// differences = 0
		counter = freqMap[sList[i]]
		if counter > maxCounter {
			maxCounter = counter
		}
		// differences = 1
		if sList[i+1]-sList[i] <= 1 {
			counter = freqMap[sList[i]] + freqMap[sList[i+1]]
			if counter > maxCounter {
				maxCounter = counter
			}
		}
	}

	// differences = 0 (last item)
	lastIndex := len(sList) - 1
	counter = freqMap[sList[lastIndex]]
	if counter > maxCounter {
		maxCounter = counter
	}

	fmt.Fprintf(writer, "%d", maxCounter)
}

func main() {
	RunWithFile(solution, "picking_numbers/testcases/001-in.txt")
	//RunWithString(solution, "6\n4 6 5 3 3 1")
}

// https://www.hackerrank.com/challenges/picking-numbers/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
