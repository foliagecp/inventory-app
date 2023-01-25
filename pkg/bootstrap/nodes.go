// Copyright 2022 Listware

package bootstrap

import (
	"context"

	"git.fg-tech.ru/listware/cmdb/pkg/cmdb/qdsl"
	"git.fg-tech.ru/listware/go-core/pkg/client/system"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types"
	"git.fg-tech.ru/listware/proto/sdk/pbcmdb"
)

var (
	registerObjects = []*pbcmdb.RegisterObjectMessage{}
)

type InventoryFunctionContainer struct{}

type NodeContainer struct{}

func createObjects(ctx context.Context) (err error) {
	if err = createNodeContainerObject(ctx); err != nil {
		return
	}

	if err = createInventoryMountpointObject(ctx); err != nil {
		return
	}

	if err = createInitFunctionObject(ctx); err != nil {
		return
	}

	return
}

func createNodeContainerObject(ctx context.Context) (err error) {
	// check if object exists
	elements, err := qdsl.Qdsl(ctx, types.NodeContainerPath)
	if err != nil {
		return
	}

	// TODO already exists
	if len(elements) > 0 {
		return
	}

	message, err := system.RegisterObject(types.RootID, types.NodeContainerID, types.NodeContainerLink, NodeContainer{}, true, false)
	if err != nil {
		return
	}
	registerObjects = append(registerObjects, message)
	return
}

func createInventoryMountpointObject(ctx context.Context) (err error) {
	// check if object exists
	elements, err := qdsl.Qdsl(ctx, types.FunctionContainerPath)
	if err != nil {
		return
	}

	// TODO already exists
	if len(elements) > 0 {
		return
	}

	message, err := system.RegisterObject(types.FunctionsPath, types.FunctionContainerID, types.FunctionContainerLink, InventoryFunctionContainer{}, false, true)
	if err != nil {
		return
	}
	registerObjects = append(registerObjects, message)

	return
}
