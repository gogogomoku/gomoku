package main

import (
	"gomoku/internal/board"
	"gomoku/internal/brain"
)

func main() {
	round := brain.Initialize()
	board.PrintTab(*round.Goban)
}
