package util

import (
	"log"
	"go/build"
	"flag"
	"os"
)

var (
	Log *log.Logger
)

//创建Log
func init() {
	var logPath = build.Default.GOPATH + "/src/queue/log/info.log"
	flag.Parse()
	flag.Parse()
	var file, err = os.Create(logPath)
	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	Log.Println("LogFile : " + logPath)
}
