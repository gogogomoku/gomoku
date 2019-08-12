package brain

import (
	"fmt"
	"time"

	"github.com/gogogomoku/gomoku/internal/board"
	tr "github.com/gogogomoku/gomoku/internal/tree"
)

type Move struct {
	Position int
	Value    int
}

func getPossibleMoves(node *tr.Node) []Move {
	poss := []Move{}
	for i := 0; i < (board.SIZE * board.SIZE); i++ {
		if node.Tab[i] == 0 {
			if CheckValidMove(i, node.Tab) {
				// Add move only if it is aligned in a range of 5 from other
				affectsTab := false
				lines := CheckNextN(i, node.Tab, 1)
				for _, line := range lines {
					for _, piece := range line {
						if piece != 0 {
							affectsTab = true
							break
						}
					}
				}
				if affectsTab {
					poss = append(poss, Move{Position: i, Value: 0})
				}
			}
		}
	}
	return poss
}

func addNewLayer(poss []Move, node *tr.Node, playerId int) {
	for i, m := range poss {
		newTab := append([]int{}, node.Tab...)
		newTab[m.Position] = playerId
		captureDirections := checkCapture(m.Position, &node.Tab, playerId)
		// Virtual Capturing
		capturePairs(m.Position, captureDirections, &newTab)
		new := tr.Node{
			Id:       i + node.Id,
			Value:    0,
			Tab:      newTab,
			Position: m.Position,
			Player:   playerId,
		}
		tr.AddChild(node, &new)
	}
}

func SuggestMove() {

	startTime := time.Now()
	//Create tree
	tree := tr.Node{Id: 1, Value: 0, Tab: Game.Goban.Tab, Player: Game.CurrentPlayer.Id}
	poss := getPossibleMoves(&tree)
	opponent := 1
	if Game.CurrentPlayer.Id == 1 {
		opponent = 2
	}
	// Players move
	addNewLayer(poss, &tree, Game.CurrentPlayer.Id)
	// opponents move
	for _, ch := range tree.Children {
		poss := getPossibleMoves(ch)
		addNewLayer(poss, ch, opponent)
	}
	// Players move
	for _, ch := range tree.Children {
		for _, ch2 := range ch.Children {
			poss := getPossibleMoves(ch2)
			addNewLayer(poss, ch2, Game.CurrentPlayer.Id)
		}
	}
	// // opponents move
	// for _, ch := range tree.Children {
	// 	for _, ch2 := range ch.Children {
	// 		for _, ch3 := range ch2.Children {
	// 			poss := getPossibleMoves(ch3)
	// 			addNewLayer(poss, ch3, opponent)
	// 		}
	// 	}
	// }
	// // Players move
	// for _, ch := range tree.Children {
	// 	for _, ch2 := range ch.Children {
	// 		for _, ch3 := range ch2.Children {
	// 			for _, ch4 := range ch3.Children {
	// 				poss := getPossibleMoves(ch4)
	// 				addNewLayer(poss, ch4, Game.CurrentPlayer.Id)
	// 			}
	// 		}
	// 	}
	// }

	// Launch algo
	LaunchMinimaxPruning(&tree, 3)

	Game.SuggestedPosition = tree.SelectedChild.Position
	duration := time.Since(startTime)
	fmt.Println("Time spent on suggestion:", duration)
}
