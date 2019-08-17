package brain

import (
	"testing"

	"github.com/gogogomoku/gomoku/internal/board"
)

func TestStartRound(t *testing.T) {
	t.Skip("Skipping, undecided about testing global state")
}

func TestCheckValidMove(t *testing.T) {
	// Game.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	Game.Goban.Tab = [board.TOT_SIZE]int{}
	Game.CurrentPlayer = Game.P1
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
			Game.Goban.Tab[v] = 2
		}
		for _, v := range table.currentPlayerPositions {
			Game.Goban.Tab[v] = 1
		}

		isValidMove := CheckValidMove(table.position, Game.Goban.Tab)
		if table.expectedIsValid != isValidMove {
			t.Errorf("position %d, valid: %t, expected: %t", table.position, isValidMove, table.expectedIsValid)
		}
		// Game.Goban.Tab = make([]int, board.SIZE*board.SIZE)
		Game.Goban.Tab = [board.TOT_SIZE]int{}
	}
}

func BenchmarkCheckValidMove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckValidMove(i%(board.SIZE*board.SIZE), Game.Goban.Tab)
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
	// nTiles = math.Pow(Game.Goban.Size, 2)
	// Game.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	Game.Goban.Tab = [board.TOT_SIZE]int{}
	Game.CurrentPlayer = Game.P1
	for i := 0; i < b.N; i++ {
		getNextIndexForDirection(i%(board.SIZE*board.SIZE), N)
	}
}

func TestReturnNextPiece(t *testing.T) {
	// Game.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	Game.Goban.Tab = [board.TOT_SIZE]int{}
	for i := range Game.Goban.Tab {
		Game.Goban.Tab[i] = i
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
		actualNextIndex, actualEdge := ReturnNextPiece(table.position, table.direction, &Game.Goban.Tab)
		if table.expectedEdge != actualEdge {
			t.Errorf("Wrong edge at index %d, expected %t, got %t", table.position, table.expectedEdge, actualEdge)
		}
		if table.expectedNextIndex != actualNextIndex {
			t.Errorf("Wrong nextIndex, expected %d, got %d", table.expectedNextIndex, actualNextIndex)
		}
	}
}

func BenchmarkReturnNextPiece(b *testing.B) {
	// Game.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	Game.Goban.Tab = [board.TOT_SIZE]int{}
	Game.CurrentPlayer = Game.P1
	for i := 0; i < b.N; i++ {
		ReturnNextPiece((board.SIZE * board.SIZE), NE, &Game.Goban.Tab)
	}
}
