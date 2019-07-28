package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/gogogomoku/gomoku/internal/board"
	"github.com/gogogomoku/gomoku/internal/server"

	"github.com/gogogomoku/gomoku/internal/brain"

	"github.com/akamensky/argparse"
)

func main() {
	parser := argparse.NewParser("gomoku", "A great Gomoku game, and solving algorithm")
	s := parser.Flag("s", "server", &argparse.Options{Help: "Launch web server"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	if *s {
		server.StartServer()
		for {

		}
	} else {
		fmt.Println("Start gomoku | no server")
		brain.StartRound()
		for brain.GameRound.Winner == 0 {
			brain.SuggestMove()
			fmt.Println("Suggestion: ", brain.GameRound.SuggestedPosition)
			fmt.Println("Player's turn: ", brain.GameRound.CurrentPlayer.Id)
			if brain.GameRound.CurrentPlayer.Id == 1 {
				brain.HandleMove(brain.GameRound.CurrentPlayer.Id, brain.GameRound.SuggestedPosition)
			} else {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Enter new move: ")
				text, _ := reader.ReadString('\n')
				choice, _ := strconv.Atoi(text[:len(text)-1])
				fmt.Println(choice)
				brain.HandleMove(brain.GameRound.CurrentPlayer.Id, choice)
			}
			board.PrintBoard(brain.GameRound.Goban)
		}
	}
}
