// Copyright 2022 Listware

package agent

import (
	"git.fg-tech.ru/listware/go-core/pkg/client/system"
	"git.fg-tech.ru/listware/go-core/pkg/module"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/netlink"
)

func (a *Agent) netlink(uuid string) {
	updateChan, deleteChan := netlink.Subscribe(a.ctx)

	for {
		select {
		case link := <-updateChan:
			netlinkFunc, err := genNetlinkFunction(uuid, a.netlinkpath(link.LinkName()), link, updateEvent)
			if err != nil {
				log.Error(err)
				continue
			}

			if err = a.executor.ExecAsync(a.ctx, netlinkFunc); err != nil {
				log.Error(err)
				continue
			}

		case link := <-deleteChan:
			netlinkFunc, err := genNetlinkFunction(uuid, a.netlinkpath(link.LinkName()), link, deleteEvent)
			if err != nil {
				log.Error(err)
				continue
			}

			if err = a.executor.ExecAsync(a.ctx, netlinkFunc); err != nil {
				log.Error(err)
				continue
			}

		case <-a.ctx.Done():
			return
		}
	}
}

func (a *Agent) createNetlinks(ctx module.Context) (err error) {
	for _, link := range a.links {
		netlinkFunc, err := genNetlinkFunction(ctx.Self().Id, a.netlinkpath(link.LinkName()), link, updateEvent)
		if err != nil {
			return err
		}
		msg, err := module.ToMessage(netlinkFunc)
		if err != nil {
			return err
		}

		ctx.Send(msg)
	}

	go a.netlink(ctx.Self().Id)
	return
}

func (a *Agent) createNetlink(ctx module.Context, link *netlink.Netlink) (err error) {
	create, err := system.CreateChild(ctx.Self().Id, types.NetlinkID, link.LinkName(), link)
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

func (a *Agent) updateNetlink(ctx module.Context, uuid string, link netlink.Netlink) (err error) {
	update, err := system.UpdateObject(uuid, link)
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
