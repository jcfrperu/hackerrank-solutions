package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"math/big"
)

func solution(lines []string, writer *bufio.Writer) {
	n := ParseLong(lines[0])

	// n must be >= 1
	var factorial = big.NewInt(int64(1))

	if n == 1 {
		factorial = big.NewInt(n)
	} else {
		i := int64(1)
		for i <= n {
			//factorial *= i
			factorial = factorial.Mul(factorial, big.NewInt(i))
			i++
		}
	}

	fmt.Fprintf(writer, "%d", factorial)
}

func main() {
	//RunWithFile(solution, "extra_long_factorials/testcases/001-in.txt")
	RunWithString(solution, "25")
}

// https://www.hackerrank.com/challenges/extra-long-factorials/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
