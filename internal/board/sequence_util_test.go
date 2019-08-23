package board

import (
	// "reflect"
	"testing"

)

func TestGetDiagonalNWSESequences(t *testing.T) {
	playerId := 1
	tab := [TOT_SIZE]int{}

	for i := 0; i < TOT_SIZE; i++ {
		tab[i] = i
	}

	sequences := GetDiagonalNWSESequences(playerId, &tab)
	t.Logf("Sequences: %#v", sequences)

	_ = sequences
	_ = playerId
	_ = tab

	t.Skip("Skip skip, hooray!")
}