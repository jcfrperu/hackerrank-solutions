package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
	"sort"
)

func getTotalX(as []int, bs []int) int {

	sort.Ints(as)
	sort.Ints(bs)

	var maxA = as[len(as)-1]
	var candidatosMap = make(map[int]int, len(as))

	//fmt.Printf("as = %v\n", as)
	//fmt.Printf("bs = %v\n", bs)

	// busca1`rs una lista de candidatos
	for _, b := range bs {
		i := 1
		for b/i >= maxA {
			if b%i == 0 {
				candidatosMap[b/i] = b / i
			}
			i++
		}
	}

	var candidatos = make([]int, 0, len(as))
	for key := range candidatosMap {
		candidatos = append(candidatos, key)
	}

	sort.Ints(candidatos)
	//fmt.Printf("candidatos = %v\n", candidatos)

	// quitar los canditados que no cumplen la condicion candidatos % a(i) == 0
	var cantidatosC1 = make([]int, 0, len(as))
	for _, candidato := range candidatos {
		var cumpleTodos = true
		for _, a := range as {
			if candidato%a != 0 {
				cumpleTodos = false
				break
			}
		}
		if cumpleTodos {
			cantidatosC1 = append(cantidatosC1, candidato)
		}
	}

	//fmt.Printf("cantidatosC1 = %v\n", cantidatosC1)

	// quitar los canditados que no cumplen la condicion b(i) % canditato == 0
	var cantidatosC2 = make([]int, 0, len(as))
	for _, candidato := range cantidatosC1 {
		var cumpleTodos = true
		for _, b := range bs {
			if b%candidato != 0 {
				cumpleTodos = false
				break
			}
		}
		if cumpleTodos {
			cantidatosC2 = append(cantidatosC2, candidato)
		}
	}
	//fmt.Printf("cantidatosC2 = %v\n", cantidatosC2)

	return len(cantidatosC2)
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
