// Copyright 2022 Listware

package cpu

import (
	"git.fg-tech.ru/listware/inventory-app/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/zcalusic/sysinfo"
)

// CPU profile
type CPU struct {
	sysinfo.CPU
}

// Inventory \\
func New() (c CPU, err error) {
	logrus.Infof("cpu: %s", utils.Sys.CPU.Model)

	c = CPU{CPU: utils.Sys.CPU}

	return
}
