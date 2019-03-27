package core

import (
	"encoding/json"
	"fmt"
	"lemon-robot-dispatcher/define/define_storage_key"
	"lemon-robot-dispatcher/subutils"
	"lemon-robot-dispatcher/sysinfo"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lruhttp"
	"os"
)

func LoginToServer() {
	logger.Info("Logging in to the server as: " + sysinfo.LrConfig().LRServerUserNumber)
	responseText, err := lruhttp.RequestJson("POST", "/user/login", map[string]string{
		"number":   sysinfo.LrConfig().LRServerUserNumber,
		"password": sysinfo.LrConfig().LRServerUserPassword,
	}, map[string]string{})
	if err != nil {
		logger.Error("Cannot login to server", err)
		os.Exit(-1)
	}
	var responseMap map[string]interface{}
	if err := json.Unmarshal([]byte(responseText), &responseMap); err != nil {
		logger.Error("Cannot read server login response:"+responseText, err)
		os.Exit(1)
	}
	if responseMap["success"] != true {
		logger.Error(fmt.Sprintf("Login to server failed, server say: %s", subutils.TranslateErrCode(responseMap["code"].(string))), nil)
		os.Exit(1)
	}
	token := responseMap["data"].(string)
	StoragePut(define_storage_key.LOGIN_TOKEN, token)
	lruhttp.AppendCommonHeader(map[string]string{"Authorization": "Bearer " + token})
	logger.Info("Login successful, token: " + token)
	Ping()
}

func Ping() {
	logger.Info("Ping to server")
	responseText, err := lruhttp.RequestJson("POST", "/dispatcher/ping", nil, nil)
	if err != nil {
		logger.Error("Cannot login to server", err)
		os.Exit(-1)
	}
	fmt.Println(responseText)
}
