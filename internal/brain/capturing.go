package brain

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
