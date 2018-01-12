package cmd

import (
	"delay-queue/routers"
	"log"
	"net/http"
	"queue/config"
)

//go http request
func (cmd *Cmd) WebRequest(data string) {
	http.HandleFunc("/push", routers.Push)
	http.HandleFunc("/pop", routers.Pop)
	http.HandleFunc("/finish", routers.Delete)
	http.HandleFunc("/delete", routers.Delete)
	http.HandleFunc("/get", routers.Get)

	//log.Printf("listen %s\n", config.Setting.BindAddress)
	err := http.ListenAndServe(config.BindAdress, nil)
	log.Fatalln(err)
}
