package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

// Complete the migratoryBirds function below.
func migratoryBirds(arr []int) int {

	var acum = []int{0, 0, 0, 0, 0, 0}

	for i := range arr {
		acum[arr[i]]++
	}

	var max = arr[0]
	for i := 1; i < len(acum); i++ {
		if acum[i] > max {
			max = acum[i]
		}
	}

	var number = 0
	for i := 1; i < len(acum); i++ {
		if acum[i] == max {
			number = i
			break
		}
	}

	return number
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
