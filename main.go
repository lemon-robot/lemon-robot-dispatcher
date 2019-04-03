package main

import (
	"fmt"
	"lemon-robot-dispatcher/core"
	"lemon-robot-dispatcher/sysinfo"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lemonrobot"
	"lemon-robot-golang-commons/utils/lruhttp"
	"lemon-robot-golang-commons/utils/machine"
	"os"
)

func startUp() {
	logger.Info("Start the " + sysinfo.AppName() + " startup process")
	lemonrobot.PrintInfo(sysinfo.AppName(), sysinfo.AppVersion())
	lruhttp.SetBaseUrl(fmt.Sprintf("http://%v:%v", sysinfo.LrConfig().LRServerHost, sysinfo.LrConfig().LRServerPort))
	core.LoginToServer()
}

func main() {
	machineCode, mcErr := lrumachine.CalculateMachineCodeByMAC()
	if mcErr != nil {
		logger.Error("The system could not register because the machine code could not be generated from the MAC address.", mcErr)
		os.Exit(1)
	}
	logger.Info("The machine code has been calculated：" + machineCode)
	//startUp()
}
