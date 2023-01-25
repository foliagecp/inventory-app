// Copyright 2022 Listware

package utils

import (
	"io"
	"os/exec"
)

type ipmitool struct {
	cmd    *exec.Cmd
	stdin  io.Reader
	stdout io.Writer
}

func newIpmiTool() (it *ipmitool, err error) {
	it = &ipmitool{}
	// it.cmd, err = exec.Command("")
	return
}
