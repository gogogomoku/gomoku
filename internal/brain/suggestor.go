package brain

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gogogomoku/gomoku/internal/board"

	bolt "github.com/gogogomoku/gomoku/internal/boltdb"
	tr "github.com/gogogomoku/gomoku/internal/tree"
)

var tree *tr.Node
var maxPrepruningMoves int

func buildTree(depth int16, playerId int16) {

	fmt.Println("   Creating tree")
	tree = initializeRootNode(playerId)

	// Create first layer
	newMovesToTest := []*tr.Node{}
	poss := getPossibleMoves(&(tree.Tab), tree.Player)
	for _, move := range poss {
		newTab := Game.Goban.Tab
		newTab[move] = playerId
		newMovesToTest = append(newMovesToTest, &tr.Node{
			Position: move,
			Value:    0,
			Tab:      newTab,
			Player:   Game.CurrentPlayer.Id,
			WinMove:  false,
		})

		// Sort best moves and limit to maxPrepruningMoves
		if len(newMovesToTest) > maxPrepruningMoves {
			sort.Slice(newMovesToTest, func(i int, j int) bool {
				return newMovesToTest[i].Value > newMovesToTest[j].Value
			})
			newMovesToTest = newMovesToTest[:maxPrepruningMoves-1]
		}
	}
	addBestChildrenToNode(&newMovesToTest, tree)
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
			break
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

	fmt.Println("Make suggestion")

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

	var hash string
	if Game.CacheEnabled {
		h := sha256.New()
		tabToStr := strings.Trim(strings.Replace(fmt.Sprint(Game.Goban.Tab), " ", "", -1), "[]")
		uid := fmt.Sprintf(
			"%v%d%d%d",
			tabToStr,
			Game.CurrentPlayer.Id,
			Game.CurrentPlayer.CapturedPieces,
			Game.GetCurrentOpponent().CapturedPieces,
		)
		fmt.Println(uid)
		h.Write([]byte(uid))
		hash = fmt.Sprintf("%x", h.Sum(nil))
		startTime := time.Now()
		str := bolt.Get(bolt.Bolt.Bucket, hash)
		fmt.Println("Return from cache:", str)
		if str != "none" {
			fmt.Println("-- Found in cache")
			cacheSuggestion, err := strconv.ParseInt(str, 10, 64)
			if err == nil {
				fmt.Println("-- Returning cache suggestion")
				Game.SuggestedPosition = int16(cacheSuggestion)
				duration := time.Since(startTime)
				Game.SuggestionTimer = int16(duration.Nanoseconds() / 1000000)
				return
			} else {
				fmt.Println(err.Error())
			}
		}
	}

	// Build the tree if it doesn't exist, or re-use it
	if tree == nil || Game.Turn == 0 {
		buildTree(depth, playerId)
	} else {
		reuseTree(depth, playerId, lastMove)
	}
	fmt.Println()

	// Launch algo
	tree.BestChild = nil
	for _, ch := range tree.Children {
		if ch.WinMove {
			tree.BestChild = ch
			break
		}
	}
	if tree.BestChild == nil {
		LaunchMinimaxPruning(tree, depth)
	}
	durationSuggestor := time.Since(startTimeSuggestor)
	Game.SuggestionTimer = int16(durationSuggestor.Nanoseconds() / 1000000)
	if Game.CacheEnabled {
		bolt.Put(bolt.Bolt.Bucket, hash, fmt.Sprint(tree.BestChild.Position))
		// bolt.PrintBucket(bolt.Bolt.Bucket)
	}

	Game.SuggestedPosition = tree.BestChild.Position
	// fmt.Println("First row")
	// for _, ch := range tree.Children {
	// 	fmt.Println(ch.Position, ch.Value)
	// }
	// fmt.Println("Second row (from best child)", tree.BestChild.Position)
	// for _, ch := range tree.BestChild.Children {
	// 	fmt.Println(ch.Position, ch.Value)
	// }
	// fmt.Println("Third row (from best child)", tree.BestChild.BestChild.Position)
	// for _, ch := range tree.BestChild.BestChild.Children {
	// 	fmt.Println(ch.Position, ch.Value)
	// }
	// fmt.Println("Fourth row (from best child)", tree.BestChild.BestChild.BestChild.Position)
	// for _, ch := range tree.BestChild.BestChild.BestChild.Children {
	// 	fmt.Println(ch.Position, ch.Value)
	// }
	// fmt.Println("Fifth row (from best child)", tree.BestChild.BestChild.BestChild.BestChild.Position)
	// for _, ch := range tree.BestChild.BestChild.BestChild.BestChild.Children {
	// 	fmt.Println(ch.Position, ch.Value)
	// }
	// fmt.Println("Final state:")
	// board.PrintBoard(tree.BestChild.BestChild.BestChild.BestChild.BestChild.Tab, board.SIZE)
}
