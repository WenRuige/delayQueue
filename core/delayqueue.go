package core

import (
	"queue/model"
	"errors"
	"log"
	"math/rand"
	"time"
	"queue/config"
)

var (
	// 每个定时器对应一个bucket
	timers         []*time.Ticker
	bucketNameChan <- chan string
)

func Init() {
	InitTimer()
	bucketNameChan = generateBucketName()
}





//建立一个Timer
func InitTimer() {
	timers = make([]*time.Ticker, 10)
	for i := 0; i < 10; i++ {
		timers[i] = time.NewTicker(1 * time.Second)
		go waitTicker(timers[i])
	}
	//需要阻塞一下timer来防止协程未执行完
	time.Sleep(time.Second * 10)
}

func waitTicker(timer *time.Ticker) {
	for {
		select {
		case <-timer.C:
			println("2s timer")
			//handler()
		}
	}
}

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
	if bucket.Timestamp > int(time.Now().Unix()) {
		return
	}
	//获取Job信息
	jobObj, err := getJob(bucket.Jobid)
	if err != nil {
		log.Printf("%s |job元信息为空", err.Error())
	}
	//check job delay和当前时间相比较
	if jobObj.Delay > int(time.Now().Unix()) {
		//删除篮子内的时间
		log.Printf("当前Job未到延时时间")
	}

	err = pushToReadyQueue(jobObj.Topic, jobObj.Id)
	if err != nil {
		log.Printf("放入ready queue error|%s|", err.Error())
	}
	err = removeFromBucket(config.DefaultBucketName, jobObj.Id)
	if err != nil {
		log.Printf("删除bucket失败|%s|", err.Error())
	}
	println("success")

}

//push数据到redis中
func Push(job model.Job) (error) {
	//@todo 对于这个id,应该是使用发号器来进行实现
	//===========================     test case         =========================
	job.Id = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
	job.Topic = "TEST_TOPIC"
	job.Delay = int(time.Now().Unix()) + 30*24*60
	job.Body = "hello world"
	job.Callback = "http://www.baidu.com"
	//===========================     test case end      ==========================
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
