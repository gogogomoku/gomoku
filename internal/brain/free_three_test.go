package brain

import (
	"reflect"
	"testing"
)

func TestCheckSequenceForF3(t *testing.T) {
	tables := []struct {
		sequence []int16
		playerId int16
		expected []int16
	}{
		//Simple
		{[]int16{0, 1, 1, 1, 0}, 1, []int16{0}},
		// Wrong player
		{[]int16{0, 1, 1, 1, 0}, 2, []int16{}},
		// Mixed players (enemy in sequence)
		{[]int16{0, 1, 1, 2, 0}, 2, []int16{}},
		{[]int16{0, 1, 1, 2, 1, 0}, 1, []int16{}},

		// should find 1 at position 0
		{[]int16{0, 1, 1, 0, 1, 0}, 1, []int16{0}},
		{[]int16{0, 1, 0, 1, 1, 0}, 1, []int16{0}},
		{[]int16{0, 2, 2, 0, 2, 0}, 2, []int16{0}},

		// should find none
		{[]int16{1, 0, 1, 0, 1, 0}, 1, []int16{}},
		{[]int16{0, 1, 1, 0, 0, 0}, 1, []int16{}},
		{[]int16{1, 0, 1, 0, 0, 1}, 1, []int16{}},
		{[]int16{0, 1, 1, 1, 1, 0}, 1, []int16{}},
		{[]int16{0, 1, 0, 2, 1, 0}, 1, []int16{}},

		// should find 1 at position 1
		{[]int16{0, 0, 1, 1, 0, 1, 0}, 1, []int16{1}},

		// should find 2
		{[]int16{0, 0, 1, 1, 0, 1, 0, 1, 1, 1, 0}, 1, []int16{1, 6}},
		{[]int16{0, 0, 1, 1, 1, 0, 1, 0, 1, 1, 0}, 1, []int16{1, 5}},
		{[]int16{0, 1, 1, 1, 0, 0, 1, 0, 1, 1, 0}, 1, []int16{0, 5}},

		// should find 1
		{[]int16{0, 0, 1, 1, 0, 1, 0, 1, 2, 1, 0}, 1, []int16{1}},

		// should find 3 overlapping
		{[]int16{0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 0}, 1, []int16{0, 3, 5}},
	}

	for _, table := range tables {
		actual := CheckSequenceForF3(table.sequence, table.playerId)
		if !reflect.DeepEqual(actual, table.expected) {
			t.Errorf("⛔️ F3 in sequence %v: got %v, expected %v\n", table.sequence, actual, table.expected)
		} else {
			t.Logf("✅ for sequence %v\n", table.sequence)
		}
	}
}
