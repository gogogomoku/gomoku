package brain

import (
	"github.com/gogogomoku/gomoku/internal/board"
)

func checkCapture(position int16, tab *[board.TOT_SIZE]int16, playerId int16) []int16 {
	captureDirections := []int16{}
	var direction int16
	for direction = 0; direction < 8; direction++ {
		counter := 0
		tmpPosition := position
		for j := 0; j < 2; j++ {
			nextIndex, edge := getNextIndexForDirection(tmpPosition, direction)
			nextIndexValue, edge := ReturnNextPiece(tmpPosition, direction, tab)
			if edge || nextIndexValue == playerId || nextIndexValue == 0 {
				break
			}
			counter++
			tmpPosition = nextIndex
		}
		if counter == 2 {
			nextIndexValue, _ := ReturnNextPiece(tmpPosition, direction, tab)
			if nextIndexValue == playerId {
				captureDirections = append(captureDirections, direction)
			}
		}
	}
	return captureDirections
}

func capturePairs(position int16, captureDirections []int16, tab *[board.TOT_SIZE]int16) {
	for _, captureDirection := range captureDirections {
		tmpPosition := position
		for j := 0; j < 2; j++ {
			nextIndex, _ := getNextIndexForDirection(tmpPosition, captureDirection)
			tmpPosition = nextIndex
			(*tab)[nextIndex] = 0
		}
		if tab == &Game.Goban.Tab {
			Game.CurrentPlayer.CapturedPieces += 2
		}
	}
}
