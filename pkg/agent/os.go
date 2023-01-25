// Copyright 2022 Listware

package agent

import (
	"git.fg-tech.ru/listware/go-core/pkg/client/system"
	"git.fg-tech.ru/listware/go-core/pkg/module"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types"
)

func (a *Agent) createOs(ctx module.Context) (err error) {
	create, err := system.CreateChild(ctx.Self().Id, types.OsID, types.OsLink, a.os)
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

func (a *Agent) updateOs(ctx module.Context, uuid string) (err error) {
	update, err := system.UpdateObject(uuid, a.os)
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
