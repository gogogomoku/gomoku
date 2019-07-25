package brain

import (
	"gomoku/internal/board"
	"reflect"
	"testing"
)

func TestStartRound(t *testing.T) {
	t.Skip("Skipping, undecided about testing global state")
}

func TestGetNextIndexForDirection(t *testing.T) {
	tables := []struct {
		position          int
		direction         int
		expectedEdge      bool
		expectedNextIndex int
	}{
		// NW corner
		{0, N, true, -42},
		{0, S, false, board.SIZE},
		{0, E, false, 1},
		{0, W, true, -42},

		// NE corner
		{board.SIZE - 1, N, true, -42},
		{board.SIZE - 1, S, false, (board.SIZE - 1) + board.SIZE},
		{board.SIZE - 1, E, true, -42},
		{board.SIZE - 1, W, false, board.SIZE - 2},

		//etc...
	}
	for _, table := range tables {
		actualNextIndex, actualEdge := getNextIndexForDirection(table.position, table.direction)
		if table.expectedEdge != actualEdge {
			t.Errorf("Wrong edge, expected %t, got %t", table.expectedEdge, actualEdge)
		}
		if table.expectedNextIndex != actualNextIndex {
			t.Errorf("Wrong nextIndex, expected %d, got %d", table.expectedNextIndex, actualNextIndex)
		}
	}
}

func TestReturnNextPiece(t *testing.T) {
	GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	for i := range GameRound.Goban.Tab {
		GameRound.Goban.Tab[i] = i
	}
	center := (board.SIZE * board.SIZE) / 2
	if board.SIZE%2 == 0 {
		center += board.SIZE / 2
	}
	tables := []struct {
		position          int
		direction         int
		expectedEdge      bool
		expectedNextIndex int
	}{
		// NW corner
		{0, N, true, -42},
		{0, S, false, board.SIZE},
		{0, E, false, 1},
		{0, W, true, -42},

		// NE corner
		{board.SIZE - 1, N, true, -42},
		{board.SIZE - 1, S, false, (board.SIZE - 1) + board.SIZE},
		{board.SIZE - 1, E, true, -42},
		{board.SIZE - 1, W, false, board.SIZE - 2},

		//etc...
		{center, N, false, center - board.SIZE},
		{center, S, false, center + board.SIZE},
		{center, E, false, center + 1},
		{center, W, false, center - 1},
	}
	for _, table := range tables {
		actualNextIndex, actualEdge := ReturnNextPiece(table.position, table.direction)
		if table.expectedEdge != actualEdge {
			t.Errorf("Wrong edge at index %d, expected %t, got %t", table.position, table.expectedEdge, actualEdge)
		}
		if table.expectedNextIndex != actualNextIndex {
			t.Errorf("Wrong nextIndex, expected %d, got %d", table.expectedNextIndex, actualNextIndex)
		}
	}
}

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
