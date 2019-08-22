package brain

import (
	"fmt"
	"sort"
	"time"

	"github.com/gogogomoku/gomoku/internal/board"
	tr "github.com/gogogomoku/gomoku/internal/tree"
)

var tree tr.Node

func getPossibleMoves(node *tr.Node) []int {
	poss := []int{}
	for i := 0; i < (board.SIZE * board.SIZE); i++ {
		if node.Tab[i] == 0 {
			if CheckValidMove(i, node.Tab) {
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
					poss = append(poss, i)
				}
			}
		}
	}
	return poss
}

func addNewLayerPrePruning(poss []int, node *tr.Node, playerId int) {
	newMovesToTest := []*tr.Node{}
	for i, m := range poss {
		newTab := node.Tab
		newTab[m] = playerId
		captureDirections := checkCapture(m, &node.Tab, playerId)
		// Virtual Capturing
		capturePairs(m, captureDirections, &newTab)
		new := tr.Node{
			Id:       i + node.Id,
			Value:    0,
			Tab:      newTab,
			Position: m,
			Player:   playerId,
		}
		new.Value = getHeuristicValue(new.Position, playerId, &new.Tab)
		newMovesToTest = append(newMovesToTest, &new)
	}
	sort.Slice(newMovesToTest, func(i int, j int) bool {
		return newMovesToTest[i].Value > newMovesToTest[j].Value
	})

	i := 0
	for i < 4 {
		if i < len(newMovesToTest) {
			newMovesToTest[i].Value = 0
			tr.AddChild(node, newMovesToTest[i])
		}
		i++
	}
}

func SuggestMove() {

	depth := 5

	if Game.CurrentPlayer.Id == 2 {
		if Game.Turn == 1 {
			center := (board.SIZE * board.SIZE) / 2
			if board.SIZE%2 == 0 {
				center += board.SIZE / 2
			}
			Game.SuggestedPosition = center + 1
		} else {
			Game.SuggestedPosition = tree.BestChild.BestChild.Position
		}
		Game.SuggestedPosition = board.TOT_SIZE + 1
		return
	}
	startTime := time.Now()

	//Create tree
	tree = tr.Node{Id: 1, Value: 0, Tab: Game.Goban.Tab, Player: Game.CurrentPlayer.Id}
	poss := getPossibleMoves(&tree)
	opponent := 1
	if Game.CurrentPlayer.Id == 1 {
		opponent = 2
	}

	// UGLY ----- test. Do it the smart way :)
	// Players move
	addNewLayerPrePruning(poss, &tree, Game.CurrentPlayer.Id)

	// opponents move
	for _, ch := range tree.Children {
		poss := getPossibleMoves(ch)
		addNewLayerPrePruning(poss, ch, opponent)
	}
	// Players move
	for _, ch := range tree.Children {
		for _, ch2 := range ch.Children {
			poss := getPossibleMoves(ch2)
			addNewLayerPrePruning(poss, ch2, Game.CurrentPlayer.Id)
		}
	}
	// opponents move
	for _, ch := range tree.Children {
		for _, ch2 := range ch.Children {
			for _, ch3 := range ch2.Children {
				poss := getPossibleMoves(ch3)
				addNewLayerPrePruning(poss, ch3, opponent)
			}
		}
	}
	// Players move
	for _, ch := range tree.Children {
		for _, ch2 := range ch.Children {
			for _, ch3 := range ch2.Children {
				for _, ch4 := range ch3.Children {
					poss := getPossibleMoves(ch4)
					addNewLayerPrePruning(poss, ch4, Game.CurrentPlayer.Id)
				}
			}
		}
	}
	// // opponents move
	// for _, ch := range tree.Children {
	// 	for _, ch2 := range ch.Children {
	// 		for _, ch3 := range ch2.Children {
	// 			for _, ch4 := range ch3.Children {
	// 				for _, ch5 := range ch4.Children {
	// 					poss := getPossibleMoves(ch5)
	// 					addNewLayerPrePruning(poss, ch5, opponent)
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	// // Players move
	// for _, ch := range tree.Children {
	// 	for _, ch2 := range ch.Children {
	// 		for _, ch3 := range ch2.Children {
	// 			for _, ch4 := range ch3.Children {
	// 				for _, ch5 := range ch4.Children {
	// 					for _, ch6 := range ch5.Children {
	// 						poss := getPossibleMoves(ch6)
	// 						addNewLayerPrePruning(poss, ch6, Game.CurrentPlayer.Id)
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	startTimeAlgo := time.Now()

	// Launch algo
	LaunchMinimaxPruning(&tree, depth)

	Game.SuggestedPosition = tree.BestChild.Position
	duration := time.Since(startTime)
	durationAlgo := time.Since(startTimeAlgo)
	fmt.Println("Time spent on suggestion:", duration)
	fmt.Println("Time spent on minimax/pruning:", durationAlgo)
	fmt.Println(tree.BestChild.Position, "(", tree.BestChild.Value, ")", "->",
		tree.BestChild.BestChild.Position, "(", tree.BestChild.Value, ")", "->",
		tree.BestChild.BestChild.BestChild.Position, "(", tree.BestChild.Value, ")", "->",
		tree.BestChild.BestChild.BestChild.BestChild.Position, "(", tree.BestChild.Value, ")",
	)
}
