package brain

import (
	"github.com/gogogomoku/gomoku/internal/board"
	"github.com/gogogomoku/gomoku/internal/player"
)

const MAXPIECES = 150

const (
	NotStarted = iota
	Running
	Concluded
)

const (
	N = iota
	S
	E
	W
	NE
	NW
	SE
	SW
)

type Round struct {
	P1                *player.Player
	P2                *player.Player
	Goban             board.Board
	Status            int
	CurrentPlayer     *player.Player
	Turn              int
	SuggestedPosition int
	Winner            int
}

var Game Round = Round{}

func init() {
	InitializeValues()
}

func InitializeValues() {
	Game.P1 = player.CreatePlayer(1, MAXPIECES)
	Game.P2 = player.CreatePlayer(2, MAXPIECES)
	Game.Goban = board.Board{}
	Game.Goban.Tab = make([]int, board.SIZE*board.SIZE)
	Game.Goban.Size = board.SIZE
	Game.Turn = 0
	Game.Status = NotStarted
	Game.Winner = 0
}
