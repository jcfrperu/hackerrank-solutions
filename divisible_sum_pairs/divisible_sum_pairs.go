package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	k := SplitGetIntAt(lines[0], " ", 1)
	ar := SplitInts(lines[1], " ")
	result := 0

	for i := 0; i < len(ar)-1; i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar[i]+ar[j])%k == 0 {
				result++
			}
		}
	}

	fmt.Fprintf(writer, "%d", result)
}

func main() {
	//RunWithFile(solution, "divisible_sum_pairs/testcases/001-in.txt")
	RunWithString(solution, "6 3\n1 3 2 6 1 2 ")
}

// https://www.hackerrank.com/challenges/divisible-sum-pairs/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
