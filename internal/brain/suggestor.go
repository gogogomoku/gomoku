package brain

import (
	"fmt"
	// "sort"
	// "sync"
	// "time"

	"github.com/gogogomoku/gomoku/internal/board"
	tr "github.com/gogogomoku/gomoku/internal/tree"
)

var tree *tr.Node

func getPossibleMoves(tab *[board.TOT_SIZE]int16, playerId int16) []int16 {
	poss := []int16{}
	for i := int16(0); i < (board.SIZE * board.SIZE); i++ {
		if tab[i] == 0 {
			if CheckValidMove(i, *tab, playerId) {
				affectsTab := false
				lines := CheckNextN(i, *tab, 1)
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

func build_tree(depth int16, playerId int16) {

	opponent := int16(1)
	if playerId == 1 {
		opponent = 2
	}
	tree = &tr.Node{
		Id:     1,
		Value:  0,
		Tab:    Game.Goban.Tab,
		Player: Game.CurrentPlayer.Id,
	}
	poss := getPossibleMoves(&(tree.Tab), tree.Player)
	// fmt.Println("Found", len(poss), "possible moves")
	for _, move := range poss {
		newTab := Game.Goban.Tab
		newTab[move] = playerId
		tree.Children = append(tree.Children, &tr.Node{
			Position: move,
			Id:       2,
			Value:    0,
			Tab:      newTab,
			Player:   Game.CurrentPlayer.Id,
		})
	}
	for _, ch := range tree.Children {
		build_tree_recursively(ch, depth, opponent)
	}
}

func build_tree_recursively(node *tr.Node, depth int16, playerId int16) {
	if depth >= -1 {
		poss := getPossibleMoves(&(node.Tab), node.Player)
		// fmt.Println("Found", len(poss), "possible moves")
		for _, move := range poss {
			newTab := node.Tab
			newTab[move] = playerId
			node.Children = append(node.Children, &tr.Node{
				Position: move,
				Id:       2,
				Value:    0,
				Tab:      newTab,
				Player:   Game.CurrentPlayer.Id,
			})
		}
		depth -= 1
		for _, ch := range node.Children {
			build_tree_recursively(ch, depth-1, playerId)
		}
	}
}

func reuse_tree(depth int16, playerId int16, lastMove int16) {
	// Find lastMove and move tree head to corresponding child
	for i, ch := range tree.Children {
		if ch.Position == lastMove {
			tree = tree.Children[i]
		}
	}
	for _, ch := range tree.Children {
		reuse_tree_recursively(ch, depth, playerId)
	}
}

func reuse_tree_recursively(node *tr.Node, depth int16, playerId int16) {
	fmt.Println("Check if needed new Layer, depth=", depth)
	if depth >= -1 {
		if len(node.Children) == 0 {
			fmt.Println("Building new Layer")
			poss := getPossibleMoves(&(node.Tab), node.Player)
			// fmt.Println("Found", len(poss), "possible moves")
			for _, move := range poss {
				newTab := node.Tab
				newTab[move] = playerId
				node.Children = append(node.Children, &tr.Node{
					Position: move,
					Id:       2,
					Value:    0,
					Tab:      newTab,
					Player:   Game.CurrentPlayer.Id,
				})
			}
		}
		depth -= 1
		for _, ch := range node.Children {
			build_tree_recursively(ch, depth-1, playerId)
		}
	}
}

func SuggestMove(playerId int16, lastMove int16) {

	depth := int16(3)

	if Game.Turn == 0 {
		center := int16((board.SIZE * board.SIZE) / 2)
		if board.SIZE%2 == 0 {
			center += board.SIZE / 2
		}
		Game.SuggestedPosition = center
		return
	}
	if tree == nil {
		fmt.Println("   Creating tree")
		build_tree(depth, playerId)

		// // Check depth
		// if len(tree.Children[0].Children[0].Children[0].Children[0].Children) > 0 {
		// 	fmt.Println("   Tree built to 5 levels")
		// }

		// Launch algo
		LaunchMinimaxPruning(tree, depth)
		Game.SuggestedPosition = tree.BestChild.Position

		fmt.Println()
	} else {
		fmt.Println("   Reusing tree")
		build_tree(depth, playerId)

		// if len(tree.Children[0].Children[0].Children[0].Children[0].Children) > 0 {
		// 	fmt.Println("   Tree built to 5 levels")
		// }

		// Launch algo
		LaunchMinimaxPruning(tree, depth)
		Game.SuggestedPosition = tree.BestChild.Position
	}
	fmt.Println()

	tree = nil

	return

	// Launch algo
	LaunchMinimaxPruning(tree, depth)

}
