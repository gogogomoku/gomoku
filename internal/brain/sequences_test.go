package brain

import (
	"reflect"
	"testing"

	"github.com/gogogomoku/gomoku/internal/board"
)

func TestCheckSequence(t *testing.T) {
	// Initialize
	Game.Goban.Tab = [board.TOT_SIZE]int16{}
	Game.CurrentPlayer = Game.P1
	center := int16((board.SIZE * board.SIZE) / 2)
	if board.SIZE%2 == 0 {
		center += board.SIZE / 2
	}
	tables := []struct {
		position               int16
		opponentPositions      []int16
		currentPlayerPositions []int16
		expectedSequences      []int16
	}{
		{
			position:               center,
			opponentPositions:      []int16{},
			currentPlayerPositions: []int16{},
			expectedSequences:      []int16{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			position:               center,
			opponentPositions:      []int16{center + 1, center + 2},
			currentPlayerPositions: []int16{center, center + 3},
			expectedSequences:      []int16{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			position:               center,
			opponentPositions:      []int16{center - 1, center - 2},
			currentPlayerPositions: []int16{center, center + 1, center - 3},
			expectedSequences:      []int16{0, 0, 1, 0, 0, 0, 0, 0},
		},
		{
			position:               center,
			opponentPositions:      []int16{center - board.SIZE, center - (2 * board.SIZE)},
			currentPlayerPositions: []int16{center, center + board.SIZE, center - (board.SIZE), center + 1, center - 1, center - 2, center - 3},
			expectedSequences:      []int16{1, 1, 1, 3, 0, 0, 0, 0},
		},
		{
			position:               center,
			opponentPositions:      []int16{},
			currentPlayerPositions: []int16{center, center - (board.SIZE) + 1, center - (board.SIZE) - 1, center + (board.SIZE) + 1, center + (board.SIZE) - 1},
			expectedSequences:      []int16{0, 0, 0, 0, 1, 1, 1, 1},
		},
	}
	for _, table := range tables {
		for _, v := range table.opponentPositions {
			Game.Goban.Tab[v] = 2
		}
		for _, v := range table.currentPlayerPositions {
			Game.Goban.Tab[v] = 1
		}
		sequenceLengths := CheckSequence(table.position, Game.CurrentPlayer.Id, &Game.Goban.Tab)

		if !reflect.DeepEqual(table.expectedSequences, sequenceLengths) {

			t.Errorf("Wrong sequenceLengths for %d, expected %v, got %v", table.position, table.expectedSequences, sequenceLengths)

		}
		// Game.Gban.Tab = make([]int16, board.SIZE*board.SIZE)
		Game.Goban.Tab = [board.TOT_SIZE]int16{}
	}
}

func BenchmarkCheckSequence(b *testing.B) {
	Game.Goban.Tab = [board.TOT_SIZE]int16{}
	Game.CurrentPlayer = Game.P1
	for i := 0; i < b.N; i++ {
		CheckSequence(int16(i)%(board.SIZE*board.SIZE), 1, &Game.Goban.Tab)
	}
}

func TestCompleteSequenceForPosition(t *testing.T) {
	// Initialize
	Game.Goban.Tab = [board.TOT_SIZE]int16{}
	Game.CurrentPlayer = Game.P1
	center := int16((board.SIZE * board.SIZE) / 2)
	if board.SIZE%2 == 0 {
		center += board.SIZE / 2
	}
	tables := []struct {
		position               int16
		currentPlayerPositions []int16
		expectedSequences      [][]int16
	}{
		{
			position:               center,
			currentPlayerPositions: []int16{center, center + 1, center - 1},
			expectedSequences:      [][]int16{[]int16{center, center - 1, center + 1}},
		},
		{
			position:               center,
			currentPlayerPositions: []int16{center, center + board.SIZE, center - board.SIZE},
			expectedSequences:      [][]int16{[]int16{center, center - board.SIZE, center + board.SIZE}},
		},
		{
			position:               center,
			currentPlayerPositions: []int16{center, center + (board.SIZE + 1), center - (board.SIZE + 1)},
			expectedSequences:      [][]int16{[]int16{center, center - (board.SIZE + 1), center + (board.SIZE + 1)}},
		},
		{
			position:               center,
			currentPlayerPositions: []int16{center, center + (board.SIZE - 1), center - (board.SIZE - 1)},
			expectedSequences:      [][]int16{[]int16{center, center - (board.SIZE - 1), center + (board.SIZE - 1)}},
		},
	}
	for _, table := range tables {
		for _, v := range table.currentPlayerPositions {
			Game.Goban.Tab[v] = 1
		}
		completeSequences := CompleteSequenceForPosition(table.position, Game.CurrentPlayer.Id, &Game.Goban.Tab)
		if !reflect.DeepEqual(table.expectedSequences, completeSequences) {
			t.Errorf("Wrong completeSequences for %d, expected %v, got %v", table.position, table.expectedSequences, completeSequences)
		}
		Game.Goban.Tab = [board.TOT_SIZE]int16{}
	}
}
