package brain

import (
	"fmt"

	tr "github.com/gogogomoku/gomoku/internal/tree"
)

const MaxInt = int16(^uint16(0) >> 1)
const MinInt = -MaxInt - 1

var nodeCounter int
var totalDepth int16

func minimum(a, b int16) int16 {
	if a < b {
		return a
	}
	return b
}

func maximum(a, b int16) int16 {
	if a > b {
		return a
	}
	return b
}

func getMaxChild(node *tr.Node, depth int16, max bool, alpha int16, beta int16) int16 {
	bestValue := MinInt
	for _, ch := range node.Children {
		returnValue := int16(MinimaxRecursivePruning(ch, depth-1, !max, alpha, beta))
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

func getMinChild(node *tr.Node, depth int16, max bool, alpha int16, beta int16) int16 {
	bestValue := MaxInt
	for _, ch := range node.Children {
		returnValue := int16(MinimaxRecursivePruning(ch, depth-1, !max, alpha, beta))
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

func MinimaxRecursivePruning(node *tr.Node, depth int16, max bool, alpha int16, beta int16) int16 {
	nodeCounter++
	if depth == 0 || len(node.Children) == 0 || node.WinMove {
		node.Value = getHeuristicValue(node.Player, &node.Tab, &node.Captured)
		// fmt.Println("Position: ", node.Position, "Value: ", node.Value)
		if node.Value >= 19000 {
			node.WinMove = true
		}
		return node.Value //- 2*node.Value/int16(totalDepth-depth)
	}
	if max {
		node.Value = getMaxChild(node, depth, max, alpha, beta)
		return int16(float64(node.Value))
	} else {
		node.Value = getMinChild(node, depth, max, alpha, beta)
		return int16(float64(node.Value))
	}
}

func LaunchMinimaxPruning(graph *tr.Node, depth int16) {
	alpha := MinInt
	beta := MaxInt
	nodeCounter = 0
	totalDepth = depth
	fmt.Println("Launching Minimax with depth ", depth)
	MinimaxRecursivePruning(graph, depth, true, alpha, beta)
	fmt.Println("==========================")
	fmt.Println("Final value:", graph.Value)
	fmt.Println("Nodes checked:", nodeCounter)
}
