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

var AXES = [4]int16{NS, EW, NWSE, NESW}

const N_DIAG_SEQ_DIV_2 = 18
const NWSE_FIRST_SEQ_GOBAN_POS = 342

/*
** Takes arbitrary slice of Goban sequence values
** Return starting index of each F3 found
 */

func CheckSequenceForF3(sequence []int16, playerId int16) []int16 {
	seqLen := len(sequence)
	if seqLen < 5 {
		return nil
	}

	nSubSeq := int16(seqLen - 4) // how many sub sequences to evaluate
	f3StartPoss := []int16{}

	var i int16
	for i = 0; i < nSubSeq; i++ {
		maxSubSeqLen := (maximum(int16(minimum(int16(len(sequence[i:])), int16(5))), int16(minimum(int16(len(sequence[i:])), int16(6)))))

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

func abs(n int16) int16 {
	return n * -1
}

func seqToGoban(start int16, step int16, i int16) int16 {
	return start + (step * i)
}

func gobanToSeq(start int16, step int16, i int16) int16 {
	return (i - start) / step
}

func checkAxisForF3(
	axis int16,
	len int16,
	playerId int16,
	seqOffset int16,
	startGobanI int16,
	step int16,
	tab *[board.TOT_SIZE]int16,
	whichSeq int16,
) bool {
	startSearchSeqOffset := maximum(0, seqOffset-4)
	lastStartSearchOffset := minimum(len-5, seqOffset-1)
	// no possible seq
	if lastStartSearchOffset < 0 {
		return false
	}
	// no empty start
	if seqOffset == 0 {
		return false
	}

	// go through all possible sequences until find an F3 in this axis
	for currSearch := startSearchSeqOffset; currSearch <= lastStartSearchOffset; currSearch++ {

		gobanIdx := seqToGoban(startGobanI, step, currSearch)
		// sequence not starting with 0
		if (*tab)[gobanIdx] != 0 {
			continue
		}
		// see what kind of f3 we can look for
		search5 := true
		search6 := true
		if seqToGoban(startGobanI, step, currSearch+5) >= board.TOT_SIZE {
			search6 = false
		} else if (*tab)[seqToGoban(startGobanI, step, currSearch+5)] != 0 {
			// doesn't end in 0
			search6 = false
		}
		if (seqToGoban(startGobanI, step, currSearch+4)) > board.TOT_SIZE {
			search5 = false
		} else if (*tab)[seqToGoban(startGobanI, step, currSearch+4)] != 0 {
			search5 = false
		}
		if !search6 && !search5 {
			// no closing 0 at end of either size sequence
			continue
		}
		if search5 {
			if (*tab)[seqToGoban(startGobanI, step, currSearch+1)] == playerId &&
				(*tab)[seqToGoban(startGobanI, step, currSearch+2)] == playerId &&
				(*tab)[seqToGoban(startGobanI, step, currSearch+3)] == playerId &&
				(*tab)[seqToGoban(startGobanI, step, currSearch+4)] == 0 {
				return true
			}
		}
		if search6 {
			if (*tab)[seqToGoban(startGobanI, step, currSearch+1)] == playerId &&
				(((*tab)[seqToGoban(startGobanI, step, currSearch+2)] == playerId && (*tab)[seqToGoban(startGobanI, step, currSearch+3)] == 0) || ((*tab)[seqToGoban(startGobanI, step, currSearch+2)] == 0 && (*tab)[seqToGoban(startGobanI, step, currSearch+3)] == playerId)) &&
				(*tab)[seqToGoban(startGobanI, step, currSearch+4)] == playerId {
				return true
			}
		}
	}
	return false
}

/*
** Checks if player's desired position will create 2 F3s.
 */

func Check2F3s(
	playerId int16,
	position int16,
	tab *[board.TOT_SIZE]int16,
) bool {
	nF3 := 0
	(*tab)[position] = playerId
	for i, axis := range AXES {
		// No need to check last axis if haven't found f3 yet
		if nF3 < 1 && i == 3 {
			(*tab)[position] = 0
			return false
		}
		var whichSeq int16
		var len int16
		var startGobanI int16
		var step int16
		var seqOffset int16

		switch axis {
		case NS:
			whichSeq = int16(position % board.SIZE)
			len = int16(board.SIZE)
			startGobanI = whichSeq
			step = int16(board.SIZE)
			seqOffset = int16((position - startGobanI) / step)
		case EW: // rows
			whichSeq = int16(position / board.SIZE)
			len = int16(board.SIZE)
			startGobanI = len * whichSeq
			step = int16(1)
			seqOffset = position % board.SIZE
		case NWSE: // diagonal
			row := int16(position / board.SIZE)
			col := int16(position % board.SIZE)
			whichSeq = board.SIZE - row + col - 1
			len = whichSeq + 1
			if whichSeq > board.SIZE-1 {
				len = (board.SIZE - 1) - whichSeq
			}
			if len < 5 {
				break
			}
			step = board.SIZE + 1
			startGobanI = NWSE_FIRST_SEQ_GOBAN_POS - (board.SIZE * whichSeq)
			if whichSeq >= N_DIAG_SEQ_DIV_2 {
				startGobanI = whichSeq - N_DIAG_SEQ_DIV_2
			}
			seqOffset = gobanToSeq(startGobanI, step, position)
		case NESW: // diagonal top-left to bottom-right
			row := int16(position / board.SIZE)
			col := int16(position % board.SIZE)
			whichSeq = (board.SIZE - 1 - row) + (board.SIZE - 1 - col)
			len = whichSeq + 1
			if whichSeq > board.SIZE-1 {
				tmp := whichSeq % board.SIZE
				len = board.SIZE - tmp - 1
			}
			if len < 5 {
				break
			}
			step = board.SIZE - 1
			startGobanI = 360 - (board.SIZE * whichSeq)
			if whichSeq >= N_DIAG_SEQ_DIV_2 {
				startGobanI = len - 1
			}
			seqOffset = gobanToSeq(startGobanI, step, position)
		}
		if checkAxisForF3(
			axis,
			len,
			playerId,
			seqOffset,
			startGobanI,
			step,
			tab,
			whichSeq,
		) {
			nF3++
		}
		if nF3 > 1 {
			(*tab)[position] = 0
			return true
		}
	}
	(*tab)[position] = 0
	return nF3 > 1
}
