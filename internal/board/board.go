package board

import "fmt"

type Board struct {
	Tab  []byte
	Size int
}

func PrintTab(tab Board) {
	fmt.Println(tab)
}
