package brain

import (
	// "fmt"
	"reflect"
	"testing"

	"github.com/gogogomoku/gomoku/internal/board"
	tr "github.com/gogogomoku/gomoku/internal/tree"
)

func TestLaunchMinimaxPruning(t *testing.T) {
	minimaxTest1(t)
	minimaxTest2(t)
	// /*
	// 		 **
	// 		 **                                1
	// 		 **                    -----------------------------
	// 		 **                2               3                4
	// 		 **             /    \          /     \           /   \
	// 		 **           5       6       7        8        9      10
	// 	 	 **          /  \    /  \    /  \    /   \    /  \    /  \
	// 		 **         11  12  13  14  15  16  17  18  19  20  21  22
	// */
	// center := (board.SIZE * board.SIZE) / 2
	// if board.SIZE%2 == 0 {
	// 	center += board.SIZE / 2
	// }
	// // Actual state
	// tree1 := tr.Node{Id: 1, Value: 0, Tab: [board.TOT_SIZE]int16{}, Player: 2}
	// // Player 1
	// tree2 := tr.Node{Id: 2, Value: 0, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree2, 1)
	// tree3 := tr.Node{Id: 3, Value: 0, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree3, 1)
	// tree4 := tr.Node{Id: 4, Value: 0, Tab: tree1.Tab, Player: 1}
	// // Player 2
	// tr.AddChildById(&tree1, &tree4, 1)
	// tree5 := tr.Node{Id: 5, Value: 0, Tab: tree1.Tab, Player: 2}
	// tr.AddChildById(&tree1, &tree5, 2)
	// tree6 := tr.Node{Id: 6, Value: 0, Tab: tree1.Tab, Player: 2}
	// tr.AddChildById(&tree1, &tree6, 2)
	// tree7 := tr.Node{Id: 7, Value: 0, Tab: tree1.Tab, Player: 2}
	// tr.AddChildById(&tree1, &tree7, 3)
	// tree8 := tr.Node{Id: 8, Value: 0, Tab: tree1.Tab, Player: 2}
	// tr.AddChildById(&tree1, &tree8, 3)
	// tree9 := tr.Node{Id: 9, Value: 0, Tab: tree1.Tab, Player: 2}
	// tr.AddChildById(&tree1, &tree9, 4)
	// tree10 := tr.Node{Id: 10, Value: 0, Tab: tree1.Tab, Player: 2}
	// tr.AddChildById(&tree1, &tree10, 4)
	// // Player 1
	// tree11 := tr.Node{Id: 11, Value: 0, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree11, 5)
	// tree12 := tr.Node{Id: 12, Value: 0, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree12, 5)
	// tree13 := tr.Node{Id: 13, Value: 0, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree13, 6)
	// tree14 := tr.Node{Id: 14, Value: 0, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree14, 6)
	// tree15 := tr.Node{Id: 15, Value: 0, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree15, 7)
	// tree16 := tr.Node{Id: 16, Value: 0, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree16, 7)
	// tree17 := tr.Node{Id: 17, Value: 0, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree17, 8)
	// tree18 := tr.Node{Id: 18, Value: 0, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree18, 8)
	// tree19 := tr.Node{Id: 19, Value: 0, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree19, 9)
	// tree20 := tr.Node{Id: 20, Value: 0, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree20, 9)
	// tree21 := tr.Node{Id: 21, Value: 0, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree21, 10)
	// tree22 := tr.Node{Id: 22, Value: 10, Tab: tree1.Tab, Player: 1}
	// tr.AddChildById(&tree1, &tree22, 10)
	//
	// tree11.Tab[center] = 1
	// tree11.Tab[center+1] = 1
	// tree11.Tab[center+2] = 1
	//
	// tree12.Tab[center] = 1
	// tree12.Tab[center+1] = 1
	// tree12.Tab[center+2] = 1
	//
	// LaunchMinimaxPruning(&tree1, 20)
	//
	// expectedValues := []int16{
	// 	12,
	// }
	// actualValues := []int16{
	// 	tree5.BestChild.Id,
	// }
	// if !reflect.DeepEqual(actualValues, expectedValues) {
	// 	t.Errorf("Error in Minimax Pruning. Expected: %d, got: %d", expectedValues, actualValues)
	// }
}

