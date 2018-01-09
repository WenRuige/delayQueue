package core

//篮子



//将job id 放入篮子中()
func pushBucket(key string, delayTime int, jobId int) error {
	_, err := exec("ZADD", key, delayTime, jobId)
	return err
}
