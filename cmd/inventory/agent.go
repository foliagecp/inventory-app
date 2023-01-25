// Copyright 2022 Listware

package main

import (
	"git.fg-tech.ru/listware/inventory-app/cmd/inventory/agent"
)

// runCmd represents the run command
func init() {
	rootCmd.AddCommand(agent.RootCmd)
}
