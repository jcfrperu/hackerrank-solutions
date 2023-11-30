package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {

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

	fmt.Printf("%.6f\n", positives/size)
	fmt.Printf("%.6f\n", negatives/size)
	fmt.Printf("%.6f\n", zeros/size)
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
