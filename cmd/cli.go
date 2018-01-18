package cmd

import (
	"queue/core"
	"queue/consumer"
	"queue/util"
)

type Cmd struct {
}

func (cmd *Cmd) Run() {
	//获取cli 输入的数据

	util.Init()

	//core.FlushDb()
	getParams()
	//开启生产者
	core.Init()
	//开启消费者
	consumer.Init()
	//启动web端
	cmd.WebRequest(getParams())

	//消费

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
