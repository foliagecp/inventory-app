// Copyright 2022 Listware

package netlink

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/vishvananda/netlink"
)

var (
	log = logrus.New()
)

func Subscribe(ctx context.Context) (updateChan, deleteChan chan Netlink) {
	updateChan = make(chan Netlink, 1000)
	deleteChan = make(chan Netlink, 100)

	go func() {
		linkUpdateChan := make(chan netlink.LinkUpdate, 100)
		if err := netlink.LinkSubscribe(linkUpdateChan, ctx.Done()); err != nil {
			log.Warn(err.Error())
			return
		}
		for {
			select {
			case m := <-linkUpdateChan:

				switch m.Header.Type {
				case syscall.RTM_NEWLINK:
					if filterLinks(m.Link) {
						log.Debugf("Handle link type: %s", m.Link.Type())
						updateChan <- netlinkToLink(m.Link)
					}
				case syscall.RTM_DELLINK:
					if filterLinks(m.Link) {
						log.Debugf("Handle link type: %s", m.Link.Type())
						deleteChan <- netlinkToLink(m.Link)
					}
				}

			case <-ctx.Done():
				return
			}
		}

	}()

	go func() {
		addrUpdateChan := make(chan netlink.AddrUpdate, 100)
		if err := netlink.AddrSubscribe(addrUpdateChan, ctx.Done()); err != nil {
			log.Warn(err.Error())
		}
		for {
			select {
			case m := <-addrUpdateChan:
				if link, err := netlink.LinkByIndex(m.LinkIndex); err == nil {
					if filterLinks(link) {
						log.Debugf("Handle address changes: %s", m.LinkAddress.String())
						updateChan <- netlinkToLink(link)
					}
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return
}

func New() (initLinks map[string]Netlink, err error) {
	log.Info("inventory links")

	initLinks = make(map[string]Netlink)

	links, err := netlink.LinkList()
	if err != nil {
		return
	}

	for _, link := range links {
		if filterLinks(link) {
			l := netlinkToLink(link)
			initLinks[l.LinkName()] = l
		}
	}

	return
}

// spaghetti filter
func filterLinks(link netlink.Link) bool {
	if link.Type() == "veth" || link.Type() == "bridge" {
		return false
	}
	attrs := link.Attrs()
	if (attrs.Flags & net.FlagLoopback) != 0 {
		return false
	}
	if attrs.MasterIndex != 0 {
		return false
	}
	return true
}

func netlinkToLink(link netlink.Link) Netlink {
	attrs := link.Attrs()
	l := Netlink{
		PortType: "os-def",
		PortID:   attrs.Index - 1,
		Type:     link.Type(),
		MTU:      attrs.MTU,
		Name:     attrs.Name,
		MacAddr:  attrs.HardwareAddr.String(),
		Alias:    attrs.Alias,
		State:    attrs.OperState.String(),
	}

	if attrs.ParentIndex != 0 {
		l.PortID = attrs.ParentIndex
	}
	switch attrs.EncapType {
	case "infiniband":
		l.PortType = "os-ib"
		l.Type = attrs.EncapType
	case "ether":
		l.PortType = "os-eth"
	}
	if link.Type() == "vlan" {
		if vlan, ok := link.(*netlink.Vlan); ok {
			l.VlanID = vlan.VlanId
			l.PortType = "os-vlan"
			l.PortID = (attrs.ParentIndex - 1)
		}
	}
	l.Addrs = addrs(link)

	data, err := os.ReadFile(fmt.Sprintf("/sys/class/net/%s/device/numa_node", l.Name))
	if err == nil {
		l.Numa, err = strconv.Atoi(strings.TrimSuffix(string(data), "\n"))
		if err != nil {
			log.Warn(err)
		}
	}

	return l
}
