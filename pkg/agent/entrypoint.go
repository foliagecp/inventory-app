// Copyright 2022 Listware

package agent

import (
	"fmt"

	"git.fg-tech.ru/listware/cmdb/pkg/cmdb/qdsl"
	"git.fg-tech.ru/listware/go-core/pkg/client/system"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types"
	"git.fg-tech.ru/listware/proto/sdk/pbtypes"
)

func (a *Agent) createLink(route *pbtypes.FunctionRoute) (err error) {
	function, err := a.getFunction()
	if err != nil {
		return
	}

	node, err := a.getNode()
	if err != nil {
		return
	}

	// create init link, will trigger 'init'
	createLink, err := system.CreateLink(function.Id.String(), node.Id.String(), node.Key, function.Type, route)
	if err != nil {
		return err
	}

	return a.executor.ExecSync(a.ctx, createLink)
}

func (a *Agent) entrypoint() (err error) {
	route := &pbtypes.FunctionRoute{
		Url: a.m.Addr(),
	}

	// exists node or create
	node, err := a.getNode()
	if err == nil {
		// search function
		documents, err := qdsl.Qdsl(a.ctx, fmt.Sprintf("%s.%s", node.Key, types.FunctionPath), qdsl.WithLinkId())
		if err != nil {
			return err
		}

		// func on node: exists
		for _, document := range documents {
			updateLink, err := system.UpdateAdvancedLink(document.LinkId.String(), route)
			if err != nil {
				return err
			}

			return a.executor.ExecSync(a.ctx, updateLink)
		}

		return a.createLink(route)
	}

	// TODO move to register
	nodes, err := a.getNodes()
	if err != nil {
		return
	}

	message := &pbtypes.FunctionMessage{
		Type:  types.FunctionPath,
		Route: route,
	}

	createNode, err := system.CreateChild(nodes.Id.String(), types.NodeID, a.hostname(), a.node, message)
	if err != nil {
		return
	}

	if err = a.executor.ExecSync(a.ctx, createNode); err != nil {
		return err
	}

	return a.createLink(route)
}
