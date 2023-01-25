// Copyright 2022 Listware

package agent

import (
	"encoding/json"
	"fmt"
	"strings"

	"git.fg-tech.ru/listware/cmdb/pkg/cmdb/qdsl"
	"git.fg-tech.ru/listware/go-core/pkg/module"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types"
)

func (a *Agent) workerFunction(ctx module.Context) (err error) {
	var req Request

	if err = json.Unmarshal(ctx.Message(), &req); err != nil {
		// temp
		// move def msg to trigger
		req.Query = a.nodepath()
	}

	elements, err := qdsl.Qdsl(ctx, req.Query, qdsl.WithName(), qdsl.WithId())
	if err != nil {
		return
	}

	switch {
	case req.Query == types.NodeContainerPath:
		// find nodes
		for _, doc := range elements {
			// uuid of "nodes.root"
			return a.readNode(ctx, doc.Id.String())
		}
		return fmt.Errorf("%s: not found", types.NodeContainerPath)

	case req.Query == a.nodepath():
		// find node
		for _, doc := range elements {
			// uuid of "dev0.nodes.root._"
			return a.readChilds(ctx, doc.Id.String())
		}
		return

	case req.Query == a.baseboardpath():
		// find baseboard
		for _, doc := range elements {
			// uuid of "baseboard.dev0.nodes.root._"
			return a.updateBaseboard(ctx, doc.Id.String())
		}
		return a.createBaseboard(ctx)

	case req.Query == a.cpupath():
		// find cpu
		for _, doc := range elements {
			// uuid of "cpu.dev0.nodes.root._"
			return a.updateCpu(ctx, doc.Id.String())
		}
		return a.createCpu(ctx)

	case req.Query == a.biospath():
		// find bios
		for _, doc := range elements {
			// uuid of "bios.dev0.nodes.root._"
			return a.updateBios(ctx, doc.Id.String())
		}
		return a.createBios(ctx)

	case req.Query == a.ospath():
		// find os
		for _, doc := range elements {
			// uuid of "os.dev0.nodes.root._"
			return a.updateOs(ctx, doc.Id.String())
		}
		return a.createOs(ctx)

	case req.Query == a.dimmspath():
		for _, doc := range elements {
			if _, ok := a.dimmDevs[doc.Name]; ok {
				continue
			}
			if err = a.deleteObject(ctx, doc.Id.String()); err != nil {
				return
			}
		}
		return a.createDimms(ctx)

	case req.Query == a.netlinkspath():

		for _, doc := range elements {
			if _, ok := a.links[doc.Name]; ok {
				continue
			}

			// TODO remove if not exists
			if err = a.deleteObject(ctx, doc.Id.String()); err != nil {
				return
			}
		}
		return a.createNetlinks(ctx)

	case strings.Contains(req.Query, types.DimmLink):
		// find dimm
		for _, doc := range elements {
			// uuid of "dimm*.dev0.nodes.root._"
			return a.updateDimm(ctx, doc.Id.String(), req.Name)
		}
		return a.createDimm(ctx, req.Name)

	case strings.Contains(req.Query, types.NetlinkLink):
		for _, doc := range elements {
			// uuid of "os-*.dev0.nodes.root._"
			switch req.Event {
			case updateEvent:
				return a.updateNetlink(ctx, doc.Id.String(), req.Link)
			case deleteEvent:
				return a.deleteObject(ctx, doc.Id.String())
			}
			return
		}
		return a.createNetlink(ctx, &req.Link)
	}

	return
}
