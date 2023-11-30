package main

import (
	"bufio"
	"fmt"
	. "github.com/jcfrperu/go-competitive-programming"
)

// Complete the birthdayCakeCandles function below.
func birthdayCakeCandles(ar []int32) int32 {

	mapa := make(map[int32]int32, len(ar))

	max := ar[0]

	//fmt.Printf("arreglo=%+v\n", ar )
	//fmt.Printf("mapa=%+v\n", mapa )

	for _, number := range ar {

		//fmt.Printf("item=%d\n", number)

		// if exists
		if mapa[number] != 0 {
			//fmt.Printf("entrando al != en number=%d\n", mapa[number] )
			mapa[number] += 1
		} else { // if doesnot
			mapa[number] = 1 // in go isnt necessary
		}

		if number > max {
			max = number
		}
		//fmt.Printf("mapa=%+v\n", mapa )
	}

	//fmt.Printf("mapa final=%+v\n", mapa )
	//fmt.Printf("max final=%d\n", max )

	return mapa[max]
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
