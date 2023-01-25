// Copyright 2022 Listware

package main

import (
	"git.fg-tech.ru/listware/inventory-app/cmd/inventory/bootstrap"
)

// runCmd represents the run command
func init() {
	rootCmd.AddCommand(bootstrap.RootCmd)
}
