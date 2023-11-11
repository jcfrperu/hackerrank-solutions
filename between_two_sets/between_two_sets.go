package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
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

func main() {
	//reader := bufio.NewReaderSize(strings.NewReader("2 2\n3 4\n24 48"), 16*1024*1024)
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	//stdout, err := os.Stdout, error(nil)
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int(nTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int(mTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	brrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var brr []int

	for i := 0; i < int(m); i++ {
		brrItemTemp, err := strconv.ParseInt(brrTemp[i], 10, 64)
		checkError(err)
		brrItem := int(brrItemTemp)
		brr = append(brr, brrItem)
	}

	total := getTotalX(arr, brr)

	fmt.Fprintf(writer, "%d\n", total)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
