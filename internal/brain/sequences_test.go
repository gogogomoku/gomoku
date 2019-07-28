package brain

import (
	"reflect"
	"testing"

	"github.com/gogogomoku/gomoku/internal/board"
)

func TestCheckSequence(t *testing.T) {
	// Initialize
	GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	GameRound.CurrentPlayer = GameRound.P1
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
			GameRound.Goban.Tab[v] = 2
		}
		for _, v := range table.currentPlayerPositions {
			GameRound.Goban.Tab[v] = 1
		}
		sequenceLengths := CheckSequence(table.position, GameRound.CurrentPlayer.Id)

		if !reflect.DeepEqual(table.expectedSequences, sequenceLengths) {

			t.Errorf("Wrong sequenceLengths for %d, expected %v, got %v", table.position, table.expectedSequences, sequenceLengths)

		}
		GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	}
}

func BenchmarkCheckSequence(b *testing.B) {
	GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	GameRound.CurrentPlayer = GameRound.P1
	for i := 0; i < b.N; i++ {
		CheckSequence(i%(board.SIZE*board.SIZE), 1)
	}
}

func TestCompleteSequenceForPosition(t *testing.T) {
	// Initialize
	GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	GameRound.CurrentPlayer = GameRound.P1
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
			GameRound.Goban.Tab[v] = 1
		}
		completeSequences := CompleteSequenceForPosition(table.position, GameRound.CurrentPlayer.Id)
		if !reflect.DeepEqual(table.expectedSequences, completeSequences) {
			t.Errorf("Wrong completeSequences for %d, expected %v, got %v", table.position, table.expectedSequences, completeSequences)
		}
		GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	}
}
