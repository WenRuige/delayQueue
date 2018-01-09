package core

import (
	"strconv"
	"queue/model"
)

//篮子

//将job id 放入篮子中()
func pushBucket(key string, delayTime int, jobId int) error {
	_, err := exec("ZADD", key, delayTime, jobId)
	return err
}

//从bucket中获取数据()
func getDataFromBucket(key string) error {
	res, err := exec("ZRANGE", key, 0, 0, "WITHSCORES")
	if err != nil {
		return err
	}
	if res == nil {
		return nil
	}
	var valueBytes []interface{}
	valueBytes = res.([]interface{})
	//change byte to string
	timestampStr := string(valueBytes[1].([]byte))
	jobIdStr := string(valueBytes[0].([]byte))

	item := &model.BucketItem{}
	println(string(valueBytes[0].([]byte)))

	item.Timestamp, _ = strconv.Atoi(timestampStr)
	item.Jobid, _ = strconv.Atoi(jobIdStr)

	return nil
}
