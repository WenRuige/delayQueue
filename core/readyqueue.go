package core

//预备队列

//放到预备队列中
func pushToReadyQueue(topic string, id int) error {
	_, err := exec("RPUSH", topic, id)
	return err
}

func GetReadyQueue(topic string) error {
return nil
}
