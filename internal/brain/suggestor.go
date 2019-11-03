package brain

import (
	"fmt"
	"sort"
	"sync"
	"time"

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
	fmt.Println("   Creating tree")
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
	tree.Captured[1] = Game.P1.CapturedPieces
	tree.Captured[2] = Game.P2.CapturedPieces
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
	var waitgroup sync.WaitGroup
	for i, ch := range tree.Children {
		waitgroup.Add(1)
		tmpCh := *ch
		go func(tmpCh *tr.Node, i int16, tree *tr.Node) {
			defer waitgroup.Done()
			build_tree_recursively(tmpCh, depth, opponent)
			tree.Children[i] = tmpCh
		}(&tmpCh, int16(i), tree)
	}
	waitgroup.Wait()
}

func build_tree_recursively(node *tr.Node, depth int16, playerId int16) {
	maxTestingMoves := 6
	newMovesToTest := []*tr.Node{}
	opponent := int16(1)
	if playerId == 1 {
		opponent = 2
	}
	if depth >= -1 {
		poss := getPossibleMoves(&(node.Tab), playerId)
		nodesAnalyzed++
		// fmt.Println("Found", len(poss), "possible moves")
		for _, move := range poss {
			new := &tr.Node{
				Position: move,
				Id:       2,
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
			newMovesToTest = append(newMovesToTest, new)
			if len(newMovesToTest) > maxTestingMoves {
				sort.Slice(newMovesToTest, func(i int, j int) bool {
					return newMovesToTest[i].Value > newMovesToTest[j].Value
				})
				newMovesToTest = newMovesToTest[:maxTestingMoves-1]
			}
			// node.Children = append(node.Children, new)
		}
		i := 0
		for i < len(newMovesToTest) {
			newMovesToTest[i].Value = 0
			tr.AddChild(node, newMovesToTest[i])
			i++
		}
		depth -= 1
		for _, ch := range node.Children {
			build_tree_recursively(ch, depth-1, opponent)
		}
	}
}

func reuse_tree(depth int16, playerId int16, lastMove int16) {
	opponent := int16(1)
	if playerId == 1 {
		opponent = 2
	}
	treeTmp := &tr.Node{
		Id:     1,
		Value:  0,
		Tab:    Game.Goban.Tab,
		Player: Game.CurrentPlayer.Id,
	}
	treeTmp.Captured[1] = Game.P1.CapturedPieces
	treeTmp.Captured[2] = Game.P2.CapturedPieces

	// Find lastMove and move tree children to corresponding grandchildren
	found := false
	for i, ch := range tree.Children {
		if ch.Position == lastMove {
			found = true
			treeTmp.Children = tree.Children[i].Children
			fmt.Println("   Reusing tree")

		}
	}
	if !found {
		build_tree(depth, playerId)
		return
	}
	tree = treeTmp
	var waitgroup sync.WaitGroup
	for i, ch := range tree.Children {
		waitgroup.Add(1)
		tmpCh := *ch
		go func(tmpCh *tr.Node, i int16, tree *tr.Node) {
			defer waitgroup.Done()
			reuse_tree_recursively(tmpCh, depth, opponent)
			tree.Children[i] = tmpCh
		}(&tmpCh, int16(i), tree)
	}
	waitgroup.Wait()
}

func reuse_tree_recursively(node *tr.Node, depth int16, playerId int16) {
	maxTestingMoves := 6
	newMovesToTest := []*tr.Node{}
	opponent := int16(1)
	if playerId == 1 {
		opponent = 2
	}
	if depth >= -1 {
		if len(node.Children) == 0 {
			poss := getPossibleMoves(&(node.Tab), playerId)
			nodesAnalyzed++
			// fmt.Println("Found", len(poss), "possible moves")
			for _, move := range poss {
				new := &tr.Node{
					Position: move,
					Id:       2,
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
				newMovesToTest = append(newMovesToTest, new)
				if len(newMovesToTest) > maxTestingMoves {
					sort.Slice(newMovesToTest, func(i int, j int) bool {
						return newMovesToTest[i].Value > newMovesToTest[j].Value
					})
					newMovesToTest = newMovesToTest[:maxTestingMoves-1]
				}
				// node.Children = append(node.Children, new)
			}
		}
		i := 0
		for i < len(newMovesToTest) {
			newMovesToTest[i].Value = 0
			tr.AddChild(node, newMovesToTest[i])
			i++
		}
		depth -= 1
		for _, ch := range node.Children {
			ch.Value = 0
			reuse_tree_recursively(ch, depth-1, opponent)
		}
	}
}

var nodesAnalyzed int

func SuggestMove(playerId int16, lastMove int16) {

	depth := int16(5)

	nodesAnalyzed = 0

	if Game.Turn == 0 {
		center := int16((board.SIZE * board.SIZE) / 2)
		if board.SIZE%2 == 0 {
			center += board.SIZE / 2
		}
		Game.SuggestedPosition = center
		return
	}
	startTimeAlgo := time.Now()
	if tree == nil {
		build_tree(depth, playerId)
		fmt.Println()
	} else {
		// build_tree(depth, playerId)
		reuse_tree(depth, playerId, lastMove)
	}
	fmt.Println()

	// tree = nil

	// Launch algo
	LaunchMinimaxPruning(tree, depth)
	durationAlgo := time.Since(startTimeAlgo)
	Game.SuggestionTimer = int16(durationAlgo.Nanoseconds() / 1000000)
	Game.SuggestedPosition = tree.BestChild.Position
	fmt.Println("Nodes Analyzed: ", nodesAnalyzed)

}
