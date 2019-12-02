package brain

import (
	// "fmt"
	"sort"
	"sync"

	"github.com/gogogomoku/gomoku/internal/board"
	tr "github.com/gogogomoku/gomoku/internal/tree"
)

func initializeNode(node *tr.Node, move int16, playerId int16) *tr.Node {
	new := &tr.Node{
		Position: move,
		Value:    0,
		Player:   playerId,
	}
	new.Tab = node.Tab
	new.Tab[move] = playerId
	new.Captured[1] = node.Captured[1]
	new.Captured[2] = node.Captured[2]

	// Virtual Capturing
	captureDirections := checkCapture(move, &node.Tab, playerId)
	new.Captured[playerId] += 2 * int16(len(captureDirections))
	capturePairs(move, captureDirections, &new.Tab)
	new.Value = getHeuristicValue(playerId, &new.Tab, &new.Captured)
	return new
}

func initializeRootNode(playerId int16) *tr.Node {
	root := &tr.Node{
		Value:  0,
		Tab:    Game.Goban.Tab,
		Player: Game.CurrentPlayer.Id,
	}
	root.Captured[1] = Game.P1.CapturedPieces
	root.Captured[2] = Game.P2.CapturedPieces
	return root
}

func getPossibleMoves(tab *[board.TOT_SIZE]int16, playerId int16) []int16 {
	poss := []int16{}
	possMap := map[int16]int16{}
	for i := int16(0); i < (board.SIZE * board.SIZE); i++ {
		if tab[i] == 0 {
			if CheckValidMove(i, *tab, playerId) {
				lines := CheckNextN(i, *tab, 1)
				for _, line := range lines {
					for _, piece := range line {
						if piece != 0 {
							possMap[i] = 1
							break
						}
					}
				}

				// if (i > 0 && tab[i-1] != 0) ||
				// 	(i > board.SIZE-1 && tab[i-board.SIZE] != 0) ||
				// 	(i > board.SIZE && i%board.SIZE != 0 && tab[i-board.SIZE-1] != 0) ||
				// 	(i > board.SIZE && i%board.SIZE != board.SIZE-1 && tab[i-board.SIZE+1] != 0) ||
				// 	(i < board.TOT_SIZE-1 && tab[i+1] != 0) ||
				// 	(i < board.TOT_SIZE-1-board.SIZE && tab[i+board.SIZE] != 0) ||
				// 	(i < board.TOT_SIZE-1-board.SIZE && i%board.SIZE != board.SIZE-1 && tab[i+board.SIZE+1] != 0) ||
				// 	(i < board.TOT_SIZE-1-board.SIZE && i%board.SIZE != 0 && tab[i+board.SIZE-1] != 0) {
				// 	poss = append(poss, i)
				// }
			}
		}
	}
	// Avoid duplicates using map
	for ind, _ := range possMap {
		poss = append(poss, ind)
	}
	sort.Slice(poss, func(i int, j int) bool {
		return poss[i] > poss[j]
	})
	return poss
}

func getOpponent(playerId int16) int16 {
	if playerId == 1 {
		return 2
	}
	return 1
}

func addBestChildrenToNode(newMovesToTest *[]*tr.Node, node *tr.Node) {
	for _, n := range *newMovesToTest {
		n.Value = 0
		tr.AddChild(node, n)
	}
}

func addRemainingLayers(depth int16, playerId int16, rebuild bool) {

	// Define opponent
	opponent := int16(1)
	if playerId == 1 {
		opponent = 2
	}

	var waitgroup sync.WaitGroup
	for i, ch := range tree.Children {
		waitgroup.Add(1)
		tmpCh := *ch
		go func(tmpCh *tr.Node, i int16, tree *tr.Node) {
			defer waitgroup.Done()
			if !rebuild {
				buildTreeRecursive(tmpCh, depth, opponent)
			} else {
				reuseTreeRecursive(tmpCh, depth, opponent)
			}
			tree.Children[i] = tmpCh
		}(&tmpCh, int16(i), tree)
	}
	waitgroup.Wait()
}
