package brain

import (
	"sync"

	"github.com/gogogomoku/gomoku/internal/board"
)

func convertArrayToSlice(line [board.SIZE]int) []int {
	new := make([]int, board.SIZE)
	for i := 0; i < board.SIZE; i++ {
		new[i] = line[i]
	}
	return new
}

func checkHorizontalSequences(playerId int, tab *[board.TOT_SIZE]int) int {
	score := 0
	for l := 0; l < board.SIZE; l++ {
		line := [board.SIZE]int{}
		for c := 0; c < board.SIZE; c++ {
			line[c] = (*tab)[l*board.SIZE+c]
		}
		lineSlice := convertArrayToSlice(line)
		score += checkSequence(lineSlice, playerId)
	}
	return score
}

func checkVerticalSequences(playerId int, tab *[board.TOT_SIZE]int) int {
	score := 0
	for c := 0; c < board.SIZE; c++ {
		line := [board.SIZE]int{}
		for l := 0; l < board.SIZE; l++ {
			line[l] = (*tab)[l*board.SIZE+c]
		}
		lineSlice := convertArrayToSlice(line)
		score += checkSequence(lineSlice, playerId)
	}
	return score
}

func checkDiagonalNWSESequences(playerId int, tab *[board.TOT_SIZE]int) int {
	score := 0
	for d := 1; d < board.SIZE*2; d++ {
		line := []int{}
		// First element of last row to start diagonals
		x := board.SIZE * (board.SIZE - d)
		// When transpassing board limit, start new diagonal
		for x < board.SIZE*board.SIZE {
			if x >= 0 && x < (board.SIZE*board.SIZE)-(board.SIZE*(d-board.SIZE)) {
				line = append(line, (*tab)[x])
			}
			x += board.SIZE + 1
		}
		if len(line) > 0 {
			score += checkSequence(line, playerId)
		}
	}
	return score
}

func checkDiagonalNESWSequences(playerId int, tab *[board.TOT_SIZE]int) int {
	score := 0
	for d := 1; d < board.SIZE*2; d++ {
		line := []int{}
		// Last element of first row to start diagonals
		x := (d*board.SIZE - 1) - (board.SIZE-1)*board.SIZE
		// When transpassing board limit, start new diagonal
		for x < board.SIZE*board.SIZE {
			if x >= 0 && x <= (d-1)*board.SIZE {
				line = append(line, (*tab)[x])
			}
			x += board.SIZE - 1
		}
		if len(line) > 0 {
			score += checkSequence(line, playerId)
		}
	}
	return score
}

func getHeuristicValue(position int, playerId int, tab *[board.TOT_SIZE]int) int {
	boardScorePlayerHV := 0
	boardScorePlayerDINWSE := 0
	boardScorePlayerDINESW := 0
	boardScoreOpponentHV := 0
	boardScoreOpponentDINWSE := 0
	boardScoreOpponentDINESW := 0
	opponent := 1
	if playerId == 1 {
		opponent = 2
	}
	var waitgroup sync.WaitGroup
	// Check sequences for player
	waitgroup.Add(6)
	go func() {
		defer waitgroup.Done()
		boardScorePlayerHV += checkHorizontalSequences(playerId, tab)
		boardScorePlayerHV += checkVerticalSequences(playerId, tab)
	}()
	go func() {
		defer waitgroup.Done()
		boardScorePlayerDINWSE += checkDiagonalNWSESequences(playerId, tab)
	}()
	go func() {
		defer waitgroup.Done()
		boardScorePlayerDINESW += checkDiagonalNESWSequences(playerId, tab)
	}()

	// Check sequences for opponent
	go func() {
		defer waitgroup.Done()
		boardScoreOpponentHV += checkHorizontalSequences(opponent, tab)
		boardScoreOpponentHV += checkVerticalSequences(opponent, tab)
	}()
	go func() {
		defer waitgroup.Done()
		boardScoreOpponentDINWSE += checkDiagonalNWSESequences(opponent, tab)
	}()
	go func() {
		defer waitgroup.Done()
		boardScoreOpponentDINESW += checkDiagonalNESWSequences(opponent, tab)
	}()
	waitgroup.Wait()
	// board.PrintBoard(*tab, board.SIZE)
	// fmt.Println(boardScore)
	playerScore := boardScorePlayerDINWSE + boardScorePlayerDINESW + boardScorePlayerHV
	opponentScore := boardScoreOpponentDINWSE + boardScoreOpponentDINESW + boardScoreOpponentHV
	if playerScore >= 100000 {
		playerScore = 100000
		opponentScore = -100000
	}
	if opponentScore >= 100000 {
		opponentScore = 100000
		playerScore = -100000
	}
	return playerScore - opponentScore
}

func checkLineHasId(line *[]int, playerId int) bool {
	for _, v := range *line {
		if v == playerId {
			return true
		}
	}
	return false
}

func checkSequence(line []int, playerId int) int {
	hasPlayer := checkLineHasId(&line, playerId)
	if !hasPlayer {
		return 0
	}
	i := 0
	counter := 0
	score := 0
	blocked := 0
	opponent := 1
	if playerId == 1 {
		opponent = 2
	}
	for i < len(line) {
		tmpScore := 0
		if line[i] == playerId {
			counter++
			if i == 0 || line[i-1] == opponent {
				blocked++
			}
			if i == len(line)-1 || line[i+1] == opponent {
				blocked++
			}
		} else {

			switch counter {
			case 2:
				tmpScore += 10
			case 3:
				tmpScore += 100
			case 4:
				if blocked == 0 {
					tmpScore += 5000
				} else {
					tmpScore += 100
				}

			}
			if blocked == 2 {
				tmpScore = 0
			}
			if counter == 5 {
				tmpScore = 100000
			}
			counter = 0
			blocked = 0
		}
		score += tmpScore
		i++
	}
	score += 100 * len(CheckSequenceForF3(line, playerId))
	// 	fmt.Println(line)
	// 	fmt.Println("Score;", score)
	return score
}
