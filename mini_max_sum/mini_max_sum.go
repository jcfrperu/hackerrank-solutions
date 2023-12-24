package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"sort"
)

func solution(lines []string, writer *bufio.Writer) {
	arr := SplitInts(lines[0], " ")
	sort.Ints(arr)

	sum := int64(0)
	for _, item := range arr {
		sum += int64(item)
	}

	sol01 := sum - int64(arr[len(arr)-1])
	sol02 := sum - int64(arr[0])

	fmt.Fprintf(writer, "%d %d", sol01, sol02)
}

func main() {
	//RunWithFile(solution, "mini_max_sum/testcases/001-in.txt")
	RunWithString(solution, "1 2 3 4 5")
}

// https://www.hackerrank.com/challenges/mini-max-sum/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
