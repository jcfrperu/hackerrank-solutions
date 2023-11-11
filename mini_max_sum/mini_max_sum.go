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

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int) {

	sort.Ints(arr)

	var sum int64 = 0

	for _, item := range arr {
		sum += int64(item)
	}

	fmt.Printf("%d %d", sum-int64(arr[len(arr)-1]), sum-int64(arr[0]))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)
	//reader := bufio.NewReaderSize(strings.NewReader("5 2 3 4 1"), 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
