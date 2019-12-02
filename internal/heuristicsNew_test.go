package brain

import (
	"reflect"
	"testing"
	// "github.com/gogogomoku/gomoku/internal/board"
)

func TestLaunchCheckSequence(t *testing.T) {

	// Free 4

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 0, 1, 1, 1, 1, 0, 0}, 1)
	expected := HeuristicScore{
		Free4: 1,
	}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	// Half 4

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 0, 0, 0, 1, 1, 1, 1}, 1)
	expected = HeuristicScore{
		Half4: 1,
	}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 0, 0, 0, 1, 1, 1, 1, 2}, 1)
	expected = HeuristicScore{
		Half4: 1,
	}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 0, 0, 2, 1, 1, 1, 1, 0}, 1)
	expected = HeuristicScore{
		Half4: 1,
	}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	// Flanked 4

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 0, 0, 2, 1, 1, 1, 1, 2}, 1)
	expected = HeuristicScore{}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 0, 0, 2, 1, 1, 1, 1}, 1)
	expected = HeuristicScore{}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{1, 1, 1, 1, 2, 0, 0, 0}, 1)
	expected = HeuristicScore{}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	// Half 3

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 0, 0, 0, 1, 1, 1}, 1)
	expected = HeuristicScore{
		Half3: 1,
	}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{1, 1, 1, 0}, 1)
	expected = HeuristicScore{
		Half3: 1,
	}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 0, 1, 2, 1, 1, 1, 0}, 1)
	expected = HeuristicScore{
		Half3: 1,
	}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	// Free 2

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 0, 0, 0, 1, 1, 0}, 1)
	expected = HeuristicScore{
		Free2: 1,
	}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 2, 2, 0, 1, 1, 0}, 1)
	expected = HeuristicScore{
		Free2: 1,
	}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	// Half 2

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 0, 0, 0, 1, 1, 2}, 1)
	expected = HeuristicScore{
		Half2: 1,
	}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 0, 0, 0, 1, 1}, 1)
	expected = HeuristicScore{
		Half2: 1,
	}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 0, 0, 0, 1, 1, 2, 1, 1, 0}, 1)
	expected = HeuristicScore{
		Half2: 2,
	}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}

	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkSequence([]int16{0, 0, 0, 2, 1, 1}, 1)
	expected = HeuristicScore{
		Half2: 0,
	}
	if !reflect.DeepEqual(scorePlayer, expected) {
		t.Errorf("Error in CheckSequence. Expected: %d, got: %d", expected, scorePlayer)
	}
}

func TestGetHeuristicValue(t *testing.T) {
	// center := (board.SIZE * board.SIZE) / 2
	// if board.SIZE%2 == 0 {
	// 	center += board.SIZE / 2
	// }
	// tab := [board.SIZE * board.SIZE]int16{}
	// tab[center-2] = 1
	// tab[center-1] = 1
	// tab[center] = 1
	// tab[center+1] = 1
	// score := getHeuristicValue(1, &tab, &[3]int16{})
	// if score != SEQ4_FREE_SCORE {
	// 	t.Errorf("Error in getHeuristicValue. Expected: %d, got: %d", score, SEQ4_FREE_SCORE)
	// }
}
