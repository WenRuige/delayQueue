package util

import (
	"log"
	"os"
)

var (
	Log *log.Logger
)

//创建Log
func Init() {
	logPath := "log.log"
	var file, err = os.Create(logPath)
	if err != nil {
		panic(err)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)

}
