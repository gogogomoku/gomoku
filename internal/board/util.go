package board

func MakeTab(player1Positions []int16, player2Positions []int16) *[TOT_SIZE]int16 {
	tab := [TOT_SIZE]int16{}
	for _, pos := range player1Positions {
		tab[pos] = 1
	}
	for _, pos := range player2Positions {
		tab[pos] = 2
	}
	// PrintBoard(tab, 19)
	return &tab
}
