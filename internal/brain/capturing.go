package brain

func checkCapture(position int, tab *[19 * 19]int, playerId int) []int {
	captureDirections := []int{}
	for direction := 0; direction < 8; direction++ {
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

func capturePairs(position int, captureDirections []int, tab *[19 * 19]int) {
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
