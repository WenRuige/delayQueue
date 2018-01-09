package core

import (
	"queue/model"
	"encoding/json"
	"errors"
)

func Init() {

}

// put job into job pool
func putJob(jobId int, job model.Job) error {
	res, err := json.Marshal(job)
	if err != nil {
		errors.New("Json解析失败")
	}
	_, err = exec("set", jobId, res);
	if err != nil {
		errors.New("put job into redis error")
	}
	return nil
}
