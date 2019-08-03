package brain

import (
	"math/rand"

	"github.com/gogogomoku/gomoku/internal/board"
)

type Move struct {
	Position int
	Value    int
}

func getHeuristicValue(position int, id int) int {
	val := 0
	// Better value if creating longer sequences
	sequences := CompleteSequenceForPosition(position, id)
	for _, s := range sequences {
		switch len(s) {
		case 2:
			val += 10
		case 3:
			val += 30
		case 4:
			val += 100
		case 5:
			val += 1000
		}
	}
	// Better value if can capture
	capture := checkCapture(position)
	val += len(capture) * 20
	// Better value if close to opponent
	for dir := 0; dir < 8; dir++ {
		contact, edge := ReturnNextPiece(position, dir)
		if !edge && contact != 0 && contact != GameRound.CurrentPlayer.Id {
			val += 20
		}
	}
	// Better value if blocking a sequence from opponent
	opponent := 1
	if GameRound.CurrentPlayer.Id == 1 {
		opponent = 2
	}
	sequencesOp := CompleteSequenceForPosition(position, opponent)
	for _, s := range sequencesOp {
		switch len(s) {
		case 2:
			val += 10
		case 3:
			val += 30
		case 4:
			val += 100
		case 5:
			val += 1000
		}
	}

	return val
}

func SuggestMove() {
	possible := []Move{}
	for i := 0; i < (board.SIZE * board.SIZE); i++ {
		if GameRound.Goban.Tab[i] == 0 {
			if CheckValidMove(i) {
				heur := getHeuristicValue(i, GameRound.CurrentPlayer.Id)
				possible = append(possible, Move{Position: i, Value: heur})
			}
		}
	}
	ran := 0
	if len(possible) > 1 {
		ran = int(rand.Intn(len(possible) - 1))
		best := Move{Position: possible[ran].Position, Value: possible[ran].Value}
		for _, m := range possible {
			if m.Value > best.Value {
				best.Value = m.Value
				best.Position = m.Position
			}
		}
		GameRound.SuggestedPosition = best.Position
	} else {
		GameRound.SuggestedPosition = possible[ran].Position
	}

}
