package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"math/big"
)

// Complete the extraLongFactorials function below.
func extraLongFactorials(n int64) {
	// n must be >= 1
	var fact = big.NewInt(int64(1))

	if n == 1 {
		fact = big.NewInt(n)
	} else {
		i := int64(1)
		for i <= n {
			//fact *= i
			fact = fact.Mul(fact, big.NewInt(i))
			i++
		}
	}

	fmt.Printf("%d", fact)
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
