package board

import "fmt"

const SIZE = 19
const TOT_SIZE = SIZE * SIZE

type Board struct {
	Tab  [TOT_SIZE]int16
	Size int16
}

func PrintBoard(tab [TOT_SIZE]int16, size int16) {
	var i int16
	for i = 0; i < size; i++ {
		fmt.Println(tab[size*i : size*(i+1)])
	}
	fmt.Println("")
}
