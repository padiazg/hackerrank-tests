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
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	var (
		count int32

		blocksGenerator = func() func() (bool, bool) {
			var index int
			return func() (bool, bool) {
				if index >= len(s) {
					return false, true
				}

				var (
					length = int32(0)
					sum    = int32(0)
					match  bool
				)

				for i := index; i < len(s); i++ {
					length++
					sum += s[i]
					if length > m || sum > d {
						match = false
						break
					}

					if length == m && sum == d {
						match = true
						break
					}
				}

				index++
				return match, false
			}
		}
	)

	next := blocksGenerator()
	for {
		match, finished := next()
		if finished {
			break
		}
		if match {
			count++
		}
	}

	return count
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	sTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var s []int32

	for i := 0; i < int(n); i++ {
		sItemTemp, err := strconv.ParseInt(sTemp[i], 10, 64)
		checkError(err)
		sItem := int32(sItemTemp)
		s = append(s, sItem)
	}

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	dTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	d := int32(dTemp)

	mTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	m := int32(mTemp)

	result := birthday(s, d, m)

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
