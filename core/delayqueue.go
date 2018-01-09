package core

import (
	"queue/model"
	"errors"
	"log"
	"math/rand"
	"time"
	"queue/config"
)

func Init() {
	InitTimer()
}

//timer 定时器,每两秒去请求一下handle方法
func InitTimer() {
	timer := time.NewTimer(time.Second * 2)
	for {
		select {
		case <-timer.C:
			println("2s timer")
			handler()
			timer.Reset(time.Second * 2)
		}
	}
}

//创建一个timer 轮询bucket
//查询bucket中最近的一个bucket _ job _ id
//消费该id
//回调该回调函数
//回调成功后消除,若回调响应失败则进行重试
//重试次数&重试间隔
func handler() {
	//处理器
	bucket, err := getDataFromBucket(config.DefaultBucketName)
	if err != nil {
		log.Printf("扫描bucket为空%s", err.Error())
		return
	}
	//如果篮子为空
	if bucket == nil {
		return
	}
	//@todo 这个时间需要精准
	//if bucket.Timestamp > 1000 {
	//
	//}
	//获取Job信息
	println(bucket.Timestamp)
}

//push数据到redis中
func Push(job model.Job) (error) {
	//data,err:=exec("get","samplekey");
	//valueBytes := data.([]byte)
	//str := string(valueBytes[:])
	//if err!= nil{}
	//@todo 对于这个id,应该是使用发号器来进行实现
	job.Id = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	job.Topic = "TEST_TOPIC"
	job.Delay = int(time.Now().Unix()) + 30*24*60
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
