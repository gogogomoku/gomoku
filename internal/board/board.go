package board

import "fmt"

type Board struct {
	Tab  []byte
	Size int
}

// var Tab = Board{make([]byte, 19*19), 19}

// func GetTab() *Board {
// 	return &Tab
// }

func PrintTab(tab Board) {
	fmt.Println(tab)
}
