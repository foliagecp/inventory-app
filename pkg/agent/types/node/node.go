// Copyright 2022 Listware

package node

import (
	"fmt"
	"strings"

	"git.fg-tech.ru/listware/inventory-app/pkg/utils"
	"github.com/sirupsen/logrus"
)

// Node profile
type Node struct {
	Hostname string `json:"hostname,omitempty"`
	Domain   string `json:"domain,omitempty"`
	Model    string `json:"model,omitempty"`
}

func Name() (string, error) {
	fqdn := strings.SplitN(utils.Sys.Hostname, ".", 2)
	if len(fqdn) == 0 {
		return "", fmt.Errorf("bad fqdn")
	}
	return fqdn[0], nil
}

// Inventory interface
func New() (n Node, err error) {
	logrus.Infof("node: %s", utils.Sys.Hostname)

	fqdn := strings.SplitN(utils.Sys.Hostname, ".", 2)
	if len(fqdn) == 0 {
		err = fmt.Errorf("bad fqdn")
		return
	}
	var dn string

	if len(fqdn) > 1 {
		dn = fqdn[1]
	}

	n = Node{
		Hostname: fqdn[0],
		Domain:   dn,
		Model:    utils.Sys.Board.Name,
	}

	return
}
