package core

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"lemon-robot-dispatcher/define/define_storage_key"
	"lemon-robot-dispatcher/subutils"
	"lemon-robot-dispatcher/sysinfo"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lru_http"
	"lemon-robot-golang-commons/utils/lru_machine"
	"log"
	"os"
	"runtime"
)

func LoginToServer() {
	logger.Info("Logging in to the server as: " + sysinfo.LrDispatcherConfig().LRServerUserNumber)
	responseText, err := lru_http.GetInstance().RequestJson("POST", "/user/login", map[string]string{
		"number":   sysinfo.LrDispatcherConfig().LRServerUserNumber,
		"password": sysinfo.LrDispatcherConfig().LRServerUserPassword,
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
	lru_http.GetInstance().AppendCommonHeader(map[string]string{"Authorization": "Bearer " + token})
	logger.Info("Login successful, token: " + token)
	ListenTheServer(token)
}

func ListenTheServer(token string) {
	dialer := websocket.Dialer{}
	logger.Info("The machine code has been calculated：" + lru_machine.GetInstance().GetMachineSign())
	conUrl := fmt.Sprintf("ws://%v:%v/ws/%v/%v/%v/%v/%v",
		sysinfo.LrDispatcherConfig().LRServerHost, sysinfo.LrDispatcherConfig().LRServerPort, runtime.GOOS,
		runtime.GOARCH, sysinfo.AppVersion(), lru_machine.GetInstance().GetMachineSign(), token)
	con, _, err := dialer.Dial(conUrl, nil)
	if err != nil {
		logger.Error("Cannot connect to the websocket server", err)
		os.Exit(1)
	}
	logger.Info("Websocket was successfully established")
	for {
		_, message, err := con.ReadMessage()
		if err != nil {
			logger.Error("Errors occurred while reading cancelled messages from websocket", err)
		}
		log.Printf("Receive messages from websocket: %s", message)
	}
}
