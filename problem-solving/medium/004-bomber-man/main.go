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
 * Complete the 'bomberMan' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING_ARRAY grid
 */

func bomberMan(n int32, grid []string) []string {
	// Write your code here
	var (
		g       = make([][]rune, len(grid))
		explode bool
		i       int32
		rows    = len(grid)
		cols    = len(grid[0])

		// show = func(time int32, action string) {
		// 	fmt.Printf("%4d:%s %t ----------\n", time, action, explode)
		// 	for _, line := range g {
		// 		fmt.Printf("%s\n", string(line))
		// 	}
		// }

		blow = func(r, c int) {
			var (
				up    = r - 1
				down  = r + 1
				left  = c - 1
				right = c + 1
			)

			// fmt.Printf("b(%d,%d) up:%d down:%d left:%d right:%d", r, c, up, down, left, right)
			g[r][c] = '.'
			if up >= 0 {
				if g[up][c] == 'O' {
					g[up][c] = '.'
				}
			}
			if down < rows {
				if g[down][c] == 'O' {
					g[down][c] = '.'
				}
			}
			if left >= 0 {
				if g[r][left] == 'O' {
					g[r][left] = '.'
				}
			}
			if right < cols {
				if g[r][right] == 'O' {
					g[r][right] = '.'
				}
			}
		}

		tick = func() {
			// fmt.Printf("> tick\n")
			// let's look for bombs flagged with * ready to explode
			for r := 0; r < rows; r++ {
				for c := 0; c < cols; c++ {
					if g[r][c] == '*' {
						blow(r, c)
					}
				}

			}
			explode = false
			// show(i, "t")
		}

		fill = func() {
			// fmt.Printf("> fill\n")
			for r := 0; r < rows; r++ {
				// existing bombs will explode in next iteration, so we flag them with *
				g[r] = []rune(strings.ReplaceAll(string(g[r]), "O", "*"))
				// empty spaces are fill with bombs
				g[r] = []rune(strings.ReplaceAll(string(g[r]), ".", "O"))

			}
			explode = true
			// show(i, "f")
		}
	)

	// edge case, no iteraction from bomberman
	if n == 1 {
		return grid
	}

	// edge case, bomberman fllls the missing slots with bombs
	if (n % 2) == 0 {
		full := strings.Repeat("O", len(grid[0]))
		for i := 0; i < len(grid); i++ {
			grid[i] = full
		}
		return grid
	}

	// For odd n > 3, reduce to either 3 or 5 steps
	if n > 5 {
		if (n-3)%4 == 0 {
			n = 3
		} else {
			n = 5
		}
	}

	// convert the string array to an array of rune arrays, easier to travel
	for r := 0; r < rows; r++ {
		g[r] = []rune(grid[r])
	}

	// show(0, "i")
	explode = false
	for i = int32(1); i < n; i++ {
		if explode {
			tick()
			continue
		}
		fill()
	}

	for r := 0; r < rows; r++ {
		grid[r] = strings.ReplaceAll(string(g[r]), "*", "O")
	}

	return grid
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	rTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	r := int32(rTemp)

	// cTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	// checkError(err)
	// c := int32(cTemp)

	nTemp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	n := int32(nTemp)

	var grid []string

	for i := 0; i < int(r); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := bomberMan(n, grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

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
