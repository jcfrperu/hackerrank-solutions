package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	arr := SplitInts(lines[1], " ")
	size := float64(len(arr))

	positives := float64(0)
	negatives := float64(0)
	zeros := float64(0)

	for _, item := range arr {
		if item > 0 {
			positives += 1
		} else if item < 0 {
			negatives += 1
		} else {
			zeros += 1
		}
	}

	fmt.Fprintf(writer, "%.6f\n", positives/size)
	fmt.Fprintf(writer, "%.6f\n", negatives/size)
	fmt.Fprintf(writer, "%.6f\n", zeros/size)
}

func main() {
	RunWithFile(solution, "plus_minus/testcases/001-in.txt")
	//RunWithString(solution, "6\n-4 3 -9 0 4 1")
}

// https://www.hackerrank.com/challenges/plus-minus/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
