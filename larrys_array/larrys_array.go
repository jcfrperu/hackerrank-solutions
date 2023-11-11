package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func main() {
	//reader := bufio.NewReaderSize(strings.NewReader("3\n3\n3 1 2\n4\n1 3 4 2\n5\n1 2 3 5 4"), 1024*1024)
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//stdout, err := os.Stdout, error(nil)
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int(nTemp)

		ATemp := strings.Split(readLine(reader), " ")

		var A []int

		for i := 0; i < int(n); i++ {
			AItemTemp, err := strconv.ParseInt(ATemp[i], 10, 64)
			checkError(err)
			AItem := int(AItemTemp)
			A = append(A, AItem)
		}

		result := larrysArray(A)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
