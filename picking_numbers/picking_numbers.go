package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {

	list := SplitInts(lines[1], " ")
	slist, freqMap := FrequenciesInt(list)

	// out(writer, fmt.Sprintf("%v\n", list))
	// out(writer, fmt.Sprintf("%v\n", freqMap))
	// out(writer, fmt.Sprintf("%v\n", slist))

	// solution
	var counter int
	maxCounter := 0
	for i := 0; i < len(slist)-1; i++ {
		// diferences = 0
		counter = freqMap[slist[i]]
		if counter > maxCounter {
			maxCounter = counter
		}
		// diferences = 1
		if slist[i+1]-slist[i] <= 1 {
			counter = freqMap[slist[i]] + freqMap[slist[i+1]]
			if counter > maxCounter {
				maxCounter = counter
			}
		}
	}

	// diferences = 0 (last item)
	lastIndex := len(slist) - 1
	counter = freqMap[slist[lastIndex]]
	if counter > maxCounter {
		maxCounter = counter
	}

	Out(writer, fmt.Sprintf("%d", maxCounter))
}

func main() {
	//RunWithFile(solution, "PROBLEM/testcases/001.input")
	RunWithString(solution, "1\n2 3")
}

// https://www.hackerrank.com/challenges/URL_PROBLEM/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
