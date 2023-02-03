// Copyright 2023 NJWS Inc.
// Copyright 2022 Listware

package bootstrap

import (
	"context"
	"encoding/json"

	"git.fg-tech.ru/listware/cmdb/pkg/cmdb/qdsl"
	"git.fg-tech.ru/listware/go-core/pkg/client/system"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types"
	"git.fg-tech.ru/listware/proto/sdk/pbcmdb"
	"git.fg-tech.ru/listware/proto/sdk/pbtypes"
)

var (
	registerLinks = []*pbcmdb.RegisterLinkMessage{}

	createTrigger = &pbcmdb.Trigger{
		Type: "create",
		FunctionType: &pbtypes.FunctionType{
			Namespace: types.Namespace,
			Type:      types.FunctionPath,
		},
	}

	updateTrigger = &pbcmdb.Trigger{
		Type: "update",
		FunctionType: &pbtypes.FunctionType{
			Namespace: types.Namespace,
			Type:      types.FunctionPath,
		},
	}
)

// TODO move to go-core?
type Link struct {
	Triggers map[string]map[string]*pbtypes.FunctionType
}

func (l *Link) IsExists(trigger *pbcmdb.Trigger) bool {
	if trigger, ok := l.Triggers[createTrigger.Type]; ok {
		if _, ok := trigger[createTrigger.FunctionType.Namespace+"/"+createTrigger.FunctionType.Type]; ok {
			return ok
		}
	}
	return false
}

func createLinks(ctx context.Context) (err error) {
	if err = createFunctionNodeLink(ctx); err != nil {
		return
	}
	return
}

func createFunctionNodeLink(ctx context.Context) (err error) {
	createTriggerMessage, err := system.RegisterLinkTrigger(types.FunctionID, types.NodeID, createTrigger, true)
	if err != nil {
		return
	}

	updateTriggerMessage, err := system.RegisterLinkTrigger(types.FunctionID, types.NodeID, updateTrigger, true)
	if err != nil {
		return
	}

	query := "node.function.types.root"

	elements, err := qdsl.Qdsl(ctx, query, qdsl.WithLink())
	if err != nil {
		return
	}

	for _, element := range elements {
		var link Link
		if err := json.Unmarshal(element.Link, &link); err == nil {
			if !link.IsExists(createTrigger) {
				registerLinks = append(registerLinks, createTriggerMessage)
			}
			if !link.IsExists(updateTrigger) {
				registerLinks = append(registerLinks, updateTriggerMessage)
			}
		}
		return
	}

	registerLinks = append(registerLinks, createTriggerMessage, updateTriggerMessage)

	return
}
