package board

import "fmt"

type Board struct {
	Tab  []int8
	Size int
}

func PrintTab(tab Board) {
	fmt.Println(tab)
}
