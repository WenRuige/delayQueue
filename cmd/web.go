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
	log.Printf("listen %s\n", config.BindAdress)
	err := http.ListenAndServe(config.BindAdress, nil)
	log.Fatalln(err)
}