func minimaxTest1(t *testing.T) {
	/*
	 **
	 **                                1
	 **                    -----------------------------
	 **                2               3                4
	 */
	center := (board.SIZE * board.SIZE) / 2
	if board.SIZE%2 == 0 {
		center += board.SIZE / 2
	}
	// Actual state
	tree1 := tr.Node{Id: 1, Value: 0, Tab: [board.TOT_SIZE]int16{}, Player: 2}
	// Player 1
	tree2 := tr.Node{Id: 2, Value: 0, Tab: tree1.Tab, Player: 1}
	tr.AddChildById(&tree1, &tree2, 1)
	tree3 := tr.Node{Id: 3, Value: 0, Tab: tree1.Tab, Player: 1}
	tr.AddChildById(&tree1, &tree3, 1)
	tree4 := tr.Node{Id: 4, Value: 0, Tab: tree1.Tab, Player: 1}
	tr.AddChildById(&tree1, &tree4, 1)

	tree2.Tab[center-1] = 1
	tree2.Tab[center] = 1
	tree2.Tab[center+1] = 0
	tree2.Tab[center+2] = 1

	tree3.Tab[center-1] = 1
	tree3.Tab[center] = 1
	tree3.Tab[center+1] = 1

	tree4.Tab[center-1] = 1
	tree4.Tab[center] = 1
	tree4.Tab[center+1] = 1
	tree4.Tab[center-2] = 1

	LaunchMinimaxPruning(&tree1, 1)

	if tree2.Value != F3_SCORE+SEQ2_FREE_SCORE {
		board.PrintBoard(tree2.Tab, board.SIZE)
		t.Errorf("Error in Minimax: Heuristic. Expected: %d, got: %d", tree2.Value, F3_SCORE+SEQ2_FREE_SCORE)
	}
	if tree3.Value != F3_SCORE {
		board.PrintBoard(tree3.Tab, board.SIZE)
		t.Errorf("Error in Minimax: Heuristic. Expected: %d, got: %d", tree3.Value, F3_SCORE)
	}
	if tree4.Value != SEQ4_FREE_SCORE {
		board.PrintBoard(tree4.Tab, board.SIZE)
		t.Errorf("Error in Minimax: Heuristic. Expected: %d, got: %d", tree4.Value, SEQ4_FREE_SCORE)
	}
	expectedValues := []int16{
		4,
	}
	actualValues := []int16{
		tree1.BestChild.Id,
	}
	if !reflect.DeepEqual(actualValues, expectedValues) {
		t.Errorf("Error in Minimax Pruning. Expected: %d, got: %d", expectedValues, actualValues)
	}
}

