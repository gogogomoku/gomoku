package brain

import (
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
			t.Logf("✅ for sequence %v\n", table.sequence)
		}
	}
}
