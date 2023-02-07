// Copyright 2023 NJWS Inc.

package main

import (
	"fmt"
	"os"

	"git.fg-tech.ru/listware/inventory-app/pkg/inventory"
)

func main() {
	if err := inventory.CLI.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
