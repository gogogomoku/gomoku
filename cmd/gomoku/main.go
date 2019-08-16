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
	"github.com/pkg/profile"
)

func main() {
	defer profile.Start().Stop()
	parser := argparse.NewParser("gomoku", "A great Gomoku game, and solving algorithm")
	s := parser.Flag("s", "server", &argparse.Options{Help: "Launch web server"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	if *s || os.Getenv("RUNSERVER") == "true" {
		server.StartServer()
	} else {
		fmt.Println("Start gomoku | no server")
		brain.StartRound()
		for brain.Game.Winner == 0 {
			brain.SuggestMove()
			fmt.Println("Suggestion: ", brain.Game.SuggestedPosition)
			fmt.Println("Player's turn: ", brain.Game.CurrentPlayer.Id)
			if brain.Game.CurrentPlayer.Id == 1 {
				brain.HandleMove(brain.Game.CurrentPlayer.Id, brain.Game.SuggestedPosition)
			} else {
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("Enter new move: ")
				text, _ := reader.ReadString('\n')
				choice, _ := strconv.Atoi(text[:len(text)-1])
				fmt.Println(choice)
				brain.HandleMove(brain.Game.CurrentPlayer.Id, choice)
			}
			board.PrintBoard(brain.Game.Goban.Tab, brain.Game.Goban.Size)
		}
	}
}
