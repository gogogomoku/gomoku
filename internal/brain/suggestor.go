package brain

import "github.com/gogogomoku/gomoku/internal/board"
import "math/rand"

func SuggestMove() int {
	suggestion := int(rand.Intn(board.SIZE * board.SIZE))
	return suggestion
}

func getSuggestion() {
	for {
		GameRound.SuggestedPosition = SuggestMove()
		if checkValidMove(GameRound.SuggestedPosition) {
			break
		}
	}
}
