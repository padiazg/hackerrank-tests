package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the catAndMouse function below.
func catAndMouse(x int32, y int32, z int32) string {
	// x => Cat A
	// y => Cat B
	// z => Mouse
	var (
		catAtoMouse = int32(0)
		catBtoMouse = int32(0)
	)

	// Cat A relative position
	if z > x {
		catAtoMouse = z - x
	} else {
		catAtoMouse = x - z
	}

	// Cat B relative position
	if z > y {
		catBtoMouse = z - y
	} else {
		catBtoMouse = y - z
	}

	// fmt.Printf("Cat A=%d Cat B=%d Mouse=%d catAtoMouse=%d catBtoMouse=%d\n", x, y, z, catAtoMouse, catBtoMouse)
	switch {
	case catAtoMouse == catBtoMouse:
		return "Mouse C"
	case catAtoMouse < catBtoMouse:
		return "Cat A"
	default:
		return "Cat B"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		xyz := strings.Split(readLine(reader), " ")

		xTemp, err := strconv.ParseInt(xyz[0], 10, 64)
		checkError(err)
		x := int32(xTemp)

		yTemp, err := strconv.ParseInt(xyz[1], 10, 64)
		checkError(err)
		y := int32(yTemp)

		zTemp, err := strconv.ParseInt(xyz[2], 10, 64)
		checkError(err)
		z := int32(zTemp)

		result := catAndMouse(x, y, z)

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
