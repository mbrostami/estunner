package commands

import (
	//"bytes"
	"fmt"
	//"github.com/mbrostami/estunner/src/debug"
	"github.com/urfave/cli"
	"log"
	"net/http"
	//"os/exec"
)

func Status() cli.Command {
	return cli.Command{
		Name:  "status",
		Usage: "Status of ElasticSearch",
		// Category: "Read only commands",
		Action: func(c *cli.Context) error {
			log.SetFlags(0) // remove time from logs
			host := c.String("host")
			if host == "" {
				log.Fatalf("Please specify host with --host")
				return nil
			}
			resp, err := http.Get(host + "/_stats")
			// cmd := exec.Command("docker", "exec", "-t", "trumpet", "bash", "-c", "ls .")
			// var out bytes.Buffer
			// var stderr bytes.Buffer
			// cmd.Stdout = &out
			// cmd.Stderr = &stderr
			// cmd.Run()
			// log.Fatal(out.String())
			// fmt.Println(out.String())
			fmt.Println(err)
			fmt.Println(resp)
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "host",
				Usage: "Host that ESTuner will connect to",
			},
		},
	}
}
