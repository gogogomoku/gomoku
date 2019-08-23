package brain

import (
	"github.com/gogogomoku/gomoku/internal/board"
)

const (
	NS = iota
	EW
	NWSE
	NESW
)

/*
** Takes arbitrary slice of Goban sequence values
** Return starting index of each F3 found
 */

func CheckSequenceForF3(sequence []int, playerId int) []int {
	seqLen := len(sequence)
	if seqLen < 5 {
		return nil
	}

	nSubSeq := seqLen - 4 // how many sub sequences to evaluate
	f3StartPoss := []int{}

	for i := 0; i < nSubSeq; i++ {
		maxSubSeqLen := maximum(minimum(len(sequence[i:]), 5), minimum(len(sequence[i:]), 6))

		// length is < 5: no sequence
		// sequence must have no value at west extreme
		if maxSubSeqLen < 5 || sequence[i:][0] != 0 {
			continue
		}

		// check for only f3 5-len sequence
		subSeq5 := sequence[i : i+5]
		if subSeq5[0] == 0 && subSeq5[4] == 0 {
			if subSeq5[1] == playerId && subSeq5[1] == subSeq5[2] && subSeq5[2] == subSeq5[3] {
				f3StartPoss = append(f3StartPoss, i)
			}
		}

		// if only 5 in sequence and doesn't match above pattern, not an f3
		if maxSubSeqLen == 5 {
			continue
		}

		// now try for broken sequence of 3
		subSeq6 := sequence[i : i+6]

		if subSeq6[5] != 0 {
			continue // east extreme must be 0
		}

		// check middle of sequence
		totalMine := 0
		for k, val := range subSeq6[1:5] {
			switch {
			case val > 0 && val != playerId:
				totalMine = -10
			case val == playerId:
				totalMine++
			case (k == 0 || k == 3) && val == 0:
				totalMine = -10
			}
			if totalMine < 0 {
				break
			}
		}

		if totalMine == 3 {
			f3StartPoss = append(f3StartPoss, i)
		}
	}
	return f3StartPoss
}

/*
** For given board position,
** checks for creation of 2 f3s
 */

 func GetDoubleF3StartPos(position int, tab [board.TOT_SIZE]int, playerId int) [4]int {
	axes := [4]int{}
	sequences := make([][]int, 4)

	// axes[NS] = board.GetColumnForPosition(position, &tab)
	// axes[EW] = board.GetRowForPosition(position, &tab)
	// axes[NWSE] = board.GetIndexNWSEForPosition(position, &tab)
	// axes[NESW] = board.GetIndexNESWForPosition(position, &tab)

	// sequences[NS] = board.GetNSSequenceForColumn(axes[NS], &tab)
	// sequences[EW] = board.GetEWSequenceForRow(axes[EW], &tab)
	// sequences[NESW] = *(board.GetDiagonalNESWSequence(axes[NESW], &tab))
	// sequences[NWSE] = *(board.GetDiagonalNWSESequence(axes[NWSE], &tab))

	_ = axes
	_ = sequences

	return axes
}

func GetTooManyF3Created(position int, tab [board.TOT_SIZE]int, playerId int) bool {
	return false
}