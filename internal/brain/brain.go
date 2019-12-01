package brain

import (
	"github.com/gogogomoku/gomoku/internal/board"
	"github.com/gogogomoku/gomoku/internal/player"

	bolt "github.com/gogogomoku/gomoku/internal/boltdb"
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
	P1                           *player.Player
	P2                           *player.Player
	Goban                        board.Board
	Status                       int16
	CurrentPlayer                *player.Player
	Turn                         int16
	SuggestedPosition            int16
	SuggestionTimer              int16
	Winner                       int16
	CacheEnabled                 bool
	CacheDB                      *bolt.BboltBucket
	InvalidMovesForCurrentPlayer []int16
}

func (round Round) GetCurrentOpponent() *player.Player {
	if round.CurrentPlayer.Id == round.P1.Id {
		return round.P2
	}
	return round.P1
}

var Game Round = Round{}

func init() {
	InitializeValues(0, 0)
}

func InitializeValues(aiStatus1 int16, aiStatus2 int16) {
	Game.P1 = player.CreatePlayer(1, MAXPIECES, aiStatus1)
	Game.P2 = player.CreatePlayer(2, MAXPIECES, aiStatus2)
	Game.Goban = board.Board{}
	Game.Goban.Tab = [board.TOT_SIZE]int16{}
	Game.Goban.Size = board.SIZE
	Game.Turn = 0
	Game.Status = NotStarted
	Game.Winner = 0
	Game.SuggestionTimer = 0
	Game.SuggestedPosition = 0
	Game.CurrentPlayer = nil
	Game.InvalidMovesForCurrentPlayer = []int16{}
}
