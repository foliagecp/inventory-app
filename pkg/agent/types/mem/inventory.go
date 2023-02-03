// Copyright 2023 NJWS Inc.
// Copyright 2022 Listware

package mem

import (
	"github.com/sirupsen/logrus"
	"github.com/yumaojun03/dmidecode"
	"github.com/yumaojun03/dmidecode/parser/memory"
)

// Channels for get updates
var (
	log = logrus.New()
)

func deviceToProfile(device *memory.MemoryDevice) (dev MemoryDevice) {
	dev = MemoryDevice{
		Size:         device.Size,
		FormFactor:   device.FormFactor.String(),
		Device:       device.DeviceLocator,
		Bank:         device.BankLocator,
		Type:         device.Type.String(),
		Speed:        device.Speed,
		Manufacturer: device.Manufacturer,
		Serial:       device.SerialNumber,
		Part:         device.PartNumber,
	}

	return
}

// Inventory implement inventory interface
func Inventory() (devs map[string]MemoryDevice, err error) {
	log.Info("inventory mem")

	devs = make(map[string]MemoryDevice)

	dmi, err := dmidecode.New()
	if err != nil {
		log.Warn(err)
		err = nil
		return
	}
	devices, err := dmi.MemoryDevice()
	if err != nil {
		log.Warn(err)
		err = nil
		return
	}

	for _, device := range devices {
		dev := deviceToProfile(device)
		devs[dev.Name()] = dev
	}
	return
}
