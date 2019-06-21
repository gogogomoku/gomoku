package main

import (
	"fmt"
	"gomoku/internal/board"
	"gomoku/internal/brain"
)

func main() {
	fmt.Println(brain.GameRound.Status)
	board.PrintTab(brain.GameRound.Goban)
	brain.StartRound()
	brain.HandleMove(brain.GameRound.CurrentPlayer.Id, 1)
	board.PrintTab(brain.GameRound.Goban)
	brain.HandleMove(brain.GameRound.CurrentPlayer.Id, 2)
	board.PrintTab(brain.GameRound.Goban)
	brain.HandleMove(brain.GameRound.CurrentPlayer.Id, 3)
	board.PrintTab(brain.GameRound.Goban)
}
