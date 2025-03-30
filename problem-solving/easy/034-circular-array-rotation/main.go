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
 * Complete the 'circularArrayRotation' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER k
 *  3. INTEGER_ARRAY queries
 */

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	var (
		rotate = func(cicles int32) []int32 {
			b := make([]int32, len(a))
			b = append([]int32(nil), a...)

			for _ = range cicles {
				tmp := []int32{b[len(b)-1]}
				tmp = append(tmp, b[:len(b)-1]...)
				b = tmp
			}

			return b
		}

		ret = make([]int32, len(queries))
	)

	a = rotate(k)
	for i, q := range queries {
		ret[i] = a[q]
	}

	return ret
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	qTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	q := int32(qTemp)

	aTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	var queries []int32

	for i := 0; i < int(q); i++ {
		queriesItemTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := circularArrayRotation(a, k, queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
