package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	list := SplitLongs(lines[1], " ")

	sum := int64(0)
	for i := 0; i < len(list); i++ {
		sum += list[i]
	}

	fmt.Fprintf(writer, "%d", sum)
}

func main() {
	//RunWithFile(solution, "a_very_big_sum/testcases/001-in.txt")
	RunWithString(solution, "5\n1000000001 1000000002 1000000003 1000000004 1000000005")
}

// https://www.hackerrank.com/challenges/a-very-big-sum/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
