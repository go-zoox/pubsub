package main

import (
	"github.com/go-zoox/cli"
	"github.com/go-zoox/pubsub/cmd/commands"
)

func main() {
	app := cli.NewMultipleProgram(&cli.MultipleProgramConfig{
		Name:  "multiple",
		Usage: "multiple is a program that has multiple commands.",
	})

	commands.Sub(app)
	commands.Pub(app)

	app.Run()
}
