package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

// Complete the larrysArray function below.
func larrysArray(A []int) string {

	i := 0
	for sePuedeIterar(A) && !estaOrdenado(A) {
		if seDebeRotar(A, i) {
			rotar(A, i)
		} else {
			i++
			if i > len(A)-3 {
				i = 0
			}
		}
	}

	if estaOrdenado(A) {
		return "YES"
	}

	return "NO"
}

func rotar(a []int, i int) {

	if i >= 0 && i <= len(a)-3 {
		temp := a[i]
		a[i] = a[i+1]
		a[i+1] = a[i+2]
		a[i+2] = temp
	}
}

func seDebeRotar(a []int, i int) bool {

	if i >= 0 && i <= len(a)-3 {
		return a[i] > a[i+1] || a[i] > a[i+2]
	}

	return false
}

func estaOrdenado(a []int) bool {

	// ordenado ascendente, como los elementos van de 1 en 1 también podría usar otro criterio
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}

	return true
}

func sePuedeIterar(a []int) bool {

	for i := 0; i <= len(a)-3; i++ {
		if seDebeRotar(a, i) {
			return true
		}
	}

	return false
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
