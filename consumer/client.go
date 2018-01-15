package consumer

//消费者客户端



func Init(){
	go consume()
}

//开启消费进程
func consume(){
	println("hello world_________________________")
}
