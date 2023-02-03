// Copyright 2023 NJWS Inc.
// Copyright 2022 Listware

package agent

import (
	"git.fg-tech.ru/listware/inventory-app/pkg/agent"
	"git.fg-tech.ru/listware/inventory-app/pkg/bootstrap"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "agent",
	Short: "Extended Management Agent",
}

// agentCmd represents the agent command
var agentRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run Extended Management Agent",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) (err error) {

		if err = bootstrap.Run(); err != nil {
			return
		}

		return agent.Run()
	},
}

// runCmd represents the run command
func init() {
	RootCmd.AddCommand(agentRunCmd)
}
