package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'matrixRotation' function below.
 *
 * The function accepts following parameters:
 *  1. 2D_INTEGER_ARRAY matrix
 *  2. INTEGER r
 */

func matrixRotation(matrix [][]int32, r int32) {
	// Write your code here

	fmt.Printf("")
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	mTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	m := int32(mTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	n := int32(nTemp)

	rTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	r := int32(rTemp)

	var matrix [][]int32
	for i := 0; i < int(m); i++ {
		matrixRowTemp := strings.Split(strings.TrimRight(readLine(reader), " \t\r\n"), " ")

		var matrixRow []int32
		for _, matrixRowItem := range matrixRowTemp {
			matrixItemTemp, err := strconv.ParseInt(matrixRowItem, 10, 64)
			checkError(err)
			matrixItem := int32(matrixItemTemp)
			matrixRow = append(matrixRow, matrixItem)
		}

		if len(matrixRow) != int(n) {
			panic("Bad input")
		}

		matrix = append(matrix, matrixRow)
	}

	matrixRotation(matrix, r)
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
