// Copyright 2022 Listware

package agent

import (
	"git.fg-tech.ru/listware/go-core/pkg/client/system"
	"git.fg-tech.ru/listware/go-core/pkg/module"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types"
)

func (a *Agent) createDimms(ctx module.Context) (err error) {
	for _, dev := range a.dimmDevs {
		dimmFunc, err := genDimmFunction(ctx.Self().Id, a.dimmpath(dev.Name()), dev.Name())
		if err != nil {
			return err
		}
		msg, err := module.ToMessage(dimmFunc)
		if err != nil {
			return err
		}

		ctx.Send(msg)
	}
	return
}

func (a *Agent) createDimm(ctx module.Context, name string) (err error) {
	create, err := system.CreateChild(ctx.Self().Id, types.MemoryID, name, a.dimmDevs[name])
	if err != nil {
		return
	}
	msg, err := module.ToMessage(create)
	if err != nil {
		return
	}

	ctx.Send(msg)
	return
}

func (a *Agent) updateDimm(ctx module.Context, id, name string) (err error) {
	update, err := system.UpdateObject(id, a.dimmDevs[name])
	if err != nil {
		return
	}
	msg, err := module.ToMessage(update)
	if err != nil {
		return
	}

	ctx.Send(msg)
	return
}
