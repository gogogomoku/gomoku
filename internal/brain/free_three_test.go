package brain

import (
	"math"
	"reflect"
	"testing"
)

func TestCheckSequenceForF3(t *testing.T) {
	tables := []struct {
		sequence []int
		playerId int
		expected []int
	}{
		//Simple
		{[]int{0, 1, 1, 1, 0}, 1, []int{0}},
		// Wrong player
		{[]int{0, 1, 1, 1, 0}, 2, []int{}},
		// Mixed players (enemy in sequence)
		{[]int{0, 1, 1, 2, 0}, 2, []int{}},
		{[]int{0, 1, 1, 2, 1, 0}, 1, []int{}},

		// should find 1 at position 0
		{[]int{0, 1, 1, 0, 1, 0}, 1, []int{0}},
		{[]int{0, 1, 0, 1, 1, 0}, 1, []int{0}},
		{[]int{0, 2, 2, 0, 2, 0}, 2, []int{0}},

		// should find none
		{[]int{1, 0, 1, 0, 1, 0}, 1, []int{}},
		{[]int{0, 1, 1, 0, 0, 0}, 1, []int{}},
		{[]int{1, 0, 1, 0, 0, 1}, 1, []int{}},
		{[]int{0, 1, 1, 1, 1, 0}, 1, []int{}},
		{[]int{0, 1, 0, 2, 1, 0}, 1, []int{}},

		// should find 1 at position 1
		{[]int{0, 0, 1, 1, 0, 1, 0}, 1, []int{1}},

		// should find 2
		{[]int{0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0}, 1, []int{1, 6}},
		{[]int{0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 0}, 1, []int{1, 5}},
		{[]int{0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0}, 1, []int{0, 5}},

		// should find 1
		{[]int{0, 0, 1, 1, 0, 1, 0, 1, 2, 1, 0}, 1, []int{1}},

		// should find 3 overlapping
		{[]int{0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0}, 1, []int{0, 3, 5}},
	}

	for _, table := range tables {
		actual := CheckSequenceForF3(table.sequence, table.playerId)
		if !reflect.DeepEqual(actual, table.expected) {
			t.Errorf("⛔️ F3 in sequence %v: got %v, expected %v\n", table.sequence, actual, table.expected)
		} else {
			// t.Logf("✅ for sequence %v\n", table.sequence)
		}
	}
}

func BenchmarkCheckSequenceForF3(b *testing.B) {
	b.StopTimer()
	playerId := 1
	seq := make([]int, math.MaxInt16*8)
	for i, _ := range seq {
		switch {
		case i%4 == 0:
			seq[i] = 1
		case i%5 == 0:
			seq[i] = 2
		case i > INT_MAX-20:
			seq[i] = 1
		}
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		CheckSequenceForF3(seq, playerId)
	}
}

func TestGetDoubleF3StartPos(t *testing.T) {
	t.Skipf("Let's do this later")
	// Game.Goban.Tab = [board.TOT_SIZE]int{}
	// Game.CurrentPlayer = Game.P1
	// // center := (board.SIZE * board.SIZE) / 2
	// // if board.SIZE%2 == 0 {
	// // 	center += board.SIZE / 2
	// // }
	// twoF3Boards := [1][][]int{
	// 	{
	// 		{0, 1, 0, 1, 1, 0, 0},
	// 		{0, 0, 0, 0, 0, 1, 0},
	// 		{0, 0, 0, 0, 0, 0, 0},
	// 		{0, 0, 0, 0, 0, 1, 0},
	// 		{0, 0, 0, 0, 0, 1, 0},
	// 		{0, 0, 0, 0, 0, 0, 0},
	// 	},
	// }
	// positions := [1]int{5}
	// expectedF3Starts := [1][]int{{5}}

	// // transpose
	// for i, miniboard := range twoF3Boards {
	// 	for j, line := range miniboard {
	// 		for k, val := range line {
	// 			Game.Goban.Tab[(board.SIZE*j)+k] = val
	// 		}
	// 	}
	// 	board.PrintBoard(Game.Goban.Tab, board.SIZE)
	// 	doubleF3Starts := GetDoubleF3StartPos(positions[i], Game.Goban.Tab, Game.CurrentPlayer.Id)
	// 	if !reflect.DeepEqual(doubleF3Starts, expectedF3Starts[i]) {
	// 		t.Errorf("⛔️ Double F3s in wrong start position(s): got %v, expected %v\n", doubleF3Starts, expectedF3Starts[i])
	// 	}
	// }
}
