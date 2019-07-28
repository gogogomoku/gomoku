package brain

import (
	"reflect"
	"testing"

	"github.com/gogogomoku/gomoku/internal/board"
)

func TestCheckCapture(t *testing.T) {
	// Initialize
	GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	GameRound.CurrentPlayer = GameRound.P1
	center := (board.SIZE * board.SIZE) / 2
	if board.SIZE%2 == 0 {
		center += board.SIZE / 2
	}
	tables := []struct {
		position                  int
		opponentPositions         []int
		currentPlayerPositions    []int
		expectedCaptureDirections []int
	}{
		{
			position:                  center,
			opponentPositions:         []int{},
			currentPlayerPositions:    []int{},
			expectedCaptureDirections: []int{},
		},
		{
			position:                  center,
			opponentPositions:         []int{center + 1, center + 2},
			currentPlayerPositions:    []int{center, center + 3},
			expectedCaptureDirections: []int{E},
		},
		{
			position:                  center,
			opponentPositions:         []int{center - 1, center - 2},
			currentPlayerPositions:    []int{center, center - 3},
			expectedCaptureDirections: []int{W},
		},
		{
			position:                  center,
			opponentPositions:         []int{center - board.SIZE, center - (2 * board.SIZE)},
			currentPlayerPositions:    []int{center, center - (3 * board.SIZE)},
			expectedCaptureDirections: []int{N},
		},
		{
			position:                  center,
			opponentPositions:         []int{center + board.SIZE, center + (2 * board.SIZE)},
			currentPlayerPositions:    []int{center, center + (3 * board.SIZE)},
			expectedCaptureDirections: []int{S},
		},
	}
	for _, table := range tables {
		for _, v := range table.opponentPositions {
			GameRound.Goban.Tab[v] = 2
		}
		for _, v := range table.currentPlayerPositions {
			GameRound.Goban.Tab[v] = 1
		}
		captureDirections := checkCapture(table.position)
		if !reflect.DeepEqual(table.expectedCaptureDirections, captureDirections) {

			t.Errorf("Wrong captureDirections for %d, expected %v, got %v", table.position, table.expectedCaptureDirections, captureDirections)

		}
		GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	}
}

func BenchmarkCheckCapture(b *testing.B) {
	GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	GameRound.CurrentPlayer = GameRound.P1
	for i := 0; i < b.N; i++ {
		checkCapture(i % (board.SIZE * board.SIZE))
	}
}
