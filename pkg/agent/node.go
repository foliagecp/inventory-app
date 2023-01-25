// Copyright 2022 Listware

package agent

import (
	"git.fg-tech.ru/listware/go-core/pkg/client/system"
	"git.fg-tech.ru/listware/go-core/pkg/module"
)

// readNode generate and exec worker("nodepath") func
func (a *Agent) readNode(ctx module.Context, id string) (err error) {
	nodeFunc, err := genFunction(id, a.nodepath())
	if err != nil {
		return
	}
	msg, err := module.ToMessage(nodeFunc)
	if err != nil {
		return
	}

	ctx.Send(msg)
	return
}

// readNode generate and exec worker("baseboard", "cpu", etc...) funcs
func (a *Agent) readChilds(ctx module.Context, id string) (err error) {
	update, err := system.UpdateObject(id, a.node)
	if err != nil {
		return
	}
	msg, err := module.ToMessage(update)
	if err != nil {
		return
	}

	ctx.Send(msg)

	baseboardFunc, err := genFunction(id, a.baseboardpath())
	if err != nil {
		return
	}
	msg, err = module.ToMessage(baseboardFunc)
	if err != nil {
		return
	}

	ctx.Send(msg)

	osFunc, err := genFunction(id, a.ospath())
	if err != nil {
		return
	}
	msg, err = module.ToMessage(osFunc)
	if err != nil {
		return
	}

	ctx.Send(msg)

	biosFunc, err := genFunction(id, a.biospath())
	if err != nil {
		return
	}
	msg, err = module.ToMessage(biosFunc)
	if err != nil {
		return
	}

	ctx.Send(msg)

	cpuFunc, err := genFunction(id, a.cpupath())
	if err != nil {
		return
	}
	msg, err = module.ToMessage(cpuFunc)
	if err != nil {
		return
	}

	ctx.Send(msg)

	dimmFunc, err := genFunction(id, a.dimmspath())
	if err != nil {
		return
	}
	msg, err = module.ToMessage(dimmFunc)
	if err != nil {
		return
	}

	ctx.Send(msg)

	netlinkFunc, err := genFunction(id, a.netlinkspath())
	if err != nil {
		return
	}
	msg, err = module.ToMessage(netlinkFunc)
	if err != nil {
		return
	}

	ctx.Send(msg)
	return
}

func (a *Agent) deleteObject(ctx module.Context, uuid string) (err error) {
	del, err := system.DeleteObject(uuid)
	if err != nil {
		return
	}
	msg, err := module.ToMessage(del)
	if err != nil {
		return
	}

	ctx.Send(msg)
	return
}
