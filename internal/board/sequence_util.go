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

func GetColumnForPosition(position int, tab *[TOT_SIZE]int) int {
	return position % SIZE
}

func GetRowForPosition(position int, tab *[TOT_SIZE]int) int {
	return position / SIZE
}

func GetIndexNWSEForPosition(position int, tab *[TOT_SIZE]int) int {
	column := position % SIZE
	row := position / SIZE // TODO: Switch direction of row and starts with 1

	d := (SIZE - row) + column
	return d
}

func GetIndexNESWForPosition(position int, tab *[TOT_SIZE]int) int {
	column := position % SIZE
	row := position / SIZE

	d := column + row + 1

	return d
}

func GetSequence(position int, tab *[TOT_SIZE]int, getIndex func(int, *[TOT_SIZE]int) int, getSequence func(int, *[TOT_SIZE]int) *[]int) *[]int {
	index := getIndex(position, tab)
	sequence := getSequence(index, tab)
	return sequence
}
