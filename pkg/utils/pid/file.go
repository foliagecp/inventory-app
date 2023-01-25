// Copyright 2022 Listware

package pid

import (
	"fmt"
	"os"
	"strconv"
	"syscall"
)

// File - is a pid-file
type File string

func (f File) String() string {
	return string(f)
}

// Write a pid file, but first make sure it doesn't exist with a running pid.
func (f File) Write() error {
	// Read in the pid file as a slice of bytes.
	if piddata, err := os.ReadFile(f.String()); err == nil {
		// Convert the file contents to an integer.
		if pid, err := strconv.Atoi(string(piddata)); err == nil {
			// Look for the pid in the process list.
			if process, err := os.FindProcess(pid); err == nil {
				// Send the process a signal zero kill.
				if err := process.Signal(syscall.Signal(0)); err == nil {
					// We only get an error if the pid isn't running, or it's not ours.
					return fmt.Errorf("pid already running: %d", pid)
				}
			}
		}
	}
	// If we get here, then the pidfile didn't exist,
	// or the pid in it doesn't belong to the user running this app.
	return os.WriteFile(f.String(), []byte(fmt.Sprintf("%d", os.Getpid())), 0664)
}

// Remove pid file
func (f File) Remove() error {
	return os.Remove(f.String())
}

// PID - process id
func (f File) PID() (pid int, err error) {
	var piddata []byte
	piddata, err = os.ReadFile(f.String())
	if err != nil {
		return
	}
	pid, err = strconv.Atoi(string(piddata))
	return
}
