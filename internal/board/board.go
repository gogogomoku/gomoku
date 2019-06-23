package board

import "fmt"

type Board struct {
	Tab  []int8
	Size int
}

func PrintTab(tab Board) {
	for i := 0; i <= tab.Size; i++ {
		s := make([]byte, tab.Size)
		copy(s, tab.Tab[tab.Size*i:])
		fmt.Println(s)
	}
	fmt.Println("")
}
