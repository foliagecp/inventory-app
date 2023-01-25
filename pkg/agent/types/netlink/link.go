// Copyright 2022 Listware

package netlink

import (
	"fmt"
	"strings"

	"github.com/vishvananda/netlink"
)

// Addr ...
type Addr struct {
	IP      string `json:"ip,omitempty"`
	Mask    string `json:"mask,omitempty"`
	Network string `json:"network,omitempty"`
}

// Link profile
type Netlink struct {
	PortID   int    `json:"port-id"`
	PortType string `json:"port-type"`
	Type     string `json:"type,omitempty"`
	MTU      int    `json:"mtu,omitempty"`
	Name     string `json:"name,omitempty"`
	MacAddr  string `json:"mac,omitempty"`
	State    string `json:"state,omitempaty"`
	Alias    string `json:"alias,omitempty"`
	VlanID   int    `json:"vlan-id,omitempty"`

	Numa  int    `json:"numa,omitempty"`
	Addrs []Addr `json:"addrs,omitempty"`
}

func (l *Netlink) LinkName() string {
	if l.VlanID != 0 {
		return fmt.Sprintf("%s%d", l.PortType, l.VlanID)
	}
	return fmt.Sprintf("%s%d", l.PortType, l.PortID)
}

func addrs(link netlink.Link) (res []Addr) {
	// Get all IPv4 addrs
	addrs, _ := netlink.AddrList(link, 2)
	for _, a := range addrs {
		var network string
		ipnet := strings.Split(a.IPNet.String(), "/")
		if len(ipnet) == 2 {
			network = fmt.Sprintf("%s/%s", a.IP.Mask(a.Mask).String(), ipnet[1])

		}
		addr := Addr{
			IP:      a.IP.String(),
			Mask:    ipv4MaskString(a.Mask),
			Network: network,
		}
		res = append(res, addr)
	}
	return res
}

// Net mask in dot decimal notation
func ipv4MaskString(m []byte) string {
	if len(m) != 4 {
		return ""
	}
	return fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3])
}
