package cmd

import (
	"log"
	"net/http"
	"queue/config"
	"queue/router"
)

//go http request
func (cmd *Cmd) WebRequest(data string) {
	http.HandleFunc("/push", router.Push)
	http.HandleFunc("/consume",router.Consume)
	log.Printf("listen %s\n", config.BindAddress)
	err := http.ListenAndServe(config.BindAddress, nil)
	log.Fatalln(err)
}
