package core

import (
	"queue/model"
	"errors"
	"log"
	"math/rand"
	"time"
	"queue/config"
)

//
func Init() {
	Handler()
}

func getBucket() {

}

func Handler() {
	err := getDataFromBucket(config.DefaultBucketName)
	if err != nil {

	}
}

//创建一个timer 轮询bucket
//查询bucket中最近的一个bucket _ job _ id
//消费该id
//回调该回调函数
//回调成功后消除,若回调响应失败则进行重试
//重试次数&重试间隔

func Push(job model.Job) (error) {
	//data,err:=exec("get","samplekey");
	//valueBytes := data.([]byte)
	//str := string(valueBytes[:])
	//if err!= nil{}
	//@todo 对于这个id,应该是使用发号器来进行实现
	job.Id = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	job.Topic = "TEST_TOPIC"
	job.Delay = 30000
	job.Body = ""
	job.Callback = "http://www.baidu.com"

	if job.Id == 0 || job.Topic == "" || job.Delay == 0 || job.Callback == "" {
		return errors.New("有部分数据为空")
	}
	err := putJob(job.Id, job)
	if err != nil {
		log.Printf("放入job poll error |%s", err.Error())
		return err
	}
	//默认的Bucket
	err = pushBucket(config.DefaultBucketName, job.Delay, job.Id)
	if err != nil {
		log.Printf("放入篮子error|%s", err.Error())
		return err
	}
	return nil
}
