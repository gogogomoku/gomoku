package brain

import (
	"fmt"
	"log"
	"sort"
	"sync"
	"time"

	"github.com/gogogomoku/gomoku/internal/board"
	tr "github.com/gogogomoku/gomoku/internal/tree"
)

var tree tr.Node

func getPossibleMoves(node *tr.Node, playerId int16) []int16 {
	poss := []int16{}
	for i := int16(0); i < (board.SIZE * board.SIZE); i++ {
		if node.Tab[i] == 0 {
			if CheckValidMove(i, node.Tab, playerId) {
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

func addNewLayerPrePruning(poss []int16, node *tr.Node, playerId int16) {
	maxTestingMoves := 6
	newMovesToTest := []*tr.Node{}
	for i, m := range poss {
		new := tr.Node{
			Id:    int16(i) + node.Id,
			Value: 0,
			// Tab:      newTab,
			Position: m,
			Player:   playerId,
		}
		new.Tab = node.Tab
		new.Tab[m] = playerId
		new.Captured[1] = node.Captured[1]
		new.Captured[2] = node.Captured[2]
		// Virtual Capturing
		captureDirections := checkCapture(m, &node.Tab, playerId)
		new.Captured[playerId] += 2 * int16(len(captureDirections))
		capturePairs(m, captureDirections, &new.Tab)
		new.Value = getHeuristicValue(playerId, &new.Tab, &new.Captured)
		newMovesToTest = append(newMovesToTest, &new)
		if len(newMovesToTest) > maxTestingMoves {
			sort.Slice(newMovesToTest, func(i int, j int) bool {
				return newMovesToTest[i].Value > newMovesToTest[j].Value
			})
			newMovesToTest = newMovesToTest[:maxTestingMoves-1]
		}
	}

	i := 0
	for i < len(newMovesToTest) {
		newMovesToTest[i].Value = 0
		tr.AddChild(node, newMovesToTest[i])
		i++
	}
}

func build_tree(depth int16, playerId int16) {
	//Create tree root
	tree = tr.Node{Id: 1, Value: 0, Tab: Game.Goban.Tab, Player: Game.CurrentPlayer.Id}
	tree.Captured[1] = Game.P1.CapturedPieces
	tree.Captured[2] = Game.P2.CapturedPieces

	//Create tree first layer
	poss := getPossibleMoves(&tree, playerId)
	addNewLayerPrePruning(poss, &tree, Game.CurrentPlayer.Id)

	// //Create the rest of the tree
	// opponent := 1
	// if tree.Player == 1 {
	// 	opponent = 2
	// }
	// for _, ch := range tree.Children {
	// 	build_tree_recursive(ch, depth-1, opponent)
	// }

	//Create the rest of the tree
	opponent := int16(1)
	if tree.Player == 1 {
		opponent = 2
	}
	var waitgroup sync.WaitGroup
	for i, ch := range tree.Children {
		waitgroup.Add(1)
		tmpCh := *ch
		go func(tmpCh *tr.Node, i int16, tree *tr.Node) {
			defer waitgroup.Done()
			build_tree_recursive(tmpCh, depth-1, opponent)
			tree.Children[i] = tmpCh
		}(&tmpCh, int16(i), &tree)
	}
	waitgroup.Wait()
}

func build_tree_recursive(node *tr.Node, depth int16, playerId int16) {
	opponent := int16(1)
	if playerId == 1 {
		opponent = 2
	}
	if depth > 0 {
		currentDepth := depth - 1
		poss := getPossibleMoves(node, playerId)
		addNewLayerPrePruning(poss, node, playerId)
		for _, ch := range node.Children {
			build_tree_recursive(ch, currentDepth, opponent)
		}
	}
}

func SuggestMove(playerId int16) {

	depth := int16(5)

	if Game.CurrentPlayer.AiStatus == 0 && Game.P1.AiStatus != Game.P2.AiStatus {
		if Game.Turn == 1 {
			center := int16((board.SIZE * board.SIZE) / 2)
			if board.SIZE%2 == 0 {
				center += board.SIZE / 2
			}
			CheckValidMove(center+1, Game.Goban.Tab, playerId)
			Game.SuggestedPosition = center + 1
		} else {
			if CheckValidMove(tree.BestChild.BestChild.Position, Game.Goban.Tab, playerId) {
				Game.SuggestedPosition = tree.BestChild.BestChild.Position
			} else {
				// TODO: Delete when we have tests :D
				log.Fatalf("Invalid suggested move :: %d", tree.BestChild.BestChild.Position)
			}
		}
		return
	}
	startTime := time.Now()
	build_tree(depth, playerId)
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
