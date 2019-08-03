package brain

import (
	"fmt"

	"github.com/gogogomoku/gomoku/internal/board"
	"github.com/gogogomoku/gomoku/internal/player"
)

func StartRound() {
	GameRound.Status = Running
	player.ResetPlayers(GameRound.P1, GameRound.P2, MAXPIECES)
	GameRound.CurrentPlayer = GameRound.P1
	SuggestMove()
	HandleMove(GameRound.CurrentPlayer.Id, GameRound.SuggestedPosition)
}

func CheckValidMove(position int) bool {
	return bool(position >= 0 && position <= (board.SIZE*board.SIZE)-1 && GameRound.Goban.Tab[position] == 0)
}

func getNextIndexForDirection(position int, direction int) (nextIndex int, edge bool) {
	possibleDirection := [4]bool{true, true, true, true}
	// First row
	if position < board.SIZE {
		possibleDirection[N] = false
	}
	// Last row
	if position >= (board.SIZE * (board.SIZE - 1)) {
		possibleDirection[S] = false
	}
	// East column
	if position%board.SIZE == (board.SIZE - 1) {
		possibleDirection[E] = false
	}
	// West column
	if position%board.SIZE == 0 {
		possibleDirection[W] = false
	}
	switch {
	case direction == N && possibleDirection[N]:
		return position - board.SIZE, false
	case direction == S && possibleDirection[S]:
		return position + board.SIZE, false
	case direction == E && possibleDirection[E]:
		return position + 1, false
	case direction == W && possibleDirection[W]:
		return position - 1, false
	case direction == NE && possibleDirection[N] && possibleDirection[E]:
		return position - (board.SIZE - 1), false
	case direction == NW && possibleDirection[N] && possibleDirection[W]:
		return position - (board.SIZE + 1), false
	case direction == SE && possibleDirection[S] && possibleDirection[E]:
		return position + (board.SIZE + 1), false
	case direction == SW && possibleDirection[S] && possibleDirection[W]:
		return position + (board.SIZE - 1), false
	}
	return -42, true
}

func ReturnNextPiece(position, direction int) (nextIndex int, edge bool) {
	nextIndex, edge = getNextIndexForDirection(position, direction)
	if edge {
		return -42, true
	}
	return GameRound.Goban.Tab[nextIndex], false
}

func checkWinningConditions(lastPosition int, sequences [][]int) bool {
	if GameRound.CurrentPlayer.CapturedPieces == 10 {
		return true
	}
	for _, v := range sequences {
		if len(v) >= 5 {
			return true
		}
	}
	return false
}

func updateWhoseTurn() {
	if GameRound.CurrentPlayer == GameRound.P1 {
		GameRound.CurrentPlayer = GameRound.P2
	} else {
		GameRound.CurrentPlayer = GameRound.P1
	}
}

func HandleMove(id int, position int) (code int, msg string) {
	fmt.Println("making move at...", position, "for Player...", GameRound.CurrentPlayer.Id)
	if GameRound.Winner != 0 {
		return 1, "Game is over"
	}
	if GameRound.CurrentPlayer.Id != id {
		return 1, "It is not your turn"
	}
	if !CheckValidMove(position) {
		return 1, "Move isn't valid"
	}
	if GameRound.CurrentPlayer.PiecesLeft == 0 {
		return 1, "You have no pieces left"
	}
	GameRound.Goban.Tab[position] = int(id)
	GameRound.CurrentPlayer.PiecesLeft--
	captureDirections := checkCapture(position)
	capturePairs(position, captureDirections)
	sequences := CompleteSequenceForPosition(position, id)
	win := checkWinningConditions(position, sequences)
	if win {
		GameRound.Winner = id
	}
	GameRound.Turn++
	updateWhoseTurn()
	SuggestMove()
	if GameRound.CurrentPlayer.Id == 1 {
		HandleMove(GameRound.CurrentPlayer.Id, GameRound.SuggestedPosition)
	}
	return 0, "Move done"
}
