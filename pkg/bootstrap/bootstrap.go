// Copyright 2022 Listware

package bootstrap

import (
	"context"
	"fmt"

	"git.fg-tech.ru/listware/go-core/pkg/client/system"
	"git.fg-tech.ru/listware/go-core/pkg/executor"
)

const (
	appName = "inventory"
)

func register(ctx context.Context, exec executor.Executor) (err error) {
	// create types
	if err = createTypes(ctx); err != nil {
		return
	}

	// create objects
	if err = createObjects(ctx); err != nil {
		return
	}

	// create links
	if err = createLinks(ctx); err != nil {
		return
	}

	message, err := system.Register(appName, registerTypes, registerObjects, registerLinks)
	if err != nil {
		return
	}

	return exec.ExecSync(ctx, message)
}

func Run() (err error) {
	ctx := context.Background()

	exec, err := executor.New()
	if err != nil {
		return
	}
	defer exec.Close()

	if err = register(ctx, exec); err != nil {
		fmt.Println("register: ", err)
		return
	}

	return
}
