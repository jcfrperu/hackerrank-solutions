package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	for i := 1; i <= len(lines)-2; i = i + 2 {
		threshold := SplitGetIntAt(lines[i], " ", 1)
		arrivals := SplitInts(lines[i+1], " ")

		cancelled := "YES"
		arrivalsOnTime := 0
		for _, arrival := range arrivals {
			if arrival <= 0 {
				arrivalsOnTime++
			}
		}
		if arrivalsOnTime >= threshold {
			cancelled = "NO"
		}

		fmt.Fprintf(writer, "%s\n", cancelled)
	}
}

func main() {
	RunWithFile(solution, "angry_professor/testcases/001-in.txt")
	//RunWithFile(solution, "angry_professor/testcases/003-in.txt")
	//RunWithString(solution, "1\n100 95\n-58 37 -15 -17 98 80 -6 35 23 -52 9 -25 56 -71 80 85 -89 44 88 -70 41 79 18 38 50 84 -88 -35 24 98 13 -47 62 10 57 40 -86 -99 89 -46 -84 50 43 68 -92 76 -53 -63 -19 -69 -9 -80 -32 0 20 13 29 28 50 44 18 4 47 -39 -93 -48 70 50 -59 29 -48 88 -6 54 69 3 -70 -11 -26 -40 -64 26 16 41 -78 59 65 -4 -67 32 8 -76 94 -51 42 44 80 93 -98 28")
}

// https://www.hackerrank.com/challenges/angry-professor/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
