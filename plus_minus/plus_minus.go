package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {

	size := float64(len(arr))

	positives := float64(0)
	negatives := float64(0)
	zeros := float64(0)

	for _, item := range arr {
		if item > 0 {
			positives += 1
		} else if item < 0 {
			negatives += 1
		} else {
			zeros += 1
		}
	}

	fmt.Printf("%.6f\n", positives/size)
	fmt.Printf("%.6f\n", negatives/size)
	fmt.Printf("%.6f\n", zeros/size)
}

func main() {

	//localReader := strings.NewReader("6\n-4 3 -9 0 4 1")
	localReader := os.Stdin

	reader := bufio.NewReaderSize(localReader, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
