package board

import (
	"fmt"
)

func PositionInSubsequence(subsequence *[11]int, position int) (bool, int) {
	for i, val := range *subsequence {
		switch {
		case val == -1:
			return false, -1
		case val == position:
			return true, i
		}
	}
	fmt.Println("Hi.")
	return false, -1
}

/*
** GET VALUE SEQUENCES
 */

func GetDiagonalNWSESequence(d int, tab *[TOT_SIZE]int) *[]int {
	line := [SIZE * 2]int{}
	// First element of last row to start diagonals
	x := SIZE * (SIZE - d)
	// When transpassing board limit, start new diagonal
	lineLen := 0
	for x < TOT_SIZE {
		if x >= 0 && x < (TOT_SIZE)-(SIZE*(d-SIZE)) {
			line[lineLen] = (*tab)[x]
			lineLen++
		}
		x += SIZE + 1
	}
	line[lineLen] = -1
	line2 := line[:lineLen]
	return &line2
}

func GetDiagonalNESWSequence(d int, tab *[TOT_SIZE]int) *[]int {
	line := [SIZE * 2]int{}
	// Last element of first row to start diagonals
	x := (d*SIZE - 1) - (SIZE-1)*SIZE
	// When transpassing board limit, start new diagonal
	lineLen := 0
	for x < TOT_SIZE {
		if x >= 0 && x <= (d-1)*SIZE {
			line[lineLen] = (*tab)[x]
			lineLen++
		}
		x += SIZE - 1
	}
	line[lineLen] = -1
	line2 := line[:lineLen]
	return &line2
}

func GetColSeqForCol(column int, tab *[TOT_SIZE]int) *[]int {
	sequence := make([]int, SIZE)
	for i := 0; i < SIZE; i++ {
		sequence[i] = (*tab)[column+(SIZE*i)]
	}
	return &sequence
}

func GetRowSeqForRow(row int, tab *[TOT_SIZE]int) *[]int {
	sequence := make([]int, SIZE)
	for i := 0; i < SIZE; i++ {
		sequence[i] = (*tab)[(SIZE*row)+i]
	}
	return &sequence
}

/*
** GET START INDICES
 */

func GetColumnForPosition(position int) int {
	return position % SIZE
}

func GetRowForPosition(position int) int {
	return position / SIZE
}

func GetIndexNWSEForPosition(position int) int {
	column := position % SIZE
	row := position / SIZE // TODO: Switch direction of row and starts with 1

	d := (SIZE - row) + column
	return d
}

func GetIndexNESWForPosition(position int) int {
	column := position % SIZE
	row := position / SIZE

	d := column + row + 1

	return d
}

/*
** GET INDEX SEQUENCES
 */

func GetDiagonalNWSESequenceIdx(d int, tab *[TOT_SIZE]int) *[]int {
	line := [SIZE * 2]int{}
	// First element of last row to start diagonals
	x := SIZE * (SIZE - d)
	// When transpassing board limit, start new diagonal
	lineLen := 0
	for x < TOT_SIZE {
		if x >= 0 && x < (TOT_SIZE)-(SIZE*(d-SIZE)) {
			line[lineLen] = x
			lineLen++
		}
		x += SIZE + 1
	}
	line[lineLen] = -1
	line2 := line[:lineLen]
	return &line2
}

func GetDiagonalNESWSequenceIdx(d int, tab *[TOT_SIZE]int) *[]int {
	line := [SIZE * 2]int{}
	// Last element of first row to start diagonals
	x := (d*SIZE - 1) - (SIZE-1)*SIZE
	// When transpassing board limit, start new diagonal
	lineLen := 0
	for x < TOT_SIZE {
		if x >= 0 && x <= (d-1)*SIZE {
			line[lineLen] = x
			lineLen++
		}
		x += SIZE - 1
	}
	line[lineLen] = -1
	line2 := line[:lineLen]
	return &line2
}

func GetColSeqForColIdx(column int, tab *[TOT_SIZE]int) *[]int {
	sequence := make([]int, SIZE)
	for i := 0; i < SIZE; i++ {
		sequence[i] = column + (SIZE * i)
	}
	return &sequence
}

func GetRowSeqForRowIdx(row int, tab *[TOT_SIZE]int) *[]int {
	sequence := make([]int, SIZE)
	for i := 0; i < SIZE; i++ {
		sequence[i] = (SIZE * row) + i
	}
	return &sequence
}

