// Copyright 2022 Listware

package agent

import (
	"fmt"

	"git.fg-tech.ru/listware/cmdb/pkg/cmdb/documents"
	"git.fg-tech.ru/listware/cmdb/pkg/cmdb/qdsl"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types"
)

func (a *Agent) getDocument(query string) (document *documents.Node, err error) {
	documents, err := qdsl.Qdsl(a.ctx, query, qdsl.WithKey(), qdsl.WithId(), qdsl.WithType())
	if err != nil {
		return
	}
	for _, document = range documents {
		return
	}
	err = fmt.Errorf("document '%s' not found", query)
	return
}

func (a *Agent) getFunction() (document *documents.Node, err error) {
	// search function_type init 'init.exmt.functions.root'
	return a.getDocument(types.FunctionPath)
}

func (a *Agent) getNode() (document *documents.Node, err error) {
	// search function_type init 'dev0.nodes.root'
	return a.getDocument(a.nodepath())
}

func (a *Agent) getNodes() (document *documents.Node, err error) {
	// search 'nodes.root'
	return a.getDocument(types.NodeContainerPath)
}
