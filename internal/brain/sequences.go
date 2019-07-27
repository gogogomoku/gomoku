package brain

import (
	"github.com/gogogomoku/gomoku/internal/board"
)

// For each direction in usual order, return how many contiguous pieces for playerId are present
func checkSequence(position int, playerId int) []int {
	sequenceLengths := []int{}
	for direction := 0; direction < 8; direction++ {
		counter := 0
		tmpPosition := position
		for {
			nextIndex, edge := getNextIndexForDirection(tmpPosition, direction)
			nextIndexValue, edge := ReturnNextPiece(tmpPosition, direction)
			if edge || nextIndexValue != playerId || nextIndexValue == 0 {
				break
			}
			counter++
			tmpPosition = nextIndex
		}
		sequenceLengths = append(sequenceLengths, counter)
	}
	return sequenceLengths
}

func sequenceOpposingDirections(position int, playerId int, dir []int, increase int) []int {
	partialSequences := checkSequence(position, playerId)
	if partialSequences[dir[0]] != 0 || partialSequences[dir[1]] != 0 {
		p := []int{position}
		// Add elements in dir[0]
		current := position
		for i := 0; i < partialSequences[dir[0]]; i++ {
			current -= increase
			p = append(p, current)
		}
		// Add elements in dir[1]
		current = position
		for i := 0; i < partialSequences[dir[1]]; i++ {
			current += increase
			p = append(p, current)
		}
		return p
	}
	return nil
}

func completeSequenceForPosition(position int, playerId int) [][]int {
	sequences := [][]int{}
	sequenceDirections := []struct {
		OpposingDirections []int
		IncreaseValue      int
	}{
		{
			OpposingDirections: []int{N, S},
			IncreaseValue:      board.SIZE,
		}, {
			OpposingDirections: []int{W, E},
			IncreaseValue:      1,
		}, {
			OpposingDirections: []int{NW, SE},
			IncreaseValue:      board.SIZE + 1,
		}, {
			OpposingDirections: []int{NE, SW},
			IncreaseValue:      board.SIZE - 1,
		},
	}
	for _, d := range sequenceDirections {
		seq := sequenceOpposingDirections(position, playerId, d.OpposingDirections, d.IncreaseValue)
		if seq != nil {
			sequences = append(sequences, seq)
		}
	}
	return sequences
}
