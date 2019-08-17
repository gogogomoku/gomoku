package brain

import (
	"reflect"
	"testing"

	"github.com/gogogomoku/gomoku/internal/board"
)

func TestCheckSequence(t *testing.T) {
	// Initialize
	Game.Goban.Tab = [board.TOT_SIZE]int{}
	Game.CurrentPlayer = Game.P1
	center := (board.SIZE * board.SIZE) / 2
	if board.SIZE%2 == 0 {
		center += board.SIZE / 2
	}
	tables := []struct {
		position               int
		opponentPositions      []int
		currentPlayerPositions []int
		expectedSequences      []int
	}{
		{
			position:               center,
			opponentPositions:      []int{},
			currentPlayerPositions: []int{},
			expectedSequences:      []int{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			position:               center,
			opponentPositions:      []int{center + 1, center + 2},
			currentPlayerPositions: []int{center, center + 3},
			expectedSequences:      []int{0, 0, 0, 0, 0, 0, 0, 0},
		},
		{
			position:               center,
			opponentPositions:      []int{center - 1, center - 2},
			currentPlayerPositions: []int{center, center + 1, center - 3},
			expectedSequences:      []int{0, 0, 1, 0, 0, 0, 0, 0},
		},
		{
			position:               center,
			opponentPositions:      []int{center - board.SIZE, center - (2 * board.SIZE)},
			currentPlayerPositions: []int{center, center + board.SIZE, center - (board.SIZE), center + 1, center - 1, center - 2, center - 3},
			expectedSequences:      []int{1, 1, 1, 3, 0, 0, 0, 0},
		},
		{
			position:               center,
			opponentPositions:      []int{},
			currentPlayerPositions: []int{center, center - (board.SIZE) + 1, center - (board.SIZE) - 1, center + (board.SIZE) + 1, center + (board.SIZE) - 1},
			expectedSequences:      []int{0, 0, 0, 0, 1, 1, 1, 1},
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
		// Game.Gban.Tab = make([]int, board.SIZE*board.SIZE)
		Game.Goban.Tab = [board.TOT_SIZE]int{}
	}
}

func BenchmarkCheckSequence(b *testing.B) {
	Game.Goban.Tab = [board.TOT_SIZE]int{}
	Game.CurrentPlayer = Game.P1
	for i := 0; i < b.N; i++ {
		CheckSequence(i%(board.SIZE*board.SIZE), 1, &Game.Goban.Tab)
	}
}

func TestCompleteSequenceForPosition(t *testing.T) {
	// Initialize
	Game.Goban.Tab = [board.TOT_SIZE]int{}
	Game.CurrentPlayer = Game.P1
	center := (board.SIZE * board.SIZE) / 2
	if board.SIZE%2 == 0 {
		center += board.SIZE / 2
	}
	tables := []struct {
		position               int
		currentPlayerPositions []int
		expectedSequences      [][]int
	}{
		{
			position:               center,
			currentPlayerPositions: []int{center, center + 1, center - 1},
			expectedSequences:      [][]int{[]int{center, center - 1, center + 1}},
		},
		{
			position:               center,
			currentPlayerPositions: []int{center, center + board.SIZE, center - board.SIZE},
			expectedSequences:      [][]int{[]int{center, center - board.SIZE, center + board.SIZE}},
		},
		{
			position:               center,
			currentPlayerPositions: []int{center, center + (board.SIZE + 1), center - (board.SIZE + 1)},
			expectedSequences:      [][]int{[]int{center, center - (board.SIZE + 1), center + (board.SIZE + 1)}},
		},
		{
			position:               center,
			currentPlayerPositions: []int{center, center + (board.SIZE - 1), center - (board.SIZE - 1)},
			expectedSequences:      [][]int{[]int{center, center - (board.SIZE - 1), center + (board.SIZE - 1)}},
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
		Game.Goban.Tab = [board.TOT_SIZE]int{}
	}
}
