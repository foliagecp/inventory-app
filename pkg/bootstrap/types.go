// Copyright 2022 Listware

package bootstrap

import (
	"context"
	"fmt"

	"git.fg-tech.ru/listware/cmdb/pkg/cmdb/qdsl"
	"git.fg-tech.ru/listware/cmdb/pkg/cmdb/vertex/types"
	"git.fg-tech.ru/listware/go-core/pkg/client/system"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/baseboard"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/bios"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/cpu"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/mem"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/netlink"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/node"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/os"
	"git.fg-tech.ru/listware/proto/sdk/pbcmdb"
)

var (
	registerTypes = []*pbcmdb.RegisterTypeMessage{}
)

func createTypes(ctx context.Context) (err error) {
	if err = createNodeType(ctx); err != nil {
		return
	}
	if err = createBiosType(ctx); err != nil {
		return
	}
	if err = createBaseboardType(ctx); err != nil {
		return
	}
	if err = createCpuType(ctx); err != nil {
		return
	}
	if err = createOsType(ctx); err != nil {
		return
	}
	if err = createNetlinkType(ctx); err != nil {
		return
	}
	if err = createMemType(ctx); err != nil {
		return
	}
	if err = createNodeContainerType(ctx); err != nil {
		return
	}

	return
}

func createType(ctx context.Context, pt *types.Type) (err error) {
	query := fmt.Sprintf("%s.types.root", pt.Schema.Title)
	elements, err := qdsl.Qdsl(ctx, query)
	if err != nil {
		return
	}

	// TODO already exists
	if len(elements) > 0 {
		return
	}

	message, err := system.RegisterType(pt, true)
	if err != nil {
		return
	}

	registerTypes = append(registerTypes, message)
	return
}

func createNodeType(ctx context.Context) (err error) {
	pt := types.ReflectType(&node.Node{})
	return createType(ctx, pt)
}
func createBiosType(ctx context.Context) (err error) {
	pt := types.ReflectType(&bios.BIOS{})
	return createType(ctx, pt)
}
func createBaseboardType(ctx context.Context) (err error) {
	pt := types.ReflectType(&baseboard.Baseboard{})
	return createType(ctx, pt)
}
func createCpuType(ctx context.Context) (err error) {
	pt := types.ReflectType(&cpu.CPU{})
	return createType(ctx, pt)
}
func createOsType(ctx context.Context) (err error) {
	pt := types.ReflectType(&os.OS{})
	return createType(ctx, pt)
}
func createNetlinkType(ctx context.Context) (err error) {
	pt := types.ReflectType(&netlink.Netlink{})
	return createType(ctx, pt)
}
func createMemType(ctx context.Context) (err error) {
	pt := types.ReflectType(&mem.MemoryDevice{})
	return createType(ctx, pt)
}

func createNodeContainerType(ctx context.Context) (err error) {
	pt := types.ReflectType(&NodeContainer{})
	return createType(ctx, pt)
}
