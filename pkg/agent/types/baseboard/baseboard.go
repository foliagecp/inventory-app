// Copyright 2022 Listware

package baseboard

import (
	"git.fg-tech.ru/listware/inventory-app/pkg/utils"
	"github.com/sirupsen/logrus"
)

// Baseboard information.
type Baseboard struct {
	UUID    string `json:"uuid,omitemtpy"`
	Name    string `json:"name,omitempty"`
	Vendor  string `json:"vendor,omitempty"`
	Version string `json:"version,omitempty"`
	Serial  string `json:"serial,omitempty"`
}

// Inventory interface
func New() (b Baseboard, err error) {
	utils.Init()

	logrus.Infof("baseboard: %s", utils.Sys.Board.Name)

	b = Baseboard{
		UUID:    utils.Sys.HostID,
		Name:    utils.Sys.Board.Name,
		Vendor:  utils.Sys.Board.Vendor,
		Version: utils.Sys.Board.Version,
		Serial:  utils.Sys.Board.Serial,
	}

	return
}
