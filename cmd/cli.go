package cmd

import (
	"queue/core"
	//"queue/model"
	//"queue/util"
	//"flag"
	//"fmt"
	//"os"
	//	"runtime"
)

type Cmd struct {
}

func (cmd *Cmd) Run() {
	//获取cli 输入的数据

	core.FlushDb()

	getParams()

	//core.Push(model.Job{})
	//开启消费者的守护进程,进行消费
	//core.Init()

	cmd.WebRequest(getParams())
	//util.Log.Printf("Server v%s pid=%d started with processes: %d", "1.00", os.Getpid(), runtime.GOMAXPROCS(runtime.NumCPU()))
}

func getParams() string {
	// - name 默认值 使用方式
	//data := flag.String("data", "不能为空呀老铁", "数据详情")
	//flag.Parse()
	//if * data != ""{
	//	println("hw")
	//	core.Push(model.Job{})
	//}
	//fmt.Println(* data)
	return ""
}
