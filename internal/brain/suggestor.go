package brain

import (
	"fmt"
	"sort"
	"time"

	"github.com/gogogomoku/gomoku/internal/board"
	tr "github.com/gogogomoku/gomoku/internal/tree"
)

var tree *tr.Node
var maxPrepruningMoves int

func buildTree(depth int16, playerId int16) {

	fmt.Println("   Creating tree")
	tree = initializeRootNode(playerId)

	// Create first layer
	poss := getPossibleMoves(&(tree.Tab), tree.Player)
	for _, move := range poss {
		newTab := Game.Goban.Tab
		newTab[move] = playerId
		tree.Children = append(tree.Children, &tr.Node{
			Position: move,
			Value:    0,
			Tab:      newTab,
			Player:   Game.CurrentPlayer.Id,
		})
	}

	// Create the rest of the layers recursively and concurrently
	addRemainingLayers(depth, playerId, false)
}

func buildTreeRecursive(node *tr.Node, depth int16, playerId int16) {

	if depth >= -1 {
		newMovesToTest := []*tr.Node{}
		poss := getPossibleMoves(&(node.Tab), playerId)

		// For evey possible move, create a node
		for _, move := range poss {
			new := initializeNode(node, move, playerId)
			newMovesToTest = append(newMovesToTest, new)

			// Sort best moves and limit to maxPrepruningMoves
			if len(newMovesToTest) > maxPrepruningMoves {
				sort.Slice(newMovesToTest, func(i int, j int) bool {
					return newMovesToTest[i].Value > newMovesToTest[j].Value
				})
				newMovesToTest = newMovesToTest[:maxPrepruningMoves-1]
			}
		}
		addBestChildrenToNode(&newMovesToTest, node)

		// Continue to next layer
		depth -= 1
		for _, ch := range node.Children {
			buildTreeRecursive(ch, depth-1, getOpponent(playerId))
		}
	}
}

func reuseTree(depth int16, playerId int16, lastMove int16) {

	treeTmp := initializeRootNode(playerId)

	// Find lastMove and make children point to corresponding grandchildren
	found := false
	for i, ch := range tree.Children {
		if ch.Position == lastMove {
			found = true
			treeTmp.Children = tree.Children[i].Children
			fmt.Println("   Reusing tree")
		}
	}
	if !found {
		buildTree(depth, playerId)
		return
	}

	// Redefine the rest of the layers recursively and concurrently
	tree = treeTmp
	addRemainingLayers(depth, playerId, true)
}

func reuseTreeRecursive(node *tr.Node, depth int16, playerId int16) {

	if depth >= -1 {

		// For evey possible move, create a node
		newMovesToTest := []*tr.Node{}
		if len(node.Children) == 0 {
			poss := getPossibleMoves(&(node.Tab), playerId)
			for _, move := range poss {
				new := initializeNode(node, move, playerId)
				newMovesToTest = append(newMovesToTest, new)
				// Sort best moves and limit to maxPrepruningMoves
				if len(newMovesToTest) > maxPrepruningMoves {
					sort.Slice(newMovesToTest, func(i int, j int) bool {
						return newMovesToTest[i].Value > newMovesToTest[j].Value
					})
					newMovesToTest = newMovesToTest[:maxPrepruningMoves-1]
				}
			}
		}
		addBestChildrenToNode(&newMovesToTest, node)

		// Continue to next layer
		depth -= 1
		for _, ch := range node.Children {
			ch.Value = 0
			reuseTreeRecursive(ch, depth-1, getOpponent(playerId))
		}
	}
}

func SuggestMove(playerId int16, lastMove int16) {

	depth := int16(5)
	maxPrepruningMoves = 5
	startTimeSuggestor := time.Now()

	// For first turn, just suggest center
	if Game.Turn == 0 {
		center := int16((board.SIZE * board.SIZE) / 2)
		if board.SIZE%2 == 0 {
			center += board.SIZE / 2
		}
		Game.SuggestedPosition = center
		return
	}

	// Build the tree if it doesn't exist, or re-use it
	if tree == nil {
		buildTree(depth, playerId)
	} else {
		reuseTree(depth, playerId, lastMove)
	}
	fmt.Println()

	// Launch algo
	LaunchMinimaxPruning(tree, depth)
	durationSuggestor := time.Since(startTimeSuggestor)
	Game.SuggestionTimer = int16(durationSuggestor.Nanoseconds() / 1000000)
	Game.SuggestedPosition = tree.BestChild.Position
}
