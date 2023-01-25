// Copyright 2022 Listware

package agent

import (
	"os"
	"os/signal"
	"syscall"
)

func (a *Agent) osSignalCtrl() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGHUP,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
	)
	go func() {
		for {
			select {
			case sig := <-sigChan:
				switch sig {
				case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT:
					log.Infof("Get Stop signal")
					a.cancel()
				}
			case <-a.ctx.Done():
				return
			}
		}
	}()
}
