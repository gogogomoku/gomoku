package board

import "fmt"

const SIZE = 19
const TOT_SIZE = SIZE * SIZE

type Board struct {
	Tab  [TOT_SIZE]int
	Size int
}

func PrintBoard(tab [TOT_SIZE]int, size int) {
	for i := 0; i < size; i++ {
		fmt.Println(tab[size*i : size*(i+1)])
	}
	fmt.Println("")
}
