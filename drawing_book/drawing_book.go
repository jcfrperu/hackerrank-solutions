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

/* ****************************************************************** */
/* **************************   SOLUTION   ************************** */
/* ****************************************************************** */
func solution(lines []string, writer *bufio.Writer) {
    var n = parseInt(lines[0])
    var p = parseInt(lines[1])

    var turnsLeft = (p - 0) / 2
    var turnsRight = (n - p) / 2
    
    if n % 2 == 0 && p % 2 != 0 {
        turnsRight += 1
    }

    var turns = minInt(turnsLeft, turnsRight)
    
    out(writer, fmt.Sprintf("%d", turns))
}

/* ******************************************************************** */
/* ***************   UTILS: print / splits / read input   ************* */
/* ******************************************************************** */
// main method
func main() {
    file, writer := getWriter()
    defer closeFile(file)
    defer flushWriter(writer)
    lines := getLines()
    solution(lines, writer)
}

// how to get input, from a string or default server input
func getInput() io.Reader {
    // from file
    // return openFile("hackerrank/picking_numbers/testcases/001.input")
    
    // from string
    // return strings.NewReader("6\n4 6 5 3 3 1")
    
    // from stdin (server config)
    return os.Stdin
}


// go specific
func getWriter() (*os.File, *bufio.Writer) {
    outputPath := os.Getenv("OUTPUT_PATH")
    if len(outputPath) > 0 {
        file, err := os.Create(outputPath)
        checkError(err)
        return file, bufio.NewWriterSize(file, 16*1024*1024)
    }
    return os.Stdout, bufio.NewWriterSize(os.Stdout, 16*1024*1024)
}

func readFile(filePath string) string {
    data, err := os.ReadFile(filePath)
    checkError(err)
    return string(data)
}

func openFile(filePath string) *os.File {
    file, err := os.Open(filePath)
    checkError(err)
    return file
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}

func flushWriter(writer *bufio.Writer) {
    err := writer.Flush()
    checkError(err)
}

func closeFile(file *os.File) {
    err := file.Close()
    checkError(err)
}

// print in output file
func out(writer *bufio.Writer, data string) {
    trimRight := strings.TrimRight(data, " ")
    _, err := fmt.Fprintf(writer, trimRight)
    checkError(err)
}

// get values from a single line
func split(line string, separator string) []string {
    return strings.Split(line, separator)
}

func splitInts(line string, separator string) []int {
    items := split(line, separator)
    list := make([]int, 0, len(items))
    for _, item := range items {
        list = append(list, parseInt(item))
    }
    return list
}

func splitLongs(line string, separator string) []int64 {
    items := split(line, separator)
    list := make([]int64, 0, len(items))
    for _, item := range items {
        list = append(list, parseLong(item))
    }
    return list
}

func splitDoubles(line string, separator string) []float64 {
    items := split(line, separator)
    list := make([]float64, 0, len(items))
    for _, item := range items {
        list = append(list, parseDouble(item))
    }
    return list
}

func getInt(line string, part int) int {
    split := split(line, " ")
    return parseInt(split[part])
}

func getLong(line string, part int) int64 {
    split := split(line, " ")
    return parseLong(split[part])
}

func getDouble(line string, part int) float64 {
    split := split(line, " ")
    return parseDouble(split[part])
}

// parsing strings
func parseInt(s string) int {
    value, err := strconv.ParseInt(s, 10, 32)
    checkError(err)
    return int(value)
}

func parseLong(s string) int64 {
    value, err := strconv.ParseInt(s, 10, 64)
    checkError(err)
    return value
}

func parseDouble(s string) float64 {
    value, err := strconv.ParseFloat(s, 64)
    checkError(err)
    return value
}

// read input file as list of strings
func getLines() []string {
    reader := bufio.NewReaderSize(getInput(), 16*1024*1024)
    var lines = make([]string, 0, 128)
    scanner := bufio.NewScanner(reader)
    scanner.Split(bufio.ScanLines)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines
}

// algorithm specific
func frequences(list []int) ([]int, map[int]int) {
    // creating map of frequencies
    freqMap := make(map[int]int, len(list))
    for _, item := range list {
        freqMap[item] += 1
    }
    // sorted list of items without repetition
    sortedItems := make([]int, 0, len(freqMap))
    for key, _ := range freqMap {
        sortedItems = append(sortedItems, key)
    }
    sort.Ints(sortedItems)
    return sortedItems, freqMap
}

func maxInt(a int, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxIntList(list []int) int {
    max := list[0]
    for _, item := range list {
        if item > max {
            max = item
        }
    }
    return max
}

func minInt(a int, b int) int {
    if a < b {
        return a
    }
    return b
}

func minIntList(list []int) int {
    min := list[0]
    for _, item := range list {
        if item < min {
            min = item
        }
    }
    return min
}