package brain

import (
	"fmt"

	"github.com/gogogomoku/gomoku/internal/board"
	"github.com/gogogomoku/gomoku/internal/player"
)

func StartRound(AiStatus1 int16, AiStatus2 int16) {
	Game.Status = Running
	player.ResetPlayers(Game.P1, Game.P2, MAXPIECES, AiStatus1, AiStatus2)
	Game.CurrentPlayer = Game.P1
	center := int16((board.TOT_SIZE) / 2)
	if board.SIZE%2 == 0 {
		center += board.SIZE / 2
	}
	SuggestMove(Game.CurrentPlayer.Id)
}

func CheckValidMove(position int16, tab [board.TOT_SIZE]int16, playerId int16) bool {
	if position >= 0 && position <= (board.TOT_SIZE)-1 {
		if tab[position] == 0 {
			return !Check2F3s(playerId, position, &tab)
		}
	}
	return false
}

func getNextIndexForDirection(position int16, direction int16) (nextIndex int16, edge bool) {
	directions := [4]bool{true, true, true, true}
	// First row
	if position < board.SIZE {
		directions[N] = false
	}
	// Last row
	if position >= (board.SIZE * (board.SIZE - 1)) {
		directions[S] = false
	}
	// East column
	if position%board.SIZE == (board.SIZE - 1) {
		directions[E] = false
	}
	// West column
	if position%board.SIZE == 0 {
		directions[W] = false
	}
	switch {
	case direction == N && directions[N]:
		return position - board.SIZE, false
	case direction == S && directions[S]:
		return position + board.SIZE, false
	case direction == E && directions[E]:
		return position + 1, false
	case direction == W && directions[W]:
		return position - 1, false
	case direction == NE && directions[N] && directions[E]:
		return position - (board.SIZE - 1), false
	case direction == NW && directions[N] && directions[W]:
		return position - (board.SIZE + 1), false
	case direction == SE && directions[S] && directions[E]:
		return position + (board.SIZE + 1), false
	case direction == SW && directions[S] && directions[W]:
		return position + (board.SIZE - 1), false
	}
	return -42, true
}

func ReturnNextPiece(position int16, direction int16, tab *[board.TOT_SIZE]int16) (nextIndex int16, edge bool) {
	nextIndex, edge = getNextIndexForDirection(position, direction)
	if edge {
		return -42, true
	}
	return (*tab)[nextIndex], false
}

func checkOpponentCancelMyWin(lastPosition int16, tab *[board.TOT_SIZE]int16, opponent *player.Player, currentPlayer *player.Player) bool {
	opponentPossibleMoves := getPossibleMoves(tab, opponent.Id)
	for _, pos := range opponentPossibleMoves {
		captureDirections := checkCapture(pos, &(Game.Goban.Tab), opponent.Id)
		nCaptures := int16(2 * len(captureDirections))
		if opponent.CapturedPieces+nCaptures >= 10 {
			return true
		} else {
			tabCopy := *tab
			capturePairs(pos, captureDirections, &tabCopy)
			sequences := CompleteSequenceForPosition(lastPosition, currentPlayer.Id, &tabCopy)
			foundWinSeq := false
			for _, v := range sequences {
				if len(v) >= 5 {
					foundWinSeq = true
				}
			}
			if !foundWinSeq {
				return true
			}
		}
	}
	return false
}

func checkWinBySeq(lastPosition int16, sequences [][]int16) bool {
	hasWinLengthSeq := false
	for _, v := range sequences {
		if len(v) >= 5 {
			hasWinLengthSeq = !checkOpponentCancelMyWin(lastPosition, &Game.Goban.Tab, Game.GetCurrentOpponent(), Game.CurrentPlayer)
			if hasWinLengthSeq {
				return true
			}
		}
	}
	return false
}

func updateWhoseTurn() {
	if Game.CurrentPlayer == Game.P1 {
		Game.CurrentPlayer = Game.P2
	} else {
		Game.CurrentPlayer = Game.P1
	}
}

func HandleMove(playerId int16, position int16) (code int16, msg string) {

	fmt.Println("making move at...", position,
		"for Player...", Game.CurrentPlayer.Id)
	if Game.Winner != 0 {
		return 1, "Game is over"
	}
	if Game.CurrentPlayer.Id != playerId {
		return 1, "It is not your turn"
	}
	if !CheckValidMove(position, Game.Goban.Tab, playerId) {
		return 1, "Move isn't valid"
	}
	if Game.CurrentPlayer.PiecesLeft == 0 {
		return 1, "You have no pieces left"
	}
	Game.Goban.Tab[position] = int16(playerId)
	Game.CurrentPlayer.PiecesLeft--
	captureDirections := checkCapture(position, &Game.Goban.Tab, Game.CurrentPlayer.Id)
	capturePairs(position, captureDirections, &Game.Goban.Tab)
	sequences := CompleteSequenceForPosition(position, playerId, &Game.Goban.Tab)
	winByCaptures := Game.CurrentPlayer.CapturedPieces >= 10
	winBySeq := false
	if !winByCaptures {
		winBySeq = checkWinBySeq(position, sequences)
	}
	if winBySeq || winByCaptures {
		Game.SuggestedPosition = board.TOT_SIZE + 1
		Game.Winner = playerId
	} else {
		Game.Turn++
		updateWhoseTurn()
		SuggestMove(Game.CurrentPlayer.Id)
	}
	return 0, "Move done"
}
