package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// Complete the extraLongFactorials function below.
func extraLongFactorials(n int64) {
	// n must be >= 1
	var fact = big.NewInt(int64(1))

	if n == 1 {
		fact = big.NewInt(n)
	} else {
		i := int64(1)
		for i <= n {
			//fact *= i
			fact = fact.Mul(fact, big.NewInt(i))
			i++
		}
	}

	fmt.Printf("%d", fact)
}

func main() {
	//reader := bufio.NewReaderSize(strings.NewReader("25"), 1024*1024)
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := nTemp

	extraLongFactorials(n)
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
