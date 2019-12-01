package brain

import (
	// "fmt"

	"sync"

	"github.com/gogogomoku/gomoku/internal/board"
)

const (
	SEQ2_BLOCKED1_SCORE = 2
	SEQ2_FREE_SCORE     = 4
	SEQ3_BLOCKED1_SCORE = 10
	SEQ4_BLOCKED1_SCORE = 90
	SEQ4_FREE_SCORE     = 2000
	SEQ4_BROKEN         = 1000
	F3_SCORE            = 100
	WIN_SCORE           = 10000
)

//Add tests
//Add tests
//Add tests
var CAPTURED_SCORE = [5]int16{0, 1, 2, 100, 1000}

func convertArrayToSlice(line [board.SIZE]int16) []int16 {
	new := make([]int16, board.SIZE)
	for i := 0; i < board.SIZE; i++ {
		new[i] = line[i]
	}
	return new
}

func checkHorizontalSequences(playerId int16, tab *[board.TOT_SIZE]int16) int16 {
	score := int16(0)
	for l := 0; l < board.SIZE; l++ {
		line := [board.SIZE]int16{}
		for c := 0; c < board.SIZE; c++ {
			line[c] = (*tab)[l*board.SIZE+c]
		}
		lineSlice := convertArrayToSlice(line)
		score += checkSequence(lineSlice, playerId)
	}
	return score
}

func checkVerticalSequences(playerId int16, tab *[board.TOT_SIZE]int16) int16 {
	score := int16(0)
	for c := 0; c < board.SIZE; c++ {
		line := [board.SIZE]int16{}
		for l := 0; l < board.SIZE; l++ {
			line[l] = (*tab)[l*board.SIZE+c]
		}
		lineSlice := convertArrayToSlice(line)
		score += checkSequence(lineSlice, playerId)
	}
	return score
}

func checkDiagonalNWSESequences(playerId int16, tab *[board.TOT_SIZE]int16) int16 {
	score := int16(0)
	for d := 1; d < board.SIZE+board.SIZE/2; d++ {
		line := [board.SIZE * 2]int16{}
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

func checkDiagonalNESWSequences(playerId int16, tab *[board.TOT_SIZE]int16) int16 {
	score := int16(0)
	for d := 1; d < board.SIZE*2; d++ {
		line := [board.SIZE * 2]int16{}
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

func getHeuristicValue(playerId int16, tab *[board.TOT_SIZE]int16, captured *[3]int16) int16 {
	boardScorePlayerHV := int16(0)
	boardScorePlayerDINWSE := int16(0)
	boardScorePlayerDINESW := int16(0)
	boardScoreOpponentHV := int16(0)
	boardScoreOpponentDINWSE := int16(0)
	boardScoreOpponentDINESW := int16(0)
	opponent := int16(1)
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
	if captured[playerId] > 0 && captured[playerId] <= 8 {
		playerScore += CAPTURED_SCORE[(captured[playerId]/2)-1]
	} else if captured[playerId] > 8 {
		playerScore += WIN_SCORE
	}
	opponentScore := boardScoreOpponentDINWSE + boardScoreOpponentDINESW + boardScoreOpponentHV
	if captured[opponent] > 0 && captured[opponent] <= 8 {
		opponentScore += CAPTURED_SCORE[(captured[opponent]/2)-1]
	} else if captured[opponent] > 8 {
		playerScore += WIN_SCORE
	}
	if playerScore >= WIN_SCORE {
		playerScore = WIN_SCORE
	} else if opponentScore >= WIN_SCORE {
		opponentScore = WIN_SCORE
		playerScore = 0
	} else {
		opponentScore += opponentScore / 5
	}
	return playerScore - opponentScore
}

func checkLineHasId(line *[]int16, playerId int16) bool {
	for _, v := range *line {
		if v == playerId {
			return true
		}
	}
	return false
}

func getSequenceScore(counter int16, blocked int16, line *[]int16, i int16) int16 {
	tmpScore := int16(0)

	switch counter {
	case 2:
		if blocked == 0 {
			tmpScore += SEQ2_FREE_SCORE
		} else if blocked == 1 {
			tmpScore += SEQ2_BLOCKED1_SCORE
		}
		// Check 2 sequence of 2 separated by empty space

		if i < int16(len(*line)-2) && (*line)[i] == 0 {
			player := (*line)[i-1]
			if (*line)[i+1] == (*line)[i+2] && (*line)[i+1] == player {
				if i < int16(len(*line)-3) && (*line)[i+3] == 0 {
					tmpScore += SEQ4_FREE_SCORE
				} else {
					tmpScore += SEQ4_BROKEN
				}
			}
		}
	case 3:
		if blocked == 1 {
			tmpScore += SEQ3_BLOCKED1_SCORE
		}
		if blocked != 2 {
			// Check 2 sequence of 1(or more) and 3 separated by empty space
			// AFTER SEQ_3
			if i < int16(len(*line)-2) && (*line)[i] == 0 {
				if (*line)[i+1] == (*line)[i-1] {
					tmpScore += SEQ4_BROKEN
				}
			}
			// Check 2 sequence of 1(or more) and 3 separated by empty space
			// BEFORE SEQ_3
			if i > 4 && (*line)[i-4] == 0 {
				if (*line)[i-5] == (*line)[i-3] {
					tmpScore += SEQ4_BROKEN
				}
			}
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
	return tmpScore
}

// Need to not check all sequences
func checkSequence(line []int16, playerId int16) int16 {
	hasPlayer := checkLineHasId(&line, playerId)
	if !hasPlayer {
		return 0
	}

	i := int16(0)
	counter := int16(0)
	score := int16(0)
	blocked := int16(0)
	opponent := int16(1)
	if playerId == 1 {
		opponent = 2
	}
	for i < int16(len(line)) {
		tmpScore := int16(0)
		if line[i] == playerId {
			counter++
			if counter >= 5 {
				return WIN_SCORE
			}
			if i == 0 || line[i-1] == opponent {
				blocked++
			}
			if i == int16(len(line)-1) || line[i+1] == opponent {
				blocked++
			}
		} else {
			tmpScore += getSequenceScore(counter, blocked, &line, i)
			counter = 0
			blocked = 0
		}
		score += tmpScore
		i++
	}
	score += int16(len(CheckSequenceForF3(line, playerId))) * F3_SCORE
	return score
}
