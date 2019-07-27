package brain

import (
	"math/rand"

	"github.com/gogogomoku/gomoku/internal/board"
)

func SuggestMove() {
	// fmt.Println("Launch suggestion")
	// bestSuggestions := getBestSuggestions()
	// fmt.Println("Got suggestion")
	// fmt.Println(bestSuggestions)
	// ran := int(rand.Intn(len(bestSuggestions)))
	// fmt.Println(ran, bestSuggestions[ran])
	// GameRound.SuggestedPosition = bestSuggestions[ran]
	// fmt.Println("Ended suggestion")
	// possible := []int{}
	// for i := 0; i < board.SIZE*board.SIZE; i++ {
	// 	if GameRound.Goban.Tab[i] == 0 {
	// 		if CheckValidMove(i) {
	// 			seq := CheckSequence(i, GameRound.CurrentPlayer.Id)
	// 			// fmt.Println("----------->", seq)
	// 			possible = append(possible, i)
	// 		}
	// 	}
	// }
	// fmt.Println(possible)

	// for {
	// 	GameRound.SuggestedPosition = int(rand.Intn(board.SIZE * board.SIZE))
	// 	if CheckValidMove(GameRound.SuggestedPosition) {
	// 		break
	// 	}
	// }

	possible := []int{}
	best := []int{}
	for i := 0; i < board.SIZE*board.SIZE; i++ {
		if GameRound.Goban.Tab[i] == 0 {
			if CheckValidMove(i) {
				seq := CompleteSequenceForPosition(i, GameRound.CurrentPlayer.Id)
				// fmt.Println(i)
				// fmt.Println(seq)
				if len(seq) > 0 {
					best = append(best, i)
				}
				possible = append(possible, i)
			}
		}
	}
	ran := 0
	if len(best) > 0 {
		if len(best)-1 < 2 {
			ran = 0
		} else {
			ran = int(rand.Intn(len(best) - 1))
		}
		GameRound.SuggestedPosition = best[ran]
	} else {
		if len(best)-1 < 2 {
			ran = 0
		} else {
			ran = int(rand.Intn(len(best) - 1))
		}
		ran := int(rand.Intn(len(possible) - 1))
		GameRound.SuggestedPosition = possible[ran]
		// brain.HandleMove(brain.GameRound.CurrentPlayer.Id, possible[ran])
	}
	// board.PrintBoard(brain.GameRound.Goban)

}

// func CheckSequences(position int, playerId int) []int {
// 	seq := []int{}
// 	//N
// 	tmp := position
// 	ct := 0
// 	for {
// 		tmp -= board.SIZE
// 		if (tmp < 0) {
//
// 		}
// 	}
// 	return seq
// }
