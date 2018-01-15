package consumer

import "queue/core"

//消费者客户端

func Init() {
	topic := "TEST_TOPIC"
	go consume(topic)
}

//开启消费进程
func consume(topic string) {
	//println("hello world_________________________")

	for {
		println("fqqqqq")
	}
	core.GetReadyQueue(topic)

}
