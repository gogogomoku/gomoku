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

func getDiagonalNWSESequence(tab *[TOT_SIZE]int, d int) *[SIZE * 2]int {
	line := [SIZE * 2]int{}
	x := SIZE * (SIZE - d)
	lineLen := 0
	for x < TOT_SIZE {
		if x >= 0 && x < (TOT_SIZE)-(SIZE*(d-SIZE)) {
			line[lineLen] = (*tab)[x]
			lineLen++
		}
		x += SIZE + 1
	}
	line[lineLen] = -1
	// line2 := line[:lineLen]
	return &line

}

func GetDiagonalNWSESequences(playerId int, tab *[TOT_SIZE]int) [(SIZE + SIZE/2) - 1][SIZE * 2]int {
	nwseSequences := [(SIZE + SIZE/2) - 1][SIZE * 2]int{}
	for d := 1; d < SIZE+SIZE/2; d++ {
		nwseSequences[d - 1] = *getDiagonalNWSESequence(tab, d)
	}
	return nwseSequences
}

func GetDiagonalNESWSequences(playerId int, tab *[TOT_SIZE]int) [][]int {
	return nil
}