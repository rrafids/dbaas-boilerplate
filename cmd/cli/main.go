package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"source.golabs.io/engineering-platforms/productivity/dbaas/dspro/cmd"
	"source.golabs.io/engineering-platforms/productivity/dbaas/dspro/datastores"
)

var (
	app = &cli.App{
		Name:        "DSPRO: Datastore Provisioner",
		Description: "A datastore deployment automation tool",
		Commands: []*cli.Command{
			{
				Name:  "generate",
				Usage: "generate the terraform file from yml",
				Action: func(c *cli.Context) error {
					generator := &cmd.Generator{
						//TODO: Validation for appname as per dev portal and ENV
						Datastore: c.Args().Get(0),
						Path:      c.Args().Get(1),
					}

					// This will call cli handler function
					return cliHandler.Generate(datastores.GetDatastoreObject(c.Args().Get(0)))
				},
			},
		},
	}
)

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatalln(err.Error())
	}
}