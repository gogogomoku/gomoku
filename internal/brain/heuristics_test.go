package brain

import (
	"testing"

	"github.com/gogogomoku/gomoku/internal/board"
)

func TestLaunchCheckSequence(t *testing.T) {
	score := checkSequence([]int{0, 0, 1, 1, 1, 1, 0, 0}, 1)
	if score != SEQ4_FREE_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", SEQ4_FREE_SCORE, score)
	}
	score = checkSequence([]int{0, 0, 1, 1, 1, 0, 0}, 1)
	if score != F3_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", F3_SCORE, score)
	}
	score = checkSequence([]int{0, 0, 1, 1, 1, 2, 0}, 1)
	if score != SEQ3_BLOCKED1_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", SEQ3_BLOCKED1_SCORE, score)
	}
	score = checkSequence([]int{0, 2, 1, 1, 1, 1, 0}, 1)
	if score != SEQ4_BLOCKED1_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", SEQ4_BLOCKED1_SCORE, score)
	}
	score = checkSequence([]int{0, 2, 1, 1, 1, 1, 1}, 1)
	if score != WIN_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", WIN_SCORE, score)
	}
	score = checkSequence([]int{0, 2, 1, 1, 1, 1, 1, 2}, 1)
	if score != WIN_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", WIN_SCORE, score)
	}
	score = checkSequence([]int{0, 2, 1, 1, 1, 1, 1, 0}, 1)
	if score != WIN_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", WIN_SCORE, score)
	}
	score = checkSequence([]int{0, 2, 1, 1, 1, 1, 0}, 1)
	if score != SEQ4_BLOCKED1_SCORE {
		t.Errorf("Error in checkSequence. Expected: %d, got: %d", SEQ4_BLOCKED1_SCORE, score)
	}
}

func TestGetHeuristicValue(t *testing.T) {
	center := (board.SIZE * board.SIZE) / 2
	if board.SIZE%2 == 0 {
		center += board.SIZE / 2
	}
	tab := [board.SIZE * board.SIZE]int{}
	tab[center-2] = 1
	tab[center-1] = 1
	tab[center] = 1
	tab[center+1] = 1
	score := getHeuristicValue(0, 1, &tab)
	if score != SEQ4_FREE_SCORE {
		t.Errorf("Error in getHeuristicValue. Expected: %d, got: %d", score, SEQ4_FREE_SCORE)
	}
}