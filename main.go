package main

import (
	"os"
	"strings"
	"log"

	"github.com/urfave/cli/v2"
	"github.com/jgoett154/grocy-backup/commands"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "server",
				Value: "http://localhost",
				Usage: "URL of the Grocy server to connect to.",
			},
			&cli.StringFlag{
				Name: "api-key",
				Value: "",
				Usage: "API key used when connecting to the Grocy server.",
			},
		},
		Before: func(ctx *cli.Context) error {
			server := ctx.String("server")

			// Removing trailing slash (/)
			if strings.HasSuffix(server, "/") {
				server = server[:len(server)-1]
			}

			// Automatically direct to the API endpoint, if not given
			if !strings.HasSuffix(server, "/api") {
				server = server + "/api"
			}

			return ctx.Set("server", server)
		},
		Commands: []*cli.Command{
			{
				Name: "backup",
				Usage: "Used to create a backup of the data on the Grocy server.",
				Action: commands.Backup,
			},
			{
				Name: "restore",
				Usage: "",
				Action: commands.Restore,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Print(err)
	}
}
