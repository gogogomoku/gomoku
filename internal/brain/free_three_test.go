package brain

import (
	"reflect"
	"testing"

	"github.com/gogogomoku/gomoku/internal/board"
)

func BenchmarkTestCheckSequenceForF3(b *testing.B) {
	sequences := make([][]int16, 4)

	// 467 ns/loop
	for i := 0; i < b.N; i++ {
		sequences[0] = []int16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		sequences[1] = []int16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		sequences[2] = []int16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		sequences[3] = []int16{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		for j := range sequences {
			CheckSequenceForF3(sequences[j], 1)
		}
	}
}

func TestGetNESWSeqIndex(t *testing.T) {
	tables := []struct {
		positions   []int16
		expectedSeq int16
	}{
		{
			positions:   []int16{360},
			expectedSeq: 0,
		},
		{
			positions:   []int16{341, 359},
			expectedSeq: 1,
		},
		{
			positions:   []int16{322, 340, 358},
			expectedSeq: 2,
		},
		{
			positions:   []int16{303, 321, 339, 357},
			expectedSeq: 3,
		},
		{
			positions:   []int16{284, 302, 320, 338},
			expectedSeq: 4,
		},
		{
			positions:   []int16{37, 55, 73, 127, 91, 163, 199, 235},
			expectedSeq: 17,
		},
		{
			positions:   []int16{15, 33, 51, 69, 87, 105, 123, 141, 159, 177, 195, 213, 231, 249, 267, 285},
			expectedSeq: 21,
		},
		{
			positions:   []int16{16, 34, 52, 88, 106, 124, 178, 232, 286},
			expectedSeq: 20,
		},
		{
			positions:   []int16{14, 32, 50, 68, 86, 104, 122, 140, 158, 176, 194, 212, 230, 248, 266},
			expectedSeq: 22,
		},
		{
			positions:   []int16{13, 31, 49, 67, 85, 103, 121, 139, 157, 175, 193, 211, 229, 247},
			expectedSeq: 23,
		},
		{
			positions:   []int16{12, 30, 48, 66, 84, 102, 120, 138, 156, 174, 192, 210, 228},
			expectedSeq: 24,
		},
		{
			positions:   []int16{208, 226, 244, 262, 280, 298, 316, 334, 352},
			expectedSeq: 8,
		},
		{
			positions:   []int16{75, 93, 111, 129, 147, 165, 183, 201, 219, 237, 255, 273, 291, 309, 327, 345},
			expectedSeq: 15,
		},
	}
	for _, tb := range tables {
		for _, pos := range tb.positions {
			actual := getNESWSeqIndex(pos)
			if actual != tb.expectedSeq {
				t.Errorf("🛑 Wrong NESW sequence for position %d, expect: %d, got: %d\n", pos, tb.expectedSeq, actual)
			} else {
				print("✓")
			}
		}
	}
}

func TestGetNWSESeqIndex(t *testing.T) {
	tables := []struct {
		positions   []int16
		expectedSeq int16
	}{
		{
			positions:   []int16{0, 20, 40, 60, 80, 100, 120, 140, 160, 180, 200, 220, 240, 260, 280, 300, 320, 340, 360},
			expectedSeq: 18,
		},
		{
			positions:   []int16{342},
			expectedSeq: 0,
		},
		{
			positions:   []int16{18},
			expectedSeq: 36,
		},
	}
	for _, tb := range tables {
		for _, pos := range tb.positions {
			actual := getNWSESeqIndex(pos)
			if actual != tb.expectedSeq {
				t.Errorf("🛑 Wrong NWSE sequence for position %d, expect: %d, got: %d\n", pos, tb.expectedSeq, actual)
			} else {
				print("✓")
			}
		}
	}
}

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
			print("✓")
		}
	}
}

