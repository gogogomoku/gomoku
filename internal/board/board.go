package board

import "fmt"

const SIZE = 10

type Board struct {
	Tab  []int
	Size int
}

func PrintBoard(tab []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Println(tab[size*i : size*(i+1)])
	}
	fmt.Println("")
}
