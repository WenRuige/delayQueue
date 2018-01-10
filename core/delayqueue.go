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
//消费该id,放入预备队列中

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
	println("hello world")
	//@todo 这个时间需要精准,如果延迟时间大于当前时间,表示延迟时间未到
	if bucket.Timestamp > int(time.Now().Unix()) {
		return
	}
	//获取Job信息
	jobObj, err := getJob(bucket.Jobid)
	if err != nil {
		log.Printf("%s |job元信息为空", err.Error())
	}
	println(jobObj.Delay)
	//check job delay和当前时间相比较
	if jobObj.Delay > int(time.Now().Unix()){
		//删除篮子内的时间
		log.Printf("当前Job未到延时时间")
	}

}

//push数据到redis中
func Push(job model.Job) (error) {
	//@todo 对于这个id,应该是使用发号器来进行实现
	job.Id = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	job.Topic = "TEST_TOPIC"
	job.Delay = int(time.Now().Unix()) + 30*24*60
	job.Body = "hello world"
	job.Callback = "http://www.baidu.com"

	if job.Id == 0 || job.Topic == "" || job.Delay == 0 || job.Callback == "" {
		return errors.New("有部分数据为空")
	}
	err := putJob(job.Id, job)
	if err != nil {
		log.Printf("放入job poll error |%s", err.Error())
		return err
	}
	//默认的Bucket,此处建议由多个Bucket来组成
	err = pushBucket(config.DefaultBucketName, job.Delay, job.Id)
	if err != nil {
		log.Printf("放入篮子error|%s", err.Error())
		return err
	}
	return nil
}
