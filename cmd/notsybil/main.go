package main

import (
	"context"
	"log"

	"notsybil/command"
)

func main() {
	// create and start CLI
	cmds := command.Commands{}
	if err := cmds.Run(context.Background()); err != nil {
		log.Fatalln(err)
	}
}
