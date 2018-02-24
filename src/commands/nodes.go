package commands

import (
	"github.com/mbrostami/estunner/src/debug"
	"github.com/urfave/cli"
)

func Nodes() cli.Command {
	return cli.Command{
		Name:  "nodes",
		Usage: "ElasticSearch nodes",
		// Category: "Read only commands",
		Action: func(c *cli.Context) error {
			debug.ListMethods(c.Args().First())
			return nil
		},
	}
}
