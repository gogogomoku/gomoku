package brain

import (
	"github.com/gogogomoku/gomoku/internal/board"
	"reflect"
	"testing"
)

func TestStartRound(t *testing.T) {
	t.Skip("Skipping, undecided about testing global state")
}

func TestCheckValidMove(t *testing.T) {
	GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	GameRound.CurrentPlayer = GameRound.P1
	center := (board.SIZE * board.SIZE) / 2
	if board.SIZE%2 == 0 {
		center += board.SIZE / 2
	}

	tables := []struct {
		position               int
		currentPlayerPositions []int
		opponentPositions      []int
		expectedIsValid        bool
	}{
		// Valid
		{0, []int{1}, []int{2}, true},
		{1, []int{0, 2}, []int{3}, true},
		{center, []int{center - 1}, []int{center + 1}, true},
		{center, []int{}, []int{}, true},

		// Invalid
		{0, []int{0}, []int{}, false},
		{0, []int{}, []int{0}, false},
		{-1, []int{}, []int{}, false},
		{board.SIZE * board.SIZE, []int{}, []int{}, false},
	}

	for _, table := range tables {
		for _, v := range table.opponentPositions {
			GameRound.Goban.Tab[v] = 2
		}
		for _, v := range table.currentPlayerPositions {
			GameRound.Goban.Tab[v] = 1
		}

		isValidMove := checkValidMove(table.position)
		if table.expectedIsValid != isValidMove {
			t.Errorf("position %d, valid: %t, expected: %t", table.position, isValidMove, table.expectedIsValid)
		}
		GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	}
}

func BenchmarkCheckValidMove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		checkValidMove(i % (board.SIZE * board.SIZE))
	}
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

func BenchmarkGetNextIndexForDirection(b *testing.B) {
	// nTiles = math.Pow(GameRound.Goban.Size, 2)
	GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	GameRound.CurrentPlayer = GameRound.P1
	for i := 0; i < b.N; i++ {
		getNextIndexForDirection(i%(board.SIZE*board.SIZE), N)
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

func BenchmarkReturnNextPiece(b *testing.B) {
	GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	GameRound.CurrentPlayer = GameRound.P1
	for i := 0; i < b.N; i++ {
		ReturnNextPiece((board.SIZE * board.SIZE), NE)
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

func BenchmarkCheckCapture(b *testing.B) {
	GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	GameRound.CurrentPlayer = GameRound.P1
	for i := 0; i < b.N; i++ {
		checkCapture(i % (board.SIZE * board.SIZE))
	}
}

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
		sequenceLengths := checkSequence(table.position, GameRound.CurrentPlayer.Id)

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
		checkSequence(i%(board.SIZE*board.SIZE), 1)
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
		completeSequences := completeSequenceForPosition(table.position, GameRound.CurrentPlayer.Id)
		if !reflect.DeepEqual(table.expectedSequences, completeSequences) {
			t.Errorf("Wrong completeSequences for %d, expected %v, got %v", table.position, table.expectedSequences, completeSequences)
		}
		GameRound.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	}
}
