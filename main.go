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
	lruhttp.SetBaseUrl(fmt.Sprintf("http://%v:%v", sysinfo.LrConfig().LRServerHost, sysinfo.LrConfig().LRServerPort))
	core.LoginToServer()
}

func main() {
	startUp()
}
