package board

import "fmt"

const SIZE = 19

type Board struct {
	Tab  []int8
	Size int
}

func PrintTab(tab Board) {
	fmt.Println(tab)
}
