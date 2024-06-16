package main

import (
	"bufio"
	"fmt"

	. "github.com/jcfrperu/go-competitive-programming"
)

func solution(lines []string, writer *bufio.Writer) {
	n := ParseInt(lines[0])
	list := SplitInts(lines[1], " ")

	quicksort(list, 0, n-1)
	medianIndex := (n + 1) / 2 // constraint: n is odd
	median := list[medianIndex-1]

	fmt.Fprintf(writer, "%d", median)
}

func quicksort(list []int, left int, right int) {
	if left < right {
		p := partition(list, left, right)
		// [left, left + 1, .., p - 1, p] [p + 1, p + 2,..., right - 1,  right]
		quicksort(list, left, p)
		quicksort(list, p+1, right)
	}
}

func partition(list []int, left int, right int) int {
	i := left - 1
	j := right + 1

	pivot := list[(left+right)/2]

	// [left, left + 1, .., p - 1, p] [p + 1, p + 2,..., right - 1,  right]
	for i < j {
		i++
		for i <= right && list[i] < pivot {
			i++
		}
		j--
		for j >= left && list[j] > pivot {
			j--
		}
		if i < j {
			swap(list, i, j)
		}
	}

	return j
}

func swap(list []int, i int, j int) {
	temp := list[i]
	list[i] = list[j]
	list[j] = temp
}

func main() {
	//RunWithFile(solution, "find_the_median/testcases/001-in.txt")
	RunWithString(solution, "7\n0 1 2 4 6 5 3")
}

// https://www.hackerrank.com/challenges/find-the-median/problem
// NOTE: read 'template-submit-code.go' file to submit your code to hackerrank.com
