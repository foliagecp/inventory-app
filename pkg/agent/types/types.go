// Copyright 2023 NJWS Inc.
// Copyright 2022 Listware

package types

const (
	Namespace = "proxy.foliage"

	// FunctionType will be as default qdsl to function
	FunctionType = FunctionPath
	Description  = "inventory init function"
)

const (
	NodeID              = "types/node"
	CpuID               = "types/cpu"
	OsID                = "types/os"
	BaseboardID         = "types/baseboard"
	BiosID              = "types/bios"
	MemoryID            = "types/memory-device"
	NetlinkID           = "types/netlink"
	TempID              = "types/temp"
	NodeContainerID     = "types/node-container"
	FunctionContainerID = "types/function-container"
	FunctionID          = "types/function"

	RootID = "system/root"
)

const (
	CpuLink               = "cpu"
	OsLink                = "os"
	BaseboardLink         = "baseboard"
	BiosLink              = "bios"
	DimmLink              = "dimm"
	NetlinkLink           = "os-"
	TempLink              = "temp"
	NodeContainerLink     = "nodes"
	FunctionContainerLink = "inventory"
	FunctionLink          = "init"
)

const (
	NodeContainerPath     = "nodes.root"
	FunctionsPath         = "functions.root"
	FunctionContainerPath = "inventory.functions.root"
	FunctionPath          = "init.inventory.functions.root"
)
