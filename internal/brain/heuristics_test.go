package brain

import (
	"testing"

	"github.com/gogogomoku/gomoku/internal/board"
)

func BenchmarkCheckDiagonalNWSESequences(b *testing.B) {
	blankBoard := board.MakeTab([]int16{}, []int16{})

	allP1s := make([]int16, board.TOT_SIZE)
	for i := range allP1s {
		allP1s[i] = int16(i)
	}
	boardWithManyP1s := board.MakeTab(allP1s, []int16{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		checkDiagonalNWSESequences(1, blankBoard)       // 2667 ns/op
		checkDiagonalNWSESequences(1, boardWithManyP1s) // 3811 ns/op
	}
}

func TestLaunchCheckSequence(t *testing.T) {
	score := checkSequence([]int16{0, 0, 1, 1, 1, 1, 0, 0}, 1)
	if score != SEQ4_FREE_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", SEQ4_FREE_SCORE, score)
	}
	score = checkSequence([]int16{0, 0, 1, 1, 1, 0, 0}, 1)
	if score != F3_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", F3_SCORE, score)
	}
	score = checkSequence([]int16{0, 0, 1, 1, 1, 2, 0}, 1)
	if score != SEQ3_BLOCKED1_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", SEQ3_BLOCKED1_SCORE, score)
	}
	score = checkSequence([]int16{0, 2, 1, 1, 1, 1, 0}, 1)
	if score != SEQ4_BLOCKED1_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", SEQ4_BLOCKED1_SCORE, score)
	}
	score = checkSequence([]int16{0, 2, 1, 1, 1, 1, 1}, 1)
	if score != WIN_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", WIN_SCORE, score)
	}
	score = checkSequence([]int16{0, 2, 1, 1, 1, 1, 1, 2}, 1)
	if score != WIN_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", WIN_SCORE, score)
	}
	score = checkSequence([]int16{0, 2, 1, 1, 1, 1, 1, 0}, 1)
	if score != WIN_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", WIN_SCORE, score)
	}
	score = checkSequence([]int16{0, 2, 1, 1, 1, 1, 0}, 1)
	if score != SEQ4_BLOCKED1_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", SEQ4_BLOCKED1_SCORE, score)
	}
	score = checkSequence([]int16{0, 1, 1, 0, 1, 1, 0}, 1)
	if score != SEQ4_FREE_SCORE+2*SEQ2_FREE_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", SEQ4_BROKEN+2*SEQ2_FREE_SCORE, score)
	}
	score = checkSequence([]int16{2, 1, 1, 0, 1, 1, 2}, 1)
	if score != SEQ4_BROKEN+2*SEQ2_BLOCKED1_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", SEQ4_BROKEN+2*SEQ2_BLOCKED1_SCORE, score)
	}
	score = checkSequence([]int16{2, 1, 1, 1, 0, 1, 2}, 1)
	if score != SEQ4_BROKEN+SEQ3_BLOCKED1_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", SEQ4_BROKEN+SEQ3_BLOCKED1_SCORE, score)
	}
	score = checkSequence([]int16{2, 1, 0, 1, 1, 1, 2}, 1)
	if score != SEQ4_BROKEN+SEQ3_BLOCKED1_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", SEQ4_BROKEN+SEQ3_BLOCKED1_SCORE, score)
	}
}

func TestGetHeuristicValue(t *testing.T) {
	center := (board.SIZE * board.SIZE) / 2
	if board.SIZE%2 == 0 {
		center += board.SIZE / 2
	}
	tab := [board.SIZE * board.SIZE]int16{}
	tab[center-2] = 1
	tab[center-1] = 1
	tab[center] = 1
	tab[center+1] = 1
	score := getHeuristicValue(1, &tab, &[3]int16{})
	if score != SEQ4_FREE_SCORE {
		t.Errorf("Error in getHeuristicValue. Expected: %d, got: %d", score, SEQ4_FREE_SCORE)
	}
}

func BenchmarkGetHeuristicValue(b *testing.B) {
	blankBoard := board.MakeTab([]int16{}, []int16{})

	allP1s := make([]int16, board.TOT_SIZE)
	for i := range allP1s {
		allP1s[i] = int16(i)
	}
	boardWithManyP1s := board.MakeTab(allP1s, []int16{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		getHeuristicValue(1, blankBoard, &[3]int16{})       // 15214/op
		getHeuristicValue(1, boardWithManyP1s, &[3]int16{}) // 14851 ns/op
	}
}
