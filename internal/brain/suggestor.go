package brain

import (
	"math/rand"

	"github.com/gogogomoku/gomoku/internal/board"
)

func SuggestMove() {
	possible := []int{}
	// best := []int{}
	for i := 0; i < (board.SIZE * board.SIZE); i++ {
		if GameRound.Goban.Tab[i] == 0 {
			if CheckValidMove(i) {
				// seq := CompleteSequenceForPosition(i, GameRound.CurrentPlayer.Id)
				// if len(seq) > 0 {
				// 	best = append(best, i)
				// }
				possible = append(possible, i)
			}
		}
	}
	ran := 0
	// if len(best) > 0 {
	// 	if len(best)-1 < 2 {
	// 		ran = 0
	// 	} else {
	// 		ran = int(rand.Intn(len(best) - 1))
	// 	}
	// 	GameRound.SuggestedPosition = best[ran]
	// } else {
	if len(possible)-1 < 2 {
		ran = 0
	} else {
		ran = int(rand.Intn(len(possible) - 1))
	}
	ran = int(rand.Intn(len(possible) - 1))
	GameRound.SuggestedPosition = possible[ran]
	// }
}
