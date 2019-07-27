package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"

	"github.com/gogogomoku/gomoku/internal/board"
	"github.com/gogogomoku/gomoku/internal/server"

	"github.com/gogogomoku/gomoku/internal/brain"

	"github.com/akamensky/argparse"
)

func main() {
	runtime.GOMAXPROCS(30)
	parser := argparse.NewParser("gomoku", "A great Gomoku game, and solving algorithm")
	s := parser.Flag("s", "server", &argparse.Options{Help: "Launch web server"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	if *s {
		go server.StartServer()
		for {

		}
	} else {
		fmt.Println("Start gomoku | no server")
		brain.StartRound()
		for brain.GameRound.Winner == 0 {
			possible := []int{}
			best := []int{}
			for i := 0; i < board.SIZE*board.SIZE; i++ {
				if brain.GameRound.Goban.Tab[i] == 0 {
					if brain.CheckValidMove(i) {
						seq := brain.CompleteSequenceForPosition(i, brain.GameRound.CurrentPlayer.Id)
						fmt.Println(i)
						fmt.Println(seq)
						if len(seq) > 0 {
							best = append(best, i)
						}
						possible = append(possible, i)
					}
				}
			}
			ran := 0
			if len(best) > 0 {
				if len(best) > 4 {
					ran = int(rand.Intn(len(best)))
				}
				brain.HandleMove(brain.GameRound.CurrentPlayer.Id, best[ran])
			} else {
				if len(possible) > 4 {
					ran = int(rand.Intn(len(possible)))
				}
				brain.HandleMove(brain.GameRound.CurrentPlayer.Id, possible[ran])
			}
			board.PrintBoard(brain.GameRound.Goban)
		}
	}
}
