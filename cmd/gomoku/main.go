package main

import (
	"fmt"
	"os"

	"github.com/gogogomoku/gomoku/internal/brain"
	"github.com/gogogomoku/gomoku/internal/server"

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
