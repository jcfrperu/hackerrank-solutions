package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
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

func main() {
	//reader := bufio.NewReaderSize(strings.NewReader("4\n3 2 1 3"), 1024*1024)
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//stdout, err := os.Stdout, error(nil)
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	arCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	arTemp := strings.Split(readLine(reader), " ")

	var ar []int32

	for i := 0; i < int(arCount); i++ {
		arItemTemp, err := strconv.ParseInt(arTemp[i], 10, 64)
		checkError(err)
		arItem := int32(arItemTemp)
		ar = append(ar, arItem)
	}

	result := birthdayCakeCandles(ar)

	fmt.Fprintf(writer, "%d\n", result)

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