func TestCheck2F3s(t *testing.T) {
	tables := []struct {
		playerId int16
		posGoban int16
		tab      *[board.TOT_SIZE]int16
		expected bool
	}{
		{
			playerId: 1,
			posGoban: 1,
			tab:      board.MakeTab([]int16{}, []int16{}),
			expected: false,
		},
		{
			// F3 in first column starting at tab[0] (len 5)
			playerId: 1,
			posGoban: 19,
			tab:      board.MakeTab([]int16{19 + 19, 19 + 19 + 19}, []int16{}),
			expected: false,
		},
		{
			// F3 in first column starting at tab[0] (len 6)
			playerId: 1,
			posGoban: 19,
			tab:      board.MakeTab([]int16{19 + 19 + 19, 19 + 19 + 19 + 19}, []int16{}),
			expected: false,
		},
		{
			// F3 in first column starting at tab[19] (len 6)
			playerId: 1,
			posGoban: 19 * 2,
			tab:      board.MakeTab([]int16{19 * 4, 19 * 5}, []int16{}),
			expected: false,
		},
		{
			// F3 in first row starting at tab[0] (len 5)
			playerId: 1,
			posGoban: 1,
			tab:      board.MakeTab([]int16{2, 3}, []int16{}),
			expected: false,
		},
		{
			// F3 in second row starting at tab[20] (len 5)
			// and in third column starting at tab[2] (len 5)
			playerId: 1,
			posGoban: 21,
			tab:      board.MakeTab([]int16{22, 23, 40, 59}, []int16{}),
			expected: true,
		},
		{
			// F3 in second row starting at tab[20] (len 5)
			// and in third column starting at tab[2] (len 5)
			// but enemy is mixed in, so should return false
			playerId: 1,
			posGoban: 21,
			tab:      board.MakeTab([]int16{22, 23, 40, 59}, []int16{20}),
			expected: false,
		},
		{
			// F3 in 18th NWSE diagonal, starting at pos 0, player position 20
			playerId: 1,
			posGoban: 20,
			tab:      board.MakeTab([]int16{40, 60}, []int16{}),
			expected: false,
		},
		{
			// F3 in 18th NWSE diagonal, starting at pos 20, player position 40, len 6
			playerId: 1,
			posGoban: 40,
			tab:      board.MakeTab([]int16{60, 100}, []int16{}),
			expected: false,
		},
		{
			// F3 in 10th NWSE diagonal, starting at pos 29, player position 69, len 6
			playerId: 1,
			posGoban: 69,
			tab:      board.MakeTab([]int16{49, 109}, []int16{}),
			expected: false,
		},
		{
			// F3 in 18th NESW diagonal, starting at pos 18, player position 36, len 5
			// F3 in 2nd row, starting at pos 33, len 5
			playerId: 1,
			posGoban: 36,
			tab:      board.MakeTab([]int16{54, 72, 34, 35}, []int16{}),
			expected: true,
		},
		{
			// F3 in 18th NESW diagonal, starting at pos 18, player position 36, len 5
			// Should return false, since we stop looking before NESW check if we don't already have one F3
			playerId: 1,
			posGoban: 36,
			tab:      board.MakeTab([]int16{54, 72}, []int16{}),
			expected: false,
		},
		{
			// Creates 2 F3s, vertical and NWSE diagonal
			playerId: 1,
			posGoban: 36,
			tab:      board.MakeTab([]int16{54, 72, 74, 93}, []int16{}),
			expected: true,
		},
		{
			// player creates F3 NS, also creates shape of F3 with opponent mixed in
			// Should return false
			playerId: 1,
			posGoban: 36,
			tab:      board.MakeTab([]int16{54, 72}, []int16{74, 93}),
			expected: false,
		},
		{
			// Create two diagonal F3s
			// Should return true
			playerId: 1,
			posGoban: 177,
			tab:      board.MakeTab([]int16{157, 197, 159, 195}, []int16{}),
			expected: true,
		},
		{
			// Create NWSE and NS F3s
			// Should return true
			playerId: 1,
			posGoban: 300,
			tab:      board.MakeTab([]int16{280, 260, 220, 281, 262}, []int16{}),
			expected: true,
		},
		{
			// Create NWSE and EW F3s -> true
			playerId: 1,
			posGoban: 157,
			tab:      board.MakeTab([]int16{159, 160, 177, 197}, []int16{}),
			expected: true,
		},
		{
			// Create NESW and EW F3s -> true
			playerId: 1,
			posGoban: 157,
			tab:      board.MakeTab([]int16{159, 160, 139, 121}, []int16{}),
			expected: true,
		},
		{
			// Create NESW and EW F3s again -> true
			playerId: 1,
			posGoban: 157,
			tab:      board.MakeTab([]int16{158, 159, 139, 121}, []int16{}),
			expected: true,
		},
		{
			// Create NESW and EW F3s again, again -> true
			playerId: 1,
			posGoban: 157,
			tab:      board.MakeTab([]int16{158, 159, 121, 103}, []int16{}),
			expected: true,
		},
		{
			// Create NS and EW F3s -> true
			playerId: 1,
			posGoban: 157,
			tab:      board.MakeTab([]int16{176, 214, 156, 158}, []int16{}),
			expected: true,
		},
		{
			// Create NS and EW F3s -> true
			playerId: 1,
			posGoban: 100,
			tab:      board.MakeTab([]int16{119, 138, 82, 118}, []int16{}),
			expected: true,
		},
		{
			// Create two NS -> false (not illegal)
			playerId: 1,
			posGoban: 100,
			tab:      board.MakeTab([]int16{43, 62, 138, 157}, []int16{}),
			expected: false,
		},
	}

	for _, table := range tables {
		actual := Check2F3s(table.playerId, table.posGoban, table.tab)
		if !reflect.DeepEqual(actual, table.expected) {
			t.Errorf("⛔️ F3 in F3s at position %v.\nActual: %v; Expected: %v\n", table.posGoban, actual, table.expected)
		} else {
			print("✓")
		}
	}
}