func minimaxTest2(t *testing.T) {
	/*
	 **
	 **                                1
	 **                    -----------------------------
	 **                2               3                4
	 */
	center := (board.SIZE * board.SIZE) / 2
	if board.SIZE%2 == 0 {
		center += board.SIZE / 2
	}
	// Actual state
	tree1 := tr.Node{Id: 1, Value: 0, Tab: [board.TOT_SIZE]int16{}, Player: 2}
	// Player 1
	tree2 := tr.Node{Id: 2, Value: 0, Tab: tree1.Tab, Player: 1}
	tr.AddChildById(&tree1, &tree2, 1)
	tree3 := tr.Node{Id: 3, Value: 0, Tab: tree1.Tab, Player: 1}
	tr.AddChildById(&tree1, &tree3, 1)
	tree4 := tr.Node{Id: 4, Value: 0, Tab: tree1.Tab, Player: 1}
	tr.AddChildById(&tree1, &tree4, 1)
	// Player 2
	tree5 := tr.Node{Id: 5, Value: 0, Tab: tree1.Tab, Player: 2}
	tr.AddChildById(&tree1, &tree5, 2)
	tree6 := tr.Node{Id: 6, Value: 0, Tab: tree1.Tab, Player: 2}
	tr.AddChildById(&tree1, &tree6, 2)
	tree7 := tr.Node{Id: 7, Value: 0, Tab: tree1.Tab, Player: 2}
	tr.AddChildById(&tree1, &tree7, 3)
	tree8 := tr.Node{Id: 8, Value: 0, Tab: tree1.Tab, Player: 2}
	tr.AddChildById(&tree1, &tree8, 3)
	tree9 := tr.Node{Id: 9, Value: 0, Tab: tree1.Tab, Player: 2}
	tr.AddChildById(&tree1, &tree9, 4)
	tree10 := tr.Node{Id: 10, Value: 0, Tab: tree1.Tab, Player: 2}
	tr.AddChildById(&tree1, &tree10, 4)

	tree5.Tab[center-2] = 2
	tree5.Tab[center-1] = 2
	tree5.Tab[center] = 2
	tree5.Tab[center+1] = 2
	tree5.Tab[center+2] = 2

	tree6.Tab[center-1] = 2
	tree6.Tab[center] = 2
	tree6.Tab[center+1] = 2
	tree6.Tab[center+2] = 2

	tree7.Tab[center-1] = 2
	tree7.Tab[center+1] = 2
	tree7.Tab[center+2] = 2

	// Gets pruned
	tree8.Tab[center] = 2
	tree8.Tab[center+1] = 2
	tree8.Tab[center+2] = 2

	tree9.Tab[center] = 2
	tree9.Tab[center+1] = 2
	tree9.Tab[center+2] = 2

	// Gets pruned
	tree10.Tab[center] = 2
	tree10.Tab[center+1] = 2

	LaunchMinimaxPruning(&tree1, 2)

	if tree5.Value != WIN_SCORE {
		board.PrintBoard(tree5.Tab, board.SIZE)
		t.Errorf("Error in Minimax: Heuristic. Expected: %d, got: %d", WIN_SCORE, tree5.Value)
	}
	if tree6.Value != SEQ4_FREE_SCORE {
		board.PrintBoard(tree6.Tab, board.SIZE)
		t.Errorf("Error in Minimax: Heuristic. Expected: %d, got: %d", SEQ4_FREE_SCORE, tree6.Value)
	}
	if tree7.Value != F3_SCORE+SEQ2_FREE_SCORE {
		board.PrintBoard(tree7.Tab, board.SIZE)
		t.Errorf("Error in Minimax: Heuristic. Expected: %d, got: %d", F3_SCORE+SEQ2_FREE_SCORE, tree7.Value)
	}
	if tree8.Value != 0 {
		// Gets pruned
		board.PrintBoard(tree8.Tab, board.SIZE)
		t.Errorf("Error in Minimax: Heuristic. Expected: %d, got: %d", 0, tree8.Value)
	}
	if tree9.Value != F3_SCORE {
		board.PrintBoard(tree9.Tab, board.SIZE)
		t.Errorf("Error in Minimax: Heuristic. Expected: %d, got: %d", F3_SCORE, tree9.Value)
	}
	if tree10.Value != 0 {
		// Gets pruned
		board.PrintBoard(tree10.Tab, board.SIZE)
		t.Errorf("Error in Minimax: Heuristic. Expected: %d, got: %d", 0, tree10.Value)
	}

	if tree2.BestChild.Id != 6 {
		t.Errorf("Error in Minimax Pruning. Expected: %d, got: %d", 6, tree2.BestChild.Id)
	}
	if tree3.BestChild.Id != 7 {
		t.Errorf("Error in Minimax Pruning (Didn't PRUNE!!!!). Expected: %d, got: %d", 7, tree3.BestChild.Id)
	}
	if tree4.BestChild.Id != 9 {
		t.Errorf("Error in Minimax Pruning (Didn't PRUNE!!!!). Expected: %d, got: %d", 9, tree4.BestChild.Id)
	}

	expectedValues := []int16{
		2,
		6,
	}
	actualValues := []int16{
		tree1.BestChild.Id,
		tree1.BestChild.BestChild.Id,
	}
	if !reflect.DeepEqual(actualValues, expectedValues) {
		t.Errorf("Error in Minimax Pruning. Expected: %d, got: %d", expectedValues, actualValues)
	}
}
