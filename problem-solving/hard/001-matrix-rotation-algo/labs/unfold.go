package main

import (
	"fmt"
	"strings"
)

func unfold(matrix [][]int32) {
	type Pos struct {
		left, right int32
	}

	type Level struct {
		elements []rune
		length   int32
		start    int32
		end      int32
		middle   int32
		top      Pos
		bottom   Pos
	}

	var (
		height   = int32(len(matrix))
		width    = int32(len(matrix[0]))
		minDim   = min(height, width)
		maxLevel = int32(minDim / 2)
		levels   = map[int32]Level{}
		start    = int32(0)
		middle   = width + (height - 1)
	)

	fmt.Printf("min=%d levels=%d\n", minDim, maxLevel)

	for l := int32(0); l < maxLevel; l++ {
		length := (height + width - 2) * 2
		levels[l] = Level{
			elements: []rune(strings.Repeat(".", int(length))),
			length:   length,
			start:    start,
			// end:      width,
			middle: middle,
			// top: Pos{
			// 	left:  start + length + 1,
			// 	right: start + width,
			// },
			// bottom: Pos{
			// 	left: middle - 1,
			// },
		}

		// decrease boudaries for next level
		height -= 2
		width -= 2
	}

	// show
	for k, v := range levels {
		fmt.Printf("%d length:%2d start:%2d middle:%2d %s\n", k, v.length, v.start, v.middle, string(v.elements))
	}

}
