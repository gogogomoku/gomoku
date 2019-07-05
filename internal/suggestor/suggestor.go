package suggestor

import "github.com/gogogomoku/gomoku/internal/board"
import "math/rand"

func SuggestMove() int8 {
	suggestion := int8(rand.Intn(board.SIZE * board.SIZE))
	return suggestion
}