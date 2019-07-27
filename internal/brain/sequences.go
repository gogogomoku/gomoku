package brain

import (
	"github.com/gogogomoku/gomoku/internal/board"
)

// For each direction in usual order, return how many contiguous pieces for playerId are present
func CheckSequence(position int, playerId int) []int {
	// fmt.Println("===================== Checking sequence for ", position)
	sequenceLengths := []int{0, 0, 0, 0, 0, 0, 0, 0}
	for direction := 0; direction < 8; direction++ {
		// fmt.Println("Direction: ", direction)
		counter := 0
		tmpPosition := position
		for i := 0; i < 8; i++ {
			nextIndex, edge := getNextIndexForDirection(tmpPosition, direction)
			nextIndexValue, edge := ReturnNextPiece(tmpPosition, direction)
			if edge || nextIndexValue != playerId || nextIndexValue == 0 {
				break
			}
			counter++
			tmpPosition = nextIndex
		}
		sequenceLengths[direction] = counter
		// fmt.Println("Counter: ", counter)
	}
	// fmt.Println("sequence: ", sequenceLengths)
	return sequenceLengths
}

func sequenceOpposingDirections(position int, playerId int, dir []int, increase int) []int {
	// fmt.Println("SEQ OPPOSING Start...")
	partialSequences := CheckSequence(position, playerId)
	// fmt.Println(partialSequences)
	// fmt.Println("SEQ OPPOSING call to CheckSequence() ended...")

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
	return []int{}
}

func CompleteSequenceForPosition(position int, playerId int) [][]int {
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
	// fmt.Println("COMPLETE SEQ Start...")
	for _, d := range sequenceDirections {
		seq := sequenceOpposingDirections(position, playerId, d.OpposingDirections, d.IncreaseValue)
		if len(seq) > 0 {
			sequences = append(sequences, seq)
		}
	}
	// fmt.Println("COMPLETE SEQ End...")

	return sequences
}
