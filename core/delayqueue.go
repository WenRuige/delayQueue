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



	//data,err:=exec("set","samplekey111111","123");
	//if err != nil{
	//
	//}
	//
	//println(data)

	job.Id = "1"
	job.Topic = "TEST_TOPIC"
	job.Delay = "3"
	job.Body = ""
	job.Callback = "http://www.baidu.com"

	if job.Id == "" || job.Topic == "" || job.Delay == "" || job.Callback == "" {
		return errors.New("有部分数据为空")
	}


	putJob(job.Id,job)
	return nil
}
