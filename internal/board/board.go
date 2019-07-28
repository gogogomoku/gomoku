package board

import "fmt"

const SIZE = 19

type Board struct {
	Tab  []int
	Size int
}

func PrintBoard(tab Board) {
	for i := 0; i < tab.Size; i++ {
		fmt.Println(tab.Tab[tab.Size*i : tab.Size*(i+1)])
	}
	fmt.Println("")
}
