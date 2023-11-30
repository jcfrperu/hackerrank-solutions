package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"sort"
)

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int) {

	sort.Ints(arr)

	var sum int64 = 0

	for _, item := range arr {
		sum += int64(item)
	}

	fmt.Printf("%d %d", sum-int64(arr[len(arr)-1]), sum-int64(arr[0]))
}

func solution(lines []string, writer *bufio.Writer) {

	Out(writer, fmt.Sprintf("%d", 4))
}

func main() {
	//RunWithFile(solution, "PROBLEM/testcases/001.input")
	RunWithString(solution, "1\n2 3")
}

// https://www.hackerrank.com/challenges/URL_PROBLEM/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
