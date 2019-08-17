package brain

import (
	"reflect"
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

		// check again for 5
		subSeq5 := sequence[i : i+5]
		if reflect.DeepEqual(subSeq5, []int{0, playerId, playerId, playerId, 0}) {
			f3StartPoss = append(f3StartPoss, i)
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
				break // found enemy piece
			case val == playerId:
				totalMine++
			case (k == 0 || k == 3) && val == 0:
				totalMine = -10
				break
			}
		}

		if totalMine == 3 {
			f3StartPoss = append(f3StartPoss, i)
		}
	}
	return f3StartPoss
}
