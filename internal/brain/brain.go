package brain

import (
	"github.com/gogogomoku/gomoku/internal/board"
	"github.com/gogogomoku/gomoku/internal/player"
)

const MAXPIECES = 50

const (
	NotStarted = iota
	Running
	Concluded
)

type Round struct {
	P1                *player.Player
	P2                *player.Player
	Goban             board.Board
	Status            int
	CurrentPlayer     *player.Player
	Turn              int
	SuggestedPosition int8
}

var GameRound Round = Round{}

func init() {
	GameRound.P1 = player.CreatePlayer(1, MAXPIECES)
	GameRound.P2 = player.CreatePlayer(2, MAXPIECES)
	GameRound.Goban = board.Board{}
	GameRound.Goban.Tab = make([]int8, board.SIZE*board.SIZE)
	GameRound.Goban.Size = board.SIZE
	GameRound.Turn = 0
	GameRound.Status = NotStarted
}
