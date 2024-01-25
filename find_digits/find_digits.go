package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	t := ParseInt(lines[0])
	for iter := 1; iter <= t; iter++ {
		ns := lines[iter]
		n := ParseLong(ns)

		slice := make([]int64, 0)
		for i := 0; i < len(ns); i++ {
			slice = append(slice, ParseLong(string(ns[i])))
		}

		freq, items, _, _ := Freq(slice, false)

		totalDivisors := 0
		for _, item := range items {
			if item != 0 && n%item == 0 {
				totalDivisors += freq[item]
			}
		}

		fmt.Fprintf(writer, "%d\n", totalDivisors)
	}
}

func main() {
	RunWithFile(solution, "find_digits/testcases/001-in.txt")
	//RunWithString(solution, "2\n12\n1012")
}

// https://www.hackerrank.com/challenges/find-digits/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
