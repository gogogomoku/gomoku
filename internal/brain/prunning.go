package brain

import (
	"fmt"

	tr "github.com/gogogomoku/gomoku/internal/tree"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

var nodeCounter int

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxChild(node *tr.Node, depth int, max bool, alpha int, beta int) int {
	bestValue := MinInt
	for _, ch := range node.Children {
		returnValue := MinimaxRecursivePruning(ch, depth-1, !max, alpha, beta)
		alpha = maximum(alpha, returnValue)
		bestValue = maximum(bestValue, returnValue)
		if node.SelectedChild == nil || ch.Value > node.SelectedChild.Value {
			node.SelectedChild = ch
		}
		if beta <= alpha {
			break
		}
	}
	return bestValue
}

func getMinChild(node *tr.Node, depth int, max bool, alpha int, beta int) int {
	bestValue := MaxInt
	for _, ch := range node.Children {
		returnValue := MinimaxRecursivePruning(ch, depth-1, !max, alpha, beta)
		beta = minimum(beta, returnValue)
		bestValue = minimum(bestValue, returnValue)
		if node.SelectedChild == nil || ch.Value < node.SelectedChild.Value {
			node.SelectedChild = ch
		}
		if beta <= alpha {
			break
		}
	}
	return bestValue
}

func MinimaxRecursivePruning(node *tr.Node, depth int, max bool, alpha int, beta int) int {
	nodeCounter++
	if depth == 0 || len(node.Children) == 0 {
		return getHeuristicValue(node.Position, node.Player, &node.Tab)
	}
	if max {
		node.Value = getMaxChild(node, depth, max, alpha, beta)
		return node.Value
	} else {
		node.Value = getMinChild(node, depth, max, alpha, beta)
		return node.Value
	}
}

func LaunchMinimaxPruning(graph *tr.Node, depth int) {
	alpha := MinInt
	beta := MaxInt
	nodeCounter = 0
	fmt.Println("Launching Minimax with depth ", depth)
	MinimaxRecursivePruning(graph, depth, true, alpha, beta)
	fmt.Println("==========================")
	fmt.Println("Final value:", graph.Value)
	fmt.Println("Nodes checked:", nodeCounter)
}
