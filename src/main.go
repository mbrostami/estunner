package main

import (
	"github.com/mbrostami/estunner/src/commands"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "estuner"
	app.Usage = "Tunning ElasticSearch"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		commands.Status(),
		commands.Nodes(),
	}

	app.Run(os.Args)
}
