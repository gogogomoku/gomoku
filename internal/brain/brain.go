package brain

import (
	"gomoku/internal/board"
	"gomoku/internal/player"
)

const SIZE = 19
const MAXPIECES = 50

const (
	NotStarted = iota
	Running
	Concluded
)

type Round struct {
	P1            *player.Player
	P2            *player.Player
	Goban         board.Board
	Status        int
	CurrentPlayer *player.Player
}

var GameRound Round = Round{}

func init() {
	GameRound.P1 = player.CreatePlayer(1, MAXPIECES)
	GameRound.P2 = player.CreatePlayer(2, MAXPIECES)
	GameRound.Goban = board.Board{}
	GameRound.Goban.Tab = make([]byte, SIZE*SIZE)
	GameRound.Goban.Size = SIZE
	GameRound.Status = NotStarted
}
