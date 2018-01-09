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