func TestCheckAxisForF3(t *testing.T) {
	tables := []struct {
		axis        int16
		len         int16
		playerId    int16
		seqOffset   int16
		startGobanI int16
		step        int16
		tab         *[board.TOT_SIZE]int16
		whichSeq    int16
		expectedF3  bool
	}{
		{
			axis:        NESW,
			len:         19,
			playerId:    1,
			seqOffset:   8,
			startGobanI: 18,
			step:        18,
			tab:         board.MakeTab([]int16{108, 126, 144}, []int16{}),
			whichSeq:    18,
			expectedF3:  true,
		},
		{
			axis:        EW,
			len:         19,
			playerId:    1,
			seqOffset:   3,
			startGobanI: 95,
			step:        1,
			tab:         board.MakeTab([]int16{98, 100, 101}, []int16{}),
			whichSeq:    5,
			expectedF3:  true,
		},
	}

	for _, tb := range tables {
		actual := checkAxisForF3(tb.axis, tb.len, tb.playerId, tb.seqOffset, tb.startGobanI, tb.step, tb.tab, tb.whichSeq)
		if actual != tb.expectedF3 {
			t.Errorf("⛔️ Expected to find F3: %v, actual: %v\n", tb.expectedF3, actual)
		} else {
			print("✓")
		}
	}
}

func BenchmarkCheck2F3s(b *testing.B) {
	blankBoard := board.MakeTab([]int16{}, []int16{})

	allP1s := make([]int16, board.TOT_SIZE)
	for i := range allP1s {
		allP1s[i] = int16(i)
	}
	boardWithManyP1s := board.MakeTab(allP1s, []int16{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 164 ns/op
		Check2F3s(1, board.TOT_SIZE/2, blankBoard)
		Check2F3s(1, board.TOT_SIZE/2, boardWithManyP1s)
	}
}
