package core

import (
	"queue/model"
	"encoding/json"
	"log"
)

// put job into job pool
func putJob(jobId int, job model.Job) error {
	res, err := json.Marshal(job)
	if err != nil {
		log.Printf("Json 解析失败|%s", err.Error())
		return err
	}
	_, err = exec("set", jobId, res);
	if err != nil {
		log.Printf("redis error %s", err.Error())
		return err
	}
	return nil
}

//获取当前Job
func getJob(jobId int) (*model.Job, error) {
	value, err := exec("get", jobId)
	//判断是否error
	if err != nil {
		return nil, err
	}
	//判断是否有数据
	if value == nil {
		return nil, nil
	}
	byteValue := value.([]byte)
	job := &model.Job{}
	//解析json
	err = json.Unmarshal(byteValue, job)
	if err != nil {
		return nil, err
	}
	return job, err
}
