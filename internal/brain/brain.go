package brain

import (
	"gomoku/internal/board"
	"gomoku/internal/player"
)

const (
	NotStarted = iota
	P1Turn
	P2Turn
	Concluded
)

type Round struct {
	P1     *player.Player
	P2     *player.Player
	Goban  *board.Board
	Status int
}

func Initialize() *Round {
	p1 := player.CreatePlayer(1, 45)
	p2 := player.CreatePlayer(2, 45)
	tab := make([]byte, 19*19)
	// fmt.Printf("%+v", board.Board)
	goban := board.Board{Tab: tab, Size: 19}

	round := Round{P1: p1, P2: p2, Goban: &goban, Status: NotStarted}
	return &round
}
