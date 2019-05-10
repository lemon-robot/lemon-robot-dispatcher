package main

import (
	"fmt"
	"lemon-robot-dispatcher/core"
	"lemon-robot-dispatcher/sysinfo"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lemonrobot"
	"lemon-robot-golang-commons/utils/lruhttp"
)

func startUp() {
	logger.Info("Start the " + sysinfo.AppName() + " startup process")
	lemonrobot.PrintInfo(sysinfo.AppName(), sysinfo.AppVersion())
	lru_http.SetBaseUrl(fmt.Sprintf("http://%v:%v", sysinfo.LrDispatcherConfig().LRServerHost, sysinfo.LrDispatcherConfig().LRServerPort))
	core.LoginToServer()
}

func main() {
	startUp()
}
