package consumer

import (
	"queue/core"
	"log"
)

//消费者客户端

func Init() {
	topic := "TEST_TOPIC"
	go consume(topic)
}

//开启消费进程
func consume(topic string) {
	for {
		err := core.GetReadyQueue(topic)
		if err != nil {
			log.Printf(" 消费Error%s", err.Error())
		}
	}
}
