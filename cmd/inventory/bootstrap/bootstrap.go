// Copyright 2022 Listware

package bootstrap

import (
	"git.fg-tech.ru/listware/inventory-app/pkg/bootstrap"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "Extended Management Agent Boostrap",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return bootstrap.Run()
	},
}
