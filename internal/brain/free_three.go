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
	// axes := [4]int{}
	// sequences := make([][]int, 4)

	// col := board.GetColumnForPosition(position)
	// row := board.GetRowForPosition(position)
	// nesw := board.GetIndexNESWForPosition(position)
	// nwse := board.GetIndexNWSEForPosition(position)

	tab[position] = playerId

	axes := [4][2]int{}

	axes[NS][0], axes[NS][1] = board.GetIdxAndOffset(position, board.GetOffsetAndColumnForPosition)
	axes[EW][0], axes[EW][1] = board.GetIdxAndOffset(position, board.GetOffsetAndRowForPosition)
	axes[NWSE][0], axes[NWSE][1] = board.GetIdxAndOffset(position, board.GetOffsetAndIndexNWSEForPosition)
	axes[NESW][0], axes[NESW][1] = board.GetIdxAndOffset(position, board.GetOffsetAndIndexNESWForPosition)

	for i, axis := range axes {
		start := maximum(0, axis[0]-5)
		end := axis[0]
	}

	// todo: Actually need indices, not values :P
	// sequences[NS] = *(board.GetSequence(position, &tab, board.GetColumnForPosition, board.GetColSeqForCol))
	// sequences[EW] = *(board.GetSequence(position, &tab, board.GetRowForPosition, board.GetRowSeqForRow))
	// sequences[NESW] = *(board.GetSequence(position, &tab, board.GetIndexNESWForPosition, board.GetDiagonalNESWSequence))
	// sequences[NESW] = *(board.GetSequence(position, &tab, board.GetIndexNWSEForPosition, board.GetDiagonalNWSESequence))

	// todo: Actually need indices, not values :P
	// sequences[NS] = *(board.GetSequence(position, &tab, board.GetColumnForPosition, board.GetColSeqForCol))
	// sequences[EW] = *(board.GetSequence(position, &tab, board.GetRowForPosition, board.GetRowSeqForRow))
	// sequences[NESW] = *(board.GetSequence(position, &tab, board.GetIndexNESWForPosition, board.GetDiagonalNESWSequence))
	// sequences[NESW] = *(board.GetSequence(position, &tab, board.GetIndexNWSEForPosition, board.GetDiagonalNWSESequence))

	// sequences[NS] = *(board.GetIdxSequence(position, &tab, board.GetColumnForPosition, board.GetColSeqForColIdx))
	// sequences[EW] = *(board.GetIdxSequence(position, &tab, board.GetRowForPosition, board.GetRowSeqForRowIdx))
	// sequences[NESW] = *(board.GetIdxSequence(position, &tab, board.GetIndexNESWForPosition, board.GetDiagonalNESWSequenceIdx))
	// sequences[NWSE] = *(board.GetIdxSequence(position, &tab, board.GetIndexNWSEForPosition, board.GetDiagonalNWSESequenceIdx))

	// for axis, seq := range sequences {
	// 	switch {
	// 	case seq < 5:
	// 		break
	// 	}
	// 	for i, idx := range seq {

	// 	}
	// }

	_ = axes
	_ = sequences

	return axes
}

func GetTooManyF3Created(position int, tab [board.TOT_SIZE]int, playerId int) bool {
	return false
}
