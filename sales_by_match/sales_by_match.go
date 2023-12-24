package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	var list = SplitInts(lines[1], " ")

	freqMap, _, _, _ := Frequencies(list)

	// or
	//freqMap := make(map[int]int, len(list))
	//for i := range list {
	//	freqMap[list[i]]++
	//}

	var pairs = 0
	for _, value := range freqMap {
		if value%2 == 0 {
			pairs += value / 2
		} else {
			pairs += (value - 1) / 2
		}
	}

	fmt.Fprintf(writer, "%d", pairs)
}

func main() {
	RunWithFile(solution, "sales_by_match/testcases/001-in.txt")
	//RunWithString(solution, "9\n10 20 20 10 10 30 50 10 20")
}

// https://www.hackerrank.com/challenges/sock-merchant/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
