package brain

import (
	"github.com/gogogomoku/gomoku/internal/board"
	"github.com/gogogomoku/gomoku/internal/player"
	"github.com/gogogomoku/gomoku/internal/suggestor"
)

func StartRound() {
	GameRound.Status = Running
	player.ResetPlayers(GameRound.P1, GameRound.P2, MAXPIECES)
	GameRound.CurrentPlayer = GameRound.P1
	getSuggestion()
}

func getSuggestion() {
	for {
		GameRound.SuggestedPosition = suggestor.SuggestMove()
		if checkValidMove(GameRound.SuggestedPosition) {
			break
		}
	}
}

func checkValidMove(position int) bool {
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

func checkCapture(position int) []int {
	captureDirections := []int{}
	for direction := 0; direction < 8; direction++ {
		counter := 0
		tmpPosition := position
		for j := 0; j < 2; j++ {
			nextIndex, edge := getNextIndexForDirection(tmpPosition, direction)
			nextIndexValue, edge := ReturnNextPiece(tmpPosition, direction)
			if edge || nextIndexValue == GameRound.CurrentPlayer.Id || nextIndexValue == 0 {
				break
			}
			counter++
			tmpPosition = nextIndex
		}
		if counter == 2 {
			nextIndexValue, _ := ReturnNextPiece(tmpPosition, direction)
			if nextIndexValue == GameRound.CurrentPlayer.Id {
				captureDirections = append(captureDirections, direction)
			}
		}
	}
	return captureDirections
}

// For each direction in usual order, return how many contiguous pieces for playerId are present
func checkSequence(position int, playerId int) []int {
	sequenceLengths := []int{}
	for direction := 0; direction < 8; direction++ {
		counter := 0
		tmpPosition := position
		for {
			nextIndex, edge := getNextIndexForDirection(tmpPosition, direction)
			nextIndexValue, edge := ReturnNextPiece(tmpPosition, direction)
			if edge || nextIndexValue != playerId || nextIndexValue == 0 {
				break
			}
			counter++
			tmpPosition = nextIndex
		}
		sequenceLengths = append(sequenceLengths, counter)
	}
	return sequenceLengths
}

func sequenceOpposingDirections(position int, playerId int, dir []int, increase int) []int {
	partialSequences := checkSequence(position, playerId)
	if partialSequences[dir[0]] != 0 || partialSequences[dir[1]] != 0 {
		p := []int{position}
		// Add elements in dir[0]
		current := position
		for i := 0; i < partialSequences[dir[0]]; i++ {
			current -= increase
			p = append(p, current)
		}
		// Add elements in dir[1]
		current = position
		for i := 0; i < partialSequences[dir[1]]; i++ {
			current += increase
			p = append(p, current)
		}
		return p
	}
	return nil
}

func completeSequenceForPosition(position int, playerId int) [][]int {
	sequences := [][]int{}
	sequenceDirections := []struct {
		OpposingDirections []int
		IncreaseValue      int
	}{
		{
			OpposingDirections: []int{N, S},
			IncreaseValue:      board.SIZE,
		}, {
			OpposingDirections: []int{W, E},
			IncreaseValue:      1,
		}, {
			OpposingDirections: []int{NW, SE},
			IncreaseValue:      board.SIZE + 1,
		}, {
			OpposingDirections: []int{NE, SW},
			IncreaseValue:      board.SIZE - 1,
		},
	}
	for _, d := range sequenceDirections {
		seq := sequenceOpposingDirections(position, playerId, d.OpposingDirections, d.IncreaseValue)
		if seq != nil {
			sequences = append(sequences, seq)
		}
	}
	return sequences
}

func capturePairs(position int, captureDirections []int) {
	for _, captureDirection := range captureDirections {
		tmpPosition := position
		for j := 0; j < 2; j++ {
			nextIndex, _ := getNextIndexForDirection(tmpPosition, captureDirection)
			tmpPosition = nextIndex
			GameRound.Goban.Tab[nextIndex] = 0
		}
		GameRound.CurrentPlayer.CapturedPieces += 2
	}
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
	if GameRound.Winner != 0 {
		return 1, "Game is over"
	}
	if GameRound.CurrentPlayer.Id != id {
		return 1, "It is not your turn"
	}
	if !checkValidMove(position) {
		return 1, "Move isn't valid"
	}
	if GameRound.CurrentPlayer.PiecesLeft == 0 {
		return 1, "You have no pieces left"
	}
	GameRound.Goban.Tab[position] = int(id)
	GameRound.CurrentPlayer.PiecesLeft--
	captureDirections := checkCapture(position)
	capturePairs(position, captureDirections)
	sequences := completeSequenceForPosition(position, id)
	win := checkWinningConditions(position, sequences)
	if win {
		GameRound.Winner = id
	}
	GameRound.Turn++
	updateWhoseTurn()
	getSuggestion()
	return 0, "Move done"
}
