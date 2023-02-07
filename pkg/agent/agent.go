// Copyright 2023 NJWS Inc.
// Copyright 2022 Listware

package agent

import (
	"context"
	"strings"

	"git.fg-tech.ru/listware/go-core/pkg/executor"
	"git.fg-tech.ru/listware/go-core/pkg/module"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/baseboard"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/bios"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/cpu"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/mem"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/netlink"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/node"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/os"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.New()
)

type Agent struct {
	ctx    context.Context
	cancel context.CancelFunc

	// init values
	baseboard baseboard.Baseboard
	bios      bios.BIOS
	cpu       cpu.CPU
	node      node.Node
	os        os.OS

	dimmDevs map[string]mem.MemoryDevice
	links    map[string]netlink.Netlink

	executor executor.Executor

	m module.Module
}

// Run agent
func Run() (err error) {
	a := &Agent{}
	a.ctx, a.cancel = context.WithCancel(context.Background())

	if a.executor, err = executor.New(); err != nil {
		return
	}

	if a.baseboard, err = baseboard.New(); err != nil {
		return
	}

	if a.bios, err = bios.New(); err != nil {
		return
	}

	if a.cpu, err = cpu.New(); err != nil {
		return
	}

	if a.node, err = node.New(); err != nil {
		return
	}

	if a.os, err = os.New(); err != nil {
		return
	}

	if a.dimmDevs, err = mem.Inventory(); err != nil {
		return
	}

	if a.links, err = netlink.New(); err != nil {
		return
	}

	return a.run()
}

func appendPath(paths ...string) string {
	return strings.Join(paths, ".")
}

func (a *Agent) hostname() string {
	return a.node.Hostname
}

func (a *Agent) nodepath() string {
	return appendPath(a.hostname(), types.NodeContainerPath)
}

func (a *Agent) baseboardpath() string {
	return appendPath(types.BaseboardLink, a.nodepath())
}

func (a *Agent) biospath() string {
	return appendPath(types.BiosLink, a.nodepath())
}

func (a *Agent) cpupath() string {
	return appendPath(types.CpuLink, a.nodepath())
}

func (a *Agent) ospath() string {
	return appendPath(types.OsLink, a.nodepath())
}

func (a *Agent) dimmpath(dev string) string {
	return appendPath(dev, a.nodepath())
}

func (a *Agent) dimmspath() string {
	return appendPath(memMask, a.nodepath())
}

func (a *Agent) netlinkpath(os string) string {
	return appendPath(os, a.nodepath())
}

func (a *Agent) netlinkspath() string {
	return appendPath(netlinkMask, a.nodepath())
}

func (a *Agent) run() (err error) {
	defer a.executor.Close()

	log.Infof("run system agent")

	a.osSignalCtrl()

	a.m = module.New(types.Namespace, module.WithPort(31000))

	log.Infof("use port (%d)", a.m.Port())

	if err = a.m.Bind(types.FunctionType, a.workerFunction); err != nil {
		return
	}

	// // TODO move to another app
	// if err = a.m.Bind("monit", monit.Monit); err != nil {
	// 	return
	// }

	ctx, cancel := context.WithCancel(a.ctx)

	go func() {
		defer cancel()

		if err = a.m.RegisterAndListen(ctx); err != nil {
			log.Error(err)
			return
		}

	}()

	if err = a.entrypoint(); err != nil {
		return
	}

	<-ctx.Done()

	return
}
