// Copyright 2023 NJWS Inc.

package inventory

import (
	"git.fg-tech.ru/listware/inventory-app/pkg/agent"
	"git.fg-tech.ru/listware/inventory-app/pkg/bootstrap"
	"github.com/urfave/cli/v2"
)

var (
	CLI = cli.NewApp()

	version = "v0.1.0"
)

func init() {
	CLI.Usage = "Inventory tool"
	CLI.Version = version

	CLI.Commands = []*cli.Command{
		&cli.Command{
			Name:        "bootstrap",
			Description: "Bootstrap inventory",
			Action: func(ctx *cli.Context) (err error) {
				return bootstrap.Run()
			},
		},
		&cli.Command{
			Name:        "run",
			Description: "Run inventory",
			Action: func(ctx *cli.Context) (err error) {
				if err = bootstrap.Run(); err != nil {
					return
				}
				return agent.Run()
			},
		},
	}
}
