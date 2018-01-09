package cmd

import (
	"queue/core"
	"queue/model"
	"flag"
	"fmt"

)

type Cmd struct {
}

func (cmd *Cmd) Run() {

	//获取cli 输入的数据
	clearParams()
	core.Push(model.Job{})
	//开启消费者的守护进程,进行消费
	core.Init()

}

func clearParams() {
	// - name 默认值 使用方式
	port := flag.String("port", ":8080", "http listen port")
	flag.Parse()
	fmt.Println(* port)
}
