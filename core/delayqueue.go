package core

import (
	"queue/model"
	//"errors"

	"errors"
)

func Push(job model.Job) (error) {
	//data,err:=exec("get","samplekey");
	//valueBytes := data.([]byte)
	//str := string(valueBytes[:])
	//if err!= nil{}
	//println(str)
	job.Id = 1
	job.Topic = "TEST_TOPIC"
	job.Delay = 3
	job.Body = ""
	job.Callback = "http://www.baidu.com"

	if job.Id == 0 || job.Topic == "" || job.Delay == 0 || job.Callback == "" {
		return errors.New("有部分数据为空")
	}
	err := putJob(job.Id, job)
	if err != nil {
		errors.New("put job in redis error")
	}
	pushBucket(job.Delay,job.Id)
	return nil
}
