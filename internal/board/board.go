package board

import "fmt"

const SIZE = 19

type Board struct {
	Tab  [19 * 19]int
	Size int
}

func PrintBoard(tab [19 * 19]int, size int) {
	for i := 0; i < size; i++ {
		fmt.Println(tab[size*i : size*(i+1)])
	}
	fmt.Println("")
}
