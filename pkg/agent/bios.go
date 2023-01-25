// Copyright 2022 Listware

package agent

import (
	"git.fg-tech.ru/listware/go-core/pkg/client/system"
	"git.fg-tech.ru/listware/go-core/pkg/module"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types"
)

func (a *Agent) createBios(ctx module.Context) (err error) {
	create, err := system.CreateChild(ctx.Self().Id, types.BiosID, types.BiosLink, a.bios)
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

func (a *Agent) updateBios(ctx module.Context, uuid string) (err error) {
	update, err := system.UpdateObject(uuid, a.bios)
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
