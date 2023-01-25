// Copyright 2022 Listware

package main

import (
	"encoding/json"
	"fmt"
	"os"

	"git.fg-tech.ru/listware/inventory-app/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	version         = "v0.1.0"
	release         = "dev"
	versionTemplate = `{{printf "%s Agent" .Short}}
{{printf "Version: %s" .Version}}
Release: ` + release + `
`
)

var configFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "exmt",
	Short:   "Extended Management",
	Long:    `Coming soon...`,
	Version: version,
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	//	cobra.OnInitialize(initConfig)
	rootCmd.SetVersionTemplate(versionTemplate)
	rootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is ~/.exmt.yaml)")
	rootCmd.AddCommand(autoShellCmd, testCmd)
}

/*
// initConfig reads in config file and ENV variables if set.

	func initConfig() {
		if configFile == "" {
			if home, err := homedir.Dir(); err == nil {
				configFile = path.Join(home, ".exmt.yaml")
			}
		}
		profile.SetConfigFile(configFile)
		dir, err := filepath.Abs(filepath.Dir(configFile))
		if err != nil {
			fmt.Println(err)
		}
		os.MkdirAll(dir, 0600)
		profile.ReadInConfig()
	}
*/
var autoShellCmd = &cobra.Command{
	Use:    "autoshell",
	Short:  "Generate bash completion script",
	Long:   "Generate bash completion script",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := cmd.Root().GenBashCompletionFile("/etc/bash_completion.d/exmt.sh")
		if err != nil {
			return err
		}
		return nil
	},
}

var testCmd = &cobra.Command{
	Use:    "test",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		utils.Init()
		err = json.NewEncoder(os.Stdout).Encode(utils.Sys)
		return
	},
}
