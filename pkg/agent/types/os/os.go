// Copyright 2022 Listware

package os

import (
	"git.fg-tech.ru/listware/inventory-app/pkg/utils"
	"github.com/sirupsen/logrus"
)

// OS profile
type OS struct {
	Type     string `json:"type,omitempty"`
	Platform string `json:"platform,omitempty"`
	Family   string `json:"family,omitempty"`
	Version  string `json:"version,omitempty"`
	BootTime int64  `json:"boot-time,omitempty"`
}

// Inventory interface
func New() (o OS, err error) {
	logrus.Infof("os: %s %s", utils.Sys.InfoStat.OS, utils.Sys.Platform)

	o = OS{
		Type:     utils.Sys.InfoStat.OS,
		Platform: utils.Sys.Platform,
		Family:   utils.Sys.PlatformFamily,
		Version:  utils.Sys.PlatformVersion,
		BootTime: utils.Sys.BootTime.Unix(),
	}
	return
}
