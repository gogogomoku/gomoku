package brain

import (
	"github.com/gogogomoku/gomoku/internal/board"
)

func checkHorizontalSequences(playerId int, tab *[]int) int {
	score := 0
	for l := 0; l < board.SIZE; l++ {
		line := (*tab)[l*board.SIZE : (l+1)*board.SIZE]
		score += checkSequence(line, playerId)
	}
	return score
}

func checkVerticalSequences(playerId int, tab *[]int) int {
	score := 0
	for c := 0; c < board.SIZE; c++ {
		line := []int{}
		for l := 0; l < board.SIZE; l++ {
			line = append(line, (*tab)[l*board.SIZE+c])
		}
		score += checkSequence(line, playerId)
	}
	return score
}

func checkDiagonalNWSESequences(playerId int, tab *[]int) int {
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

func checkDiagonalNESWSequences(playerId int, tab *[]int) int {
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

func getHeuristicValue(position int, playerId int, tab *[]int) int {
	boardScore := 0
	opponent := 1
	if playerId == 1 {
		opponent = 2
	}
	// Check sequences for player
	boardScore += checkHorizontalSequences(playerId, tab)
	boardScore += checkVerticalSequences(playerId, tab)
	boardScore += checkDiagonalNWSESequences(playerId, tab)
	boardScore += checkDiagonalNESWSequences(playerId, tab)
	// Check sequences for opponent
	boardScore -= checkHorizontalSequences(opponent, tab)
	boardScore -= checkVerticalSequences(opponent, tab)
	boardScore -= checkDiagonalNWSESequences(opponent, tab)
	boardScore -= checkDiagonalNESWSequences(opponent, tab)
	// board.PrintBoard(*tab, board.SIZE)
	// fmt.Println(boardScore)
	return boardScore
}

func checkSequence(line []int, playerId int) int {
	opponent := 1
	if playerId == 1 {
		opponent = 2
	}
	i := 0
	counter := 0
	score := 0
	blocked := 0
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
				tmpScore += 2000
			}
			if blocked == 1 {
				tmpScore = int(float64(tmpScore) * 0.7)
			} else if blocked == 2 {
				tmpScore = 0
			}
			if counter == 5 {
				tmpScore = 50000
			}
			counter = 0
			blocked = 0
		}
		score += tmpScore
		i++
	}
	// 	fmt.Println(line)
	// 	fmt.Println("Score;", score)
	return score
}
