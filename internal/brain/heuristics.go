package brain

import (
	// "fmt"
	"sync"

	"github.com/gogogomoku/gomoku/internal/board"
)

const (
	SEQ2_FREE_SCORE     = 10
	SEQ3_BLOCKED1_SCORE = 100
	SEQ4_BLOCKED1_SCORE = 5000
	SEQ4_FREE_SCORE     = 40000
	F3_SCORE            = 1000
	WIN_SCORE           = 100000
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
	for d := 1; d < board.SIZE+board.SIZE/2; d++ {
		line := [board.SIZE * 2]int{}
		// First element of last row to start diagonals
		x := board.SIZE * (board.SIZE - d)
		// When transpassing board limit, start new diagonal
		lineLen := 0
		for x < board.SIZE*board.SIZE {
			if x >= 0 && x < (board.SIZE*board.SIZE)-(board.SIZE*(d-board.SIZE)) {
				line[lineLen] = (*tab)[x]
				lineLen++
			}
			x += board.SIZE + 1
		}
		line[lineLen] = -1
		line2 := line[:lineLen]
		if len(line2) > 0 {
			score += checkSequence(line2, playerId)
		}
	}
	return score
}

func checkDiagonalNESWSequences(playerId int, tab *[board.TOT_SIZE]int) int {
	score := 0
	for d := 1; d < board.SIZE*2; d++ {
		line := [board.SIZE * 2]int{}
		// Last element of first row to start diagonals
		x := (d*board.SIZE - 1) - (board.SIZE-1)*board.SIZE
		// When transpassing board limit, start new diagonal
		lineLen := 0
		for x < board.SIZE*board.SIZE {
			if x >= 0 && x <= (d-1)*board.SIZE {
				line[lineLen] = (*tab)[x]
				lineLen++
			}
			x += board.SIZE - 1
		}
		line[lineLen] = -1
		line2 := line[:lineLen]
		if len(line2) > 0 {
			score += checkSequence(line2, playerId)
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
	playerScore := boardScorePlayerDINWSE + boardScorePlayerDINESW + boardScorePlayerHV
	opponentScore := boardScoreOpponentDINWSE + boardScoreOpponentDINESW + boardScoreOpponentHV
	opponentScore = int(float64(opponentScore) * 1.4)
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
			if counter >= 5 {
				return WIN_SCORE
			}
			if i == 0 || line[i-1] == opponent {
				blocked++
			}
			if i == len(line)-1 || line[i+1] == opponent {
				blocked++
			}
		} else {
			switch counter {
			case 2:
				tmpScore += SEQ2_FREE_SCORE
			case 3:
				if blocked == 1 {
					tmpScore += SEQ3_BLOCKED1_SCORE
				}
			case 4:
				if blocked == 0 {
					tmpScore += SEQ4_FREE_SCORE
				} else {
					tmpScore += SEQ4_BLOCKED1_SCORE
				}
			}
			if blocked == 2 {
				tmpScore = 0
			}
			counter = 0
			blocked = 0
		}
		score += tmpScore
		i++
	}
	f3 := len(CheckSequenceForF3(line, playerId))
	score += F3_SCORE * f3
	return score
}
