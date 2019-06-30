package main

import (
	"fmt"
	"gomoku/internal/brain"
	"gomoku/internal/server"
	"os"

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
		go server.StartServer()
	} else {
		fmt.Println("Start gomoku | no server")
		brain.StartRound()
	}
	for {

	}
}
