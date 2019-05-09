package sysinfo

import (
	"lemon-robot-dispatcher/model"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lruio"
	"os"
	"path/filepath"
)

var lrDispatcherConfigObj = &model.LrDispatcherConfig{}

const configFileName = "lemon.robot.json"

func configFilePath() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return filepath.Join(dir, configFileName)
}

func checkConfigExisted() bool {
	return lruio.PathExists(configFilePath())
}

func LrDispatcherConfig() *model.LrDispatcherConfig {
	if checkConfigExisted() {
		err := lruio.JsonToStruct(configFilePath(), &lrDispatcherConfigObj)
		if err != nil {
			logger.Error("An error occurred while parsing the configuration ["+configFileName+"], please check your config file.", err)
			os.Exit(1)
			return nil
		}
		return lrDispatcherConfigObj
	}
	logger.Warn("Configuration file  [" + configFileName + "] not found, please check your config file")
	return nil
}
