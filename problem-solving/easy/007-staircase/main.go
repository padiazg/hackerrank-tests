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
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	// Write your code here
	for bricks := 1; bricks <= int(n); bricks++ {
		fmt.Printf("%s%s\n",
			strings.Repeat(" ", int(n)-bricks),
			strings.Repeat("#", bricks),
		)
	}
}

// func staircase(n int32) {
// 	// Write your code here
// 	var (
// 		bricks int = 1
// 		spaces int = int(n) - 1
// 	)

// 	for i := 0; i < int(n); i++ {
// 		fmt.Printf("%s%s\n",
// 			strings.Repeat(" ", spaces),
// 			strings.Repeat("#", bricks),
// 		)
// 		bricks++
// 		spaces--
// 	}
// }

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
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
