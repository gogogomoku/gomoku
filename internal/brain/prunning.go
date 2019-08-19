package brain

import (
	"fmt"

	tr "github.com/gogogomoku/gomoku/internal/tree"
)

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

var nodeCounter int
var totalDepth int

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
		if node.BestChild == nil || ch.Value > node.BestChild.Value {
			node.BestChild = ch
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
		if node.BestChild == nil || ch.Value < node.BestChild.Value {
			node.BestChild = ch
		}
		if beta <= alpha {
			break
		}
	}
	return bestValue
}

func MinimaxRecursivePruning(node *tr.Node, depth int, max bool, alpha int, beta int) int {
	nodeCounter++
	if depth == 0 || len(node.Children) == 0 || node.WinMove {
		node.Value = getHeuristicValue(node.Position, node.Player, &node.Tab)
		// fmt.Println("Position: ", node.Position, "Value: ", node.Value)
		if node.Value >= 100000 {
			node.WinMove = true
		}
		return node.Value
	}
	if max {
		node.Value = getMaxChild(node, depth, max, alpha, beta)
		return int(float64(node.Value))
	} else {
		node.Value = getMinChild(node, depth, max, alpha, beta)
		return int(float64(node.Value))
	}
}

func LaunchMinimaxPruning(graph *tr.Node, depth int) {
	alpha := MinInt
	beta := MaxInt
	totalDepth = depth
	nodeCounter = 0
	fmt.Println("Launching Minimax with depth ", depth)
	MinimaxRecursivePruning(graph, depth, true, alpha, beta)
	fmt.Println("==========================")
	fmt.Println("Final value:", graph.Value)
	fmt.Println("Nodes checked:", nodeCounter)
}
