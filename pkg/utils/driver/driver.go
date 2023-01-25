// Copyright 2022 Listware

package driver

import (
	"os"
	"os/exec"
	"path"

	"github.com/sirupsen/logrus"
)

const (
	kmodSysPath = "/sys/module"
	drvsPath    = "/sys/bus/pci/drivers"
)

// Driver - disk driver: nvme, uio, etc..
type Driver string

func (drv Driver) String() string {
	return string(drv)
}

// IsLoaded - check if driver is loaded
func (drv Driver) IsLoaded() bool {
	_, err := os.Stat(path.Join(kmodSysPath, drv.String()))
	return !os.IsNotExist(err)
}

// Load kernel module
func (drv Driver) Load() error {
	if drv.IsLoaded() {
		logrus.WithField("driver", drv.String()).Debug("is already loaded")
		return nil

	}
	return exec.Command("modprobe", drv.String()).Run()
}

// Unload kernel module
func (drv Driver) Unload() error {
	if !drv.IsLoaded() {
		logrus.WithField("driver", drv.String()).Debug("is already unloaded")
		return nil
	}
	return exec.Command("modprobe", "--remove", drv.String()).Run()
}
