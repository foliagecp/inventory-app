// Copyright 2022 Listware

package bootstrap

import (
	"context"

	"git.fg-tech.ru/listware/cmdb/pkg/cmdb/qdsl"
	"git.fg-tech.ru/listware/go-core/pkg/client/system"
)

type InitFunction struct{}

func createInitFunctionObject(ctx context.Context) (err error) {
	// check if object exists
	elements, err := qdsl.Qdsl(ctx, "init.inventory.functions.root")
	if err != nil {
		return
	}

	// already exists
	if len(elements) > 0 {
		return
	}

	message, err := system.RegisterObject("inventory.functions.root", "types/function", "init", InitFunction{}, true, true)
	if err != nil {
		return
	}
	registerObjects = append(registerObjects, message)

	return
}
