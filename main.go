package main

import (
	"lemon-robot-dispatcher/core"
	"lemon-robot-dispatcher/handler"
	"lemon-robot-dispatcher/sysinfo"
	"lemon-robot-golang-commons/logger"
	"lemon-robot-golang-commons/utils/lemonrobot"
	"lemon-robot-golang-commons/utils/lruhttp"
)

const executorPath = "/Users/lemonit_cn/Documents/code/lemon-robot/lemon-robot-runner-executor/target/lemon-robot-runner-executor-0.0.1-jar-with-dependencies.jar"
const jrePath = "/Users/lemonit_cn/Downloads/jre-darwin1.8.0_192.jre/Contents/Home/bin/java"

//var (
//	runningPool map[string]structs.Running
//)

func startUp() {
	logger.Info("Start the " + sysinfo.AppName() + " startup process")
	lemonrobot.PrintInfo(sysinfo.AppName(), sysinfo.AppVersion())

	lruhttp.SetBaseUrl(sysinfo.LrConfig().LRServerUrl)
	core.LoginToServer()
	handler.StartPingHandler()

	core.WorkLock.Wait()
}

func main() {
	startUp()
	//vm := otto.New()
	//vm.Run(`
	//	function register(){
	//		return {
	//			hello: function(name){
	//				return {
	//					name: name + '111'
	//				}
	//			}
	//		}
	//	}
	//`)
	//vm.Run(`var obj = register()`)
	//val, _ := vm.Run(`JSON.stringify(obj.hello('liuri11112'))`)
	//fmt.Println(val)
	//fmt.Println(runtime.GOARCH)
	//fmt.Println(runtime.GOOS)
	//fmt.Println("=================")
	//runningPool = make(map[string] structs.Running)
	//runningKey := strconv.FormatInt(time.Now().Unix(), 10)
	//go executeInstance(runningKey)
	//fmt.Println("OVER")
	//time.Sleep(20  * time.Second)
	//runningPool[runningKey].CancelFunc()
	//fmt.Println("OVER222")
	////runningPool[runningKey].CancelFunc()
	////time.Sleep(2 * time.Second)
	//logs := runningPool[runningKey].Logs
	//for i := logs.Front(); i != nil; i = i.Next() {
	//	fmt.Println(logs.Len(), "ja:", i.Value)
	//	i.Next()
	//}
}

//func executeInstance(runningKey string) {
//	fmt.Println("Start a instance running, key: ", runningKey)
//	ctx, cancel := context.WithCancel(context.Background())
//	cmd := exec.CommandContext(ctx, jrePath, "-jar", executorPath)
//	stdout, _ := cmd.StdoutPipe()
//	stdoutReader := bufio.NewReader(stdout)
//	runningPool[runningKey] = structs.Running{Logs: list.New(), CancelFunc: cancel}
//	cmd.Start()
//	for {
//		line, _, err := stdoutReader.ReadLine()
//		if err != nil {
//			break
//		}
//		logContent := string(line)
//		runningPool[runningKey].Logs.PushBack(logContent)
//		fmt.Println("[STD", runningKey, "]", logContent)
//	}
//	fmt.Println("Instance running complete!")
//}
