package brain

import (
	// "fmt"

	// "sync"

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
	WIN_SCORE           = 19000
)

type HeuristicScore struct {
	Free5 int16
	Free4 int16
	Half4 int16
	Free3 int16
	Half3 int16
	Free2 int16
	Half2 int16
}

var CAPTURED_SCORE = [5]int16{0, 1, 2, 100, 1000}

func convertArrayToSlice(line [board.SIZE]int16) []int16 {
	new := make([]int16, board.SIZE)
	for i := 0; i < board.SIZE; i++ {
		new[i] = line[i]
	}
	return new
}

func checkHorizontalSequences(playerId int16, tab *[board.TOT_SIZE]int16) {
	for l := 0; l < board.SIZE; l++ {
		line := [board.SIZE]int16{}
		for c := 0; c < board.SIZE; c++ {
			line[c] = (*tab)[l*board.SIZE+c]
		}
		lineSlice := convertArrayToSlice(line)
		checkSequence(lineSlice, playerId)
	}
}

func checkVerticalSequences(playerId int16, tab *[board.TOT_SIZE]int16) {
	for c := 0; c < board.SIZE; c++ {
		line := [board.SIZE]int16{}
		for l := 0; l < board.SIZE; l++ {
			line[l] = (*tab)[l*board.SIZE+c]
		}
		lineSlice := convertArrayToSlice(line)
		checkSequence(lineSlice, playerId)
	}
}

func checkDiagonalNWSESequences(playerId int16, tab *[board.TOT_SIZE]int16) {
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
			checkSequence(line2, playerId)
		}
	}
}

func checkDiagonalNESWSequences(playerId int16, tab *[board.TOT_SIZE]int16) {
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
			checkSequence(line2, playerId)
		}
	}
}

func CheckSequenceForF5(line []int16, playerId int16) int16 {
	counter := 0
	for _, t := range line {
		if t == playerId {
			counter++
			if counter == 5 {
				return 1
			}
		} else {
			counter = 0
		}
	}
	return 0
}

func CheckSequenceForF4(line []int16, playerId int16) (int16, int16) {
	counter := 0
	blocked := 0
	free4 := int16(0)
	half4 := int16(0)
	for i, t := range line {
		if t == playerId {
			if i == 0 || line[i-1] == getOpponent(playerId) || i == len(line)-1 || line[i+1] == getOpponent(playerId) {
				blocked++
			}
			counter++
			if counter == 4 {
				if blocked == 0 {
					free4++
				} else if blocked == 1 {
					half4++
				}
				counter = 0
				blocked = 0
			}
		} else {
			counter = 0
			blocked = 0
		}
	}
	return free4, half4
}

func CheckSequenceForHalf3(line []int16, playerId int16) int16 {
	counter := 0
	blocked := 0
	half3 := int16(0)
	for i, t := range line {
		if t == playerId {
			if i == 0 || line[i-1] == getOpponent(playerId) || i == len(line)-1 || line[i+1] == getOpponent(playerId) {
				blocked++
			}
			counter++
			if counter == 3 && (i == len(line)-1 || line[i+1] != playerId) {
				if blocked == 1 {
					half3++
				}
				counter = 0
				blocked = 0
			}
		} else {
			counter = 0
			blocked = 0
		}
	}
	return half3
}

func CheckSequenceForF2(line []int16, playerId int16) (int16, int16) {
	counter := 0
	blocked := 0
	free2 := int16(0)
	half2 := int16(0)
	for i, t := range line {
		if t == playerId {
			if i == 0 || line[i-1] == getOpponent(playerId) || i == len(line)-1 || line[i+1] == getOpponent(playerId) {
				blocked++
			}
			counter++
			if counter == 2 && (i == len(line)-1 || line[i+1] != playerId) {
				if blocked == 0 {
					free2++
				} else if blocked == 1 {
					half2++
				}
				counter = 0
				blocked = 0
			}
		} else {
			counter = 0
			blocked = 0
		}
	}
	return free2, half2
}

func checkLineHasId(line *[]int16, playerId int16) bool {
	for _, v := range *line {
		if v == playerId {
			return true
		}
	}
	return false
}

func checkSequence(line []int16, playerId int16) {

	if checkLineHasId(&line, playerId) {
		scorePlayer.Free5 += int16(CheckSequenceForF5(line, playerId))
		free4Player, half4Player := CheckSequenceForF4(line, playerId)
		scorePlayer.Free4 += free4Player
		scorePlayer.Half4 += half4Player
		scorePlayer.Free3 += int16(len(CheckSequenceForF3(line, playerId)))
		scorePlayer.Half3 += CheckSequenceForHalf3(line, playerId)
		free2Player, half2Player := CheckSequenceForF2(line, playerId)
		scorePlayer.Free2 += free2Player
		scorePlayer.Half2 += half2Player
	}

	if checkLineHasId(&line, getOpponent(playerId)) {
		scoreOpponent.Free5 += int16(CheckSequenceForF5(line, getOpponent(playerId)))
		free4Opponent, half4Opponent := CheckSequenceForF4(line, getOpponent(playerId))
		scoreOpponent.Free4 += free4Opponent
		scoreOpponent.Half4 += half4Opponent
		scoreOpponent.Free3 += int16(len(CheckSequenceForF3(line, getOpponent(playerId))))
		scoreOpponent.Half3 += CheckSequenceForHalf3(line, getOpponent(playerId))
		free2Opponent, half2Opponent := CheckSequenceForF2(line, getOpponent(playerId))
		scoreOpponent.Free2 += free2Opponent
		scoreOpponent.Half2 += half2Opponent
	}
}

var scorePlayer HeuristicScore
var scoreOpponent HeuristicScore

func getHeuristicValue(playerId int16, tab *[board.TOT_SIZE]int16, captured *[3]int16) int16 {
	scorePlayer = HeuristicScore{}
	scoreOpponent = HeuristicScore{}
	checkHorizontalSequences(playerId, tab)
	checkVerticalSequences(playerId, tab)
	checkDiagonalNESWSequences(playerId, tab)
	checkDiagonalNWSESequences(playerId, tab)

	// fmt.Println("Score Player", scorePlayer)
	// fmt.Println("Score Opponent", scoreOpponent)

	score := int16(0)
	if scoreOpponent.Free5 > 0 {
		score -= 28000
	} else if scorePlayer.Free5 > 0 {
		score += 19500
	} else if scoreOpponent.Free4 > 0 && (scorePlayer.Half4+scorePlayer.Free4 == 0) {
		score -= 10000
	} else if scorePlayer.Free4 > 0 && (scoreOpponent.Half4+scoreOpponent.Free4 == 0) {
		score += 9500
	} else if scoreOpponent.Free3 > 0 && scoreOpponent.Half4 > 0 {
		score -= 5000
	} else if scorePlayer.Free3 > 0 && scorePlayer.Half4 > 0 {
		score += 4500
	} else if scoreOpponent.Half4 > 1 {
		score -= 2000
	} else if scorePlayer.Half4 > 1 {
		score += 1500
	}
	// score += scorePlayer.Free4 * 500
	// score += scorePlayer.Half4 * 200
	score += scorePlayer.Free3 * 100
	score += scorePlayer.Half3 * 20
	score += scorePlayer.Free2 * 5
	score += scorePlayer.Half2 * 2

	// score -= scoreOpponent.Free4 * 1000
	// score -= scoreOpponent.Half4 * 400
	score -= scoreOpponent.Free3 * 200
	score -= scoreOpponent.Half3 * 40
	score -= scoreOpponent.Free2 * 10
	score -= scoreOpponent.Half2 * 4
	return score
}
