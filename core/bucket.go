package core

import (
	"strconv"
	"queue/model"
	"queue/config"
)

//将job id 放入篮子中
//实际上bucket里面的内容是有序切不重复的
func pushBucket(key string, delayTime int, jobId int) error {
	_, err := exec("ZADD", key, delayTime, jobId)
	return err
}

//生成对应篮子的序号
//@todo:有点问题,每次都是返回bucket1
func generateBucketName() <-chan string {
	c := make(chan string)
	go func() {
		i := 1
		for {
			c <- config.DefaultBucketName + strconv.Itoa(i)
			if i >= 10 {
				i = 1
			} else {
				i++
			}
		}
	}()
	return c
}

//从bucket中获取数据()
func getDataFromBucket(key string) (*model.BucketItem, error) {
	res, err := exec("ZRANGE", key, 0, 0, "WITHSCORES")
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	var valueBytes []interface{}
	valueBytes = res.([]interface{})
	//change byte to string
	timestampStr := string(valueBytes[1].([]byte))
	jobIdStr := string(valueBytes[0].([]byte))
	//add a bucket
	item := &model.BucketItem{}
	item.Timestamp, _ = strconv.Atoi(timestampStr)
	item.Jobid, _ = strconv.Atoi(jobIdStr)
	return item, nil
}

//从篮子里面删除该数据
func removeFromBucket(bucketName string, id int) error {
	_, err := exec("ZREM", bucketName, id)
	return err
}
