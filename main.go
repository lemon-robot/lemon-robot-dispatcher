package main

import (
	"bufio"
	"container/list"
	"context"
	"fmt"
	"lemon-robot-dispatcher/structs"
	"os/exec"
	"runtime"
)

const executorPath = "/Users/lemonit_cn/Documents/code/lemon-robot/lemon-robot-runner-executor/target/lemon-robot-runner-executor-0.0.1-jar-with-dependencies.jar"
const jrePath = "/Users/lemonit_cn/Downloads/jre-darwin1.8.0_192.jre/Contents/Home/bin/java"

var (
	runningPool map[string] structs.Running
)

func main(){
	fmt.Println(runtime.GOARCH)
	fmt.Printlake(map[string] structs.Running)
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
	//}n(runtime.GOOS)
	fmt.Println("=================")
	//runningPool = m
}

func executeInstance(runningKey string) {
	fmt.Println("Start a instance running, key: ", runningKey)
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, jrePath, "-jar", executorPath)
	stdout, _ := cmd.StdoutPipe()
	stdoutReader := bufio.NewReader(stdout)
	runningPool[runningKey] = structs.Running{Logs: list.New(), CancelFunc: cancel}
	cmd.Start()
	for {
		line, _, err := stdoutReader.ReadLine()
		if err != nil {
			break
		}
		logContent := string(line)
		runningPool[runningKey].Logs.PushBack(logContent)
		fmt.Println("[STD", runningKey, "]", logContent)
	}
	fmt.Println("Instance running complete!")
}
