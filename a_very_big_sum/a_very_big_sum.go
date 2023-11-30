package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	var n = ParseInt(lines[0])
	var list = SplitLongs(lines[1], " ")

	var sum int64 = 0
	for i := 0; i < n; i++ {
		sum += list[i]
	}

	Out(writer, fmt.Sprintf("%d", sum))
}

func main() {
	//RunWithFile(solution, "a_very_big_sum/testcases/002.input")
	RunWithString(solution, "5\n1000000001 1000000002 1000000003 1000000004 1000000005")
}

// https://www.hackerrank.com/challenges/a-very-big-sum/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
