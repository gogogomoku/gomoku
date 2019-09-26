package brain

import (
	"reflect"
	"testing"

	"github.com/gogogomoku/gomoku/internal/board"
)

func TestCheckCapture(t *testing.T) {
	// Initialize
	Game.Goban.Tab = [board.TOT_SIZE]int16{}
	Game.CurrentPlayer = Game.P1
	center := int16((board.SIZE * board.SIZE) / 2)
	if board.SIZE%2 == 0 {
		center += board.SIZE / 2
	}
	tables := []struct {
		position                  int16
		opponentPositions         []int16
		currentPlayerPositions    []int16
		expectedCaptureDirections []int16
	}{
		{
			position:                  center,
			opponentPositions:         []int16{},
			currentPlayerPositions:    []int16{},
			expectedCaptureDirections: []int16{},
		},
		{
			position:                  center,
			opponentPositions:         []int16{center + 1, center + 2},
			currentPlayerPositions:    []int16{center, center + 3},
			expectedCaptureDirections: []int16{E},
		},
		{
			position:                  center,
			opponentPositions:         []int16{center - 1, center - 2},
			currentPlayerPositions:    []int16{center, center - 3},
			expectedCaptureDirections: []int16{W},
		},
		{
			position:                  center,
			opponentPositions:         []int16{center - board.SIZE, center - (2 * board.SIZE)},
			currentPlayerPositions:    []int16{center, center - (3 * board.SIZE)},
			expectedCaptureDirections: []int16{N},
		},
		{
			position:                  center,
			opponentPositions:         []int16{center + board.SIZE, center + (2 * board.SIZE)},
			currentPlayerPositions:    []int16{center, center + (3 * board.SIZE)},
			expectedCaptureDirections: []int16{S},
		},
	}
	for _, table := range tables {
		for _, v := range table.opponentPositions {
			Game.Goban.Tab[v] = 2
		}
		for _, v := range table.currentPlayerPositions {
			Game.Goban.Tab[v] = 1
		}
		captureDirections := checkCapture(table.position, &Game.Goban.Tab, Game.CurrentPlayer.Id)
		if !reflect.DeepEqual(table.expectedCaptureDirections, captureDirections) {

			t.Errorf("Wrong captureDirections for %d, expected %v, got %v", table.position, table.expectedCaptureDirections, captureDirections)

		}
		Game.Goban.Tab = [board.TOT_SIZE]int16{}
	}
}

// Broken, can probably delete
// func BenchmarkCheckCapture(b *testing.B) {
// 	Game.Goban.Tab = [board.TOT_SIZE]int16{}
// 	Game.CurrentPlayer = Game.P1
// 	for i := 0; i < b.N; i++ {
// 		checkCapture(int16(i)%(board.SIZE*board.SIZE), &Game.Goban.Tab, Game.CurrentPlayer.Id)
// 	}
// }