/*
** POINTLESS WRAPPERS
 */

func GetSequence(position int, tab *[TOT_SIZE]int, getIndex func(int) int, getSequence func(int, *[TOT_SIZE]int) *[]int) *[]int {
	index := getIndex(position)
	sequence := getSequence(index, tab)
	return sequence
}

func GetIdxSequence(position int, tab *[TOT_SIZE]int, getIndex func(int) int, getIdxSequence func(int, *[TOT_SIZE]int) *[]int) *[]int {
	index := getIndex(position)
	idxSequence := getIdxSequence(index, tab)
	return idxSequence
}

func GetIdxAndOffset(position int, getOffsetAndIdx func(int) (int, int)) (offset int, idx int) {
	offset, index := getOffsetAndIdx(position)
	return offset, index
}

/*
** GET INDEX SEQUENCES WITH PLAYER OFFSET
 */

func GetDiagonalNWSESequenceIdxPos(d int, tab *[TOT_SIZE]int, position int) (*[]int, int) {
	line := [SIZE * 2]int{}
	positionOffset := -1
	// First element of last row to start diagonals
	x := SIZE * (SIZE - d)
	// When transpassing board limit, start new diagonal
	lineLen := 0
	for x < TOT_SIZE {
		if x >= 0 && x < (TOT_SIZE)-(SIZE*(d-SIZE)) {
			line[lineLen] = x
			if position == x {
				positionOffset = lineLen
			}
			lineLen++
		}
		x += SIZE + 1
	}
	line[lineLen] = -1
	line2 := line[:lineLen]
	return (&line2), positionOffset
}

func GetDiagonalNESWSequenceIdxPos(d int, tab *[TOT_SIZE]int, position int) (*[]int, int) {
	line := [SIZE * 2]int{}
	positionOffset := -1
	// Last element of first row to start diagonals
	x := (d*SIZE - 1) - (SIZE-1)*SIZE
	// When transpassing board limit, start new diagonal
	lineLen := 0
	for x < TOT_SIZE {
		if x >= 0 && x <= (d-1)*SIZE {
			line[lineLen] = x
			if position == x {
				positionOffset = lineLen
			}
			lineLen++
		}
		x += SIZE - 1
	}
	line[lineLen] = -1
	line2 := line[:lineLen]
	return &line2, positionOffset
}

func GetColSeqForColIdxPos(column int, tab *[TOT_SIZE]int, position int) (*[]int, int) {
	sequence := make([]int, SIZE)
	offset := 0
	for i := 0; i < SIZE; i++ {
		sequence[i] = column + (SIZE * i)
		if i == position {
			offset = i
		}
	}
	return &sequence, offset
}

func GetRowSeqForRowIdxPos(row int, tab *[TOT_SIZE]int, position int) (*[]int, int) {
	sequence := make([]int, SIZE)
	offset := 0
	for i := 0; i < SIZE; i++ {
		sequence[i] = (SIZE * row) + i
		if i == position {
			offset = i
		}
	}
	return &sequence, offset
}

/*
** INDEX + OFFSET
 */

func GetOffsetAndColumnForPosition(position int) (int, int) {
	col := position % SIZE
	offset := position / SIZE
	return offset, col
}

func GetOffsetAndRowForPosition(position int) (int, int) {
	row := position / SIZE
	offset := position % SIZE
	return offset, row
}

func GetOffsetAndIndexNWSEForPosition(position int) (int, int) {
	column := position % SIZE
	row := position / SIZE // TODO: Switch direction of row and starts with 1

	d := (SIZE - row) + column

	x := SIZE * (SIZE - d)
	offset := 0
	y := 0
	for x < TOT_SIZE {
		if x >= 0 && x < (TOT_SIZE)-(SIZE*(d-SIZE)) {
			if position == x {
				offset = y
				break
			}
			y++
		}
		x += SIZE + 1
	}
	return offset, d
}

func GetOffsetAndIndexNESWForPosition(position int) (int, int) {
	column := position % SIZE
	row := position / SIZE

	d := column + row + 1

	x := (d*SIZE - 1) - (SIZE-1)*SIZE
	offset := 0
	y := 0

	for x < TOT_SIZE {
		if x >= 0 && x < (TOT_SIZE)-(SIZE*(d-SIZE)) {
			if x == position {
				offset = y
				break
			}
			y++
		}
		x += SIZE - 1
	}

	return offset, d
}
