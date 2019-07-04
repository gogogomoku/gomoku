package suggestor

import "gomoku/internal/board"
import "math/rand"

func SuggestMove() int {
	suggestion := int(rand.Intn(board.SIZE * board.SIZE))
	return suggestion
}
