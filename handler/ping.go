package handler

import (
	"lemon-robot-dispatcher/core"
	"lemon-robot-dispatcher/sysinfo"
	"time"
)

func StartPingHandler() {
	core.WorkLock.Add(1)
	ticker := time.NewTicker(time.Second * time.Duration(sysinfo.LrConfig().TimerPingInterval))
	go func() {
		for range ticker.C {
			core.Ping()
		}
	}()
}
