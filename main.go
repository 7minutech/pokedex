package main

import (
	"github.com/7minutech/pokedex/cmds"
)

func main() {
	cmds.RegisterCommands()
	cmds.CreatePokedex()
	startRepl()
}
