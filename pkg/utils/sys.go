// Copyright 2022 Listware

package utils

import (
	"math/rand"
	"os"
	"time"

	"git.fg-tech.ru/listware/inventory-app/pkg/utils/ipmi"
	"github.com/shirou/gopsutil/host"
	"github.com/sirupsen/logrus"
	"github.com/zcalusic/sysinfo"
)

var (
	forceUuidEnvName = "FORCE_BASEBOARD_UUID"

	srcUuidEnvName = "SRC_BASEBOARD_UUID"

	log = logrus.New()
)

// Sys info
var Sys *sys

type sys struct {
	sysinfo.SysInfo
	*host.InfoStat

	BootTime time.Time
}

func Init() {
	Sys = newSys()
}

func newSys() *sys {
	s := &sys{
		InfoStat: getHostInfo(),
		BootTime: bootTime(),
	}

	s.GetSysInfo()

	if v, ok := os.LookupEnv(forceUuidEnvName); ok {
		log.Infof("overidden uuid: %s", v)
		s.InfoStat.HostID = v
		return s
	}

	if v, ok := os.LookupEnv(srcUuidEnvName); ok {
		if v == "dmi" {
			log.Infof("uuid from dmi: %s", s.InfoStat.HostID)
			return s
		}
	}

	u, err := getUUIDFromIPMI()
	if u != "" {
		log.Infof("uuid from ipmi: %s", s.InfoStat.HostID)
		s.InfoStat.HostID = u
	} else {
		log.Warnf("can't read uuid from ipmi, using dmi: %v", err)
	}

	return s
}

func getUUIDFromIPMI() (string, error) {
	drvWasLoaded := ipmi.Driver.IsLoaded()
	if !drvWasLoaded {
		if err := ipmi.Driver.Load(); err != nil {
			return "", err
		}
	}
	uuid, err := ipmi.Tool.MC().GUID()
	if err != nil {
		return "", err
	}
	if !drvWasLoaded {
		err = ipmi.Driver.Unload()
	}
	return uuid, err
}

func getHostInfo() *host.InfoStat {
	hostInfo, err := host.Info()
	if err != nil {
		return nil
	}
	return hostInfo
}

// Random func
func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func bootTime() time.Time {
	if bt, err := host.BootTime(); err == nil {
		return time.Unix(int64(bt), 0)
	}
	return time.Now()
}
