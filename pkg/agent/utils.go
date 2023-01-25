// Copyright 2022 Listware

package agent

import (
	"encoding/json"

	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types"
	"git.fg-tech.ru/listware/inventory-app/pkg/agent/types/netlink"
	"git.fg-tech.ru/listware/proto/sdk/pbtypes"
)

const (
	namespace = "proxy"

	netlinkMask = "*[?@._type == 'netlink'?]"
	memMask     = "*[?@._type == 'memory-device'?]"
)

const (
	updateEvent = "update"
	deleteEvent = "delete"
)

// TODO temp
const (
	cpuDev = "cpu"
)

type Request struct {
	Query string `json:"query"`
	Name  string `json:"name"`

	// need for subscribe
	Link netlink.Netlink `json:"link"`
	// update or delete
	Event string `json:"event"`
}

func prepareFunc(id string, r Request) (fc *pbtypes.FunctionContext, err error) {
	ft := &pbtypes.FunctionType{
		Namespace: namespace,
		Type:      types.FunctionPath,
	}

	fc = &pbtypes.FunctionContext{
		Id:           id,
		FunctionType: ft,
	}
	fc.Value, err = json.Marshal(r)
	return
}

// genFunction generate function call with object uuid and qdsl
func genFunction(id, query string) (*pbtypes.FunctionContext, error) {
	r := Request{Query: query}
	return prepareFunc(id, r)
}

// genDimmFunction generate function call with object uuid and dimm name
func genDimmFunction(id, query, name string) (*pbtypes.FunctionContext, error) {
	r := Request{Query: query, Name: name}
	return prepareFunc(id, r)
}

// genNetlinkFunction generate function call with object uuid and link object
func genNetlinkFunction(id, query string, link netlink.Netlink, event string) (*pbtypes.FunctionContext, error) {
	r := Request{Query: query, Link: link, Event: event}
	return prepareFunc(id, r)
}
