// Copyright 2022 Listware

package bios

import (
	"git.fg-tech.ru/listware/inventory-app/pkg/utils"
	"github.com/sirupsen/logrus"
)

// BIOS information.
type BIOS struct {
	Vendor  string `json:"vendor,omitempty"`
	Version string `json:"version,omitempty"`
	Date    string `json:"date,omitempty"`
}

// Inventory interface
func New() (b BIOS, err error) {
	logrus.Infof("bios: %s", utils.Sys.BIOS.Vendor)

	b = BIOS{
		Vendor:  utils.Sys.BIOS.Vendor,
		Version: utils.Sys.BIOS.Version,
		Date:    utils.Sys.BIOS.Date,
	}

	return
}
