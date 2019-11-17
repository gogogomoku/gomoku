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

const (
	N_DIAG_SEQ_DIV_2         = 18
	NWSE_FIRST_SEQ_GOBAN_POS = 342
	MAX_ROW                  = 18
	MAX_COL                  = 18
)

var AXES = [4]int16{NS, EW, NWSE, NESW}

// For each direction in usual order, return how many contiguous pieces for playerId are present
func CheckSequence(position int16, playerId int16, tab *[board.TOT_SIZE]int16) []int16 {
	sequenceLengths := []int16{0, 0, 0, 0, 0, 0, 0, 0}
	for direction := int16(0); direction < 8; direction++ {
		counter := int16(0)
		tmpPosition := position
		for i := 0; i < 8; i++ {
			nextIndex, edge := getNextIndexForDirection(tmpPosition, direction)

			if edge {
				break
			}
			nextIndexValue := (*tab)[nextIndex]
			if nextIndexValue != playerId || nextIndexValue == 0 {
				break
			}
			counter++
			tmpPosition = nextIndex
		}
		sequenceLengths[direction] = counter
	}
	return sequenceLengths
}

func sequenceOpposingDirections(position int16, playerId int16, dir []int16, increase int16, tab *[board.TOT_SIZE]int16) []int16 {
	partialSequences := CheckSequence(position, playerId, tab)
	if partialSequences[dir[0]] != 0 || partialSequences[dir[1]] != 0 {
		p := []int16{position}
		// Add elements in dir[0]
		current := position
		for i := int16(0); i < partialSequences[dir[0]]; i++ {
			current -= increase
			p = append(p, current)
		}
		// Add elements in dir[1]
		current = position
		for i := int16(0); i < partialSequences[dir[1]]; i++ {
			current += increase
			p = append(p, current)
		}
		return p
	}
	return []int16{}
}

func CompleteSequenceForPosition(position int16, playerId int16, tab *[board.TOT_SIZE]int16) [][]int16 {
	sequences := [][]int16{}
	sequenceDirections := []struct {
		OpposingDirections []int16
		IncreaseValue      int16
	}{
		{
			OpposingDirections: []int16{N, S},
			IncreaseValue:      board.SIZE,
		}, {
			OpposingDirections: []int16{W, E},
			IncreaseValue:      1,
		}, {
			OpposingDirections: []int16{NW, SE},
			IncreaseValue:      board.SIZE + 1,
		}, {
			OpposingDirections: []int16{NE, SW},
			IncreaseValue:      board.SIZE - 1,
		},
	}
	for _, d := range sequenceDirections {
		seq := sequenceOpposingDirections(position, playerId, d.OpposingDirections, d.IncreaseValue, tab)
		if len(seq) > 0 {
			sequences = append(sequences, seq)
		}
	}
	return sequences
}

// Return N next pieces for every directions
func CheckNextN(position int16, tab [board.TOT_SIZE]int16, size int16) [][]int16 {
	lines := make([][]int16, 8)
	for direction := int16(0); direction < 8; direction++ {
		tmpPosition := position
		for i := int16(0); i < size; i++ {
			nextIndex, edge := getNextIndexForDirection(tmpPosition, direction)

			if edge {
				break
			}
			nextIndexValue := tab[nextIndex]
			lines[direction] = append(lines[direction], nextIndexValue)
			tmpPosition = nextIndex
		}
	}
	return lines
}
