package brain

import (
	"fmt"
	"sort"
	"time"

	"github.com/gogogomoku/gomoku/internal/board"
	tr "github.com/gogogomoku/gomoku/internal/tree"
)

type Move struct {
	Position int
	Value    int
}

var tree tr.Node

func getPossibleMoves(node *tr.Node) []Move {
	poss := []Move{}
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
					poss = append(poss, Move{Position: i, Value: 0})
				}
			}
		}
	}
	return poss
}

func addNewLayer(poss []Move, node *tr.Node, playerId int) {
	for i, m := range poss {
		// newTab := append([]int{}, node.Tab...)
		newTab := node.Tab
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

func addNewLayerPrePrunning(poss []Move, node *tr.Node, playerId int) {
	newMovesToTest := []*tr.Node{}
	for i, m := range poss {
		// newTab := append([]int{}, node.Tab...)
		newTab := node.Tab
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
		new.Value = getHeuristicValue(new.Position, playerId, &new.Tab)
		// tr.AddChild(node, &new)
		newMovesToTest = append(newMovesToTest, &new)
	}
	// fmt.Println("POSSIBLE MOVES: ", len(poss))
	// fmt.Println("*** APPLYING PRE-PRUNNING ***")
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

	if Game.CurrentPlayer.Id == 2 {
		// if Game.Turn == 1 {
		// 	Game.SuggestedPosition = 0
		// } else {
		// 	Game.SuggestedPosition = tree.BestChild.BestChild.Position
		// }
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
	// Players move
	addNewLayerPrePrunning(poss, &tree, Game.CurrentPlayer.Id)
	// opponents move
	for _, ch := range tree.Children {
		poss := getPossibleMoves(ch)
		addNewLayerPrePrunning(poss, ch, opponent)
	}
	// Players move
	for _, ch := range tree.Children {
		for _, ch2 := range ch.Children {
			poss := getPossibleMoves(ch2)
			addNewLayerPrePrunning(poss, ch2, Game.CurrentPlayer.Id)
		}
	}
	// opponents move
	for _, ch := range tree.Children {
		for _, ch2 := range ch.Children {
			for _, ch3 := range ch2.Children {
				poss := getPossibleMoves(ch3)
				addNewLayerPrePrunning(poss, ch3, opponent)
			}
		}
	}
	// Players move
	for _, ch := range tree.Children {
		for _, ch2 := range ch.Children {
			for _, ch3 := range ch2.Children {
				for _, ch4 := range ch3.Children {
					poss := getPossibleMoves(ch4)
					addNewLayerPrePrunning(poss, ch4, Game.CurrentPlayer.Id)
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
	// 					addNewLayerPrePrunning(poss, ch5, opponent)
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
	// 						addNewLayerPrePrunning(poss, ch6, Game.CurrentPlayer.Id)
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }

	// Launch algo
	LaunchMinimaxPruning(&tree, 5)

	Game.SuggestedPosition = tree.BestChild.Position
	duration := time.Since(startTime)
	fmt.Println("Time spent on suggestion:", duration)
}
