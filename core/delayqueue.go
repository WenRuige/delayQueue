package core

import (
	"queue/model"
	//"errors"
	"log"
	"time"
	"queue/config"
	"errors"
)

func FlushDb() {
	exec("flushdb")
}

func Init() {
	RedisPool = initRedisPool()
	InitTimer()

}
func InitTimer() {
	//一个三秒的定时器
	//println("hello")
	//t := time.NewTicker(1 * time.Second)
	//for {
	//	select {
	//	case <-t.C:
	//		handler(config.DefaultBucketName)
	//	}
	//}
	handler(config.DefaultBucketName)
}

func handler(bucketName string) {
	//处理器
	bucketItem, err := getDataFromBucket(bucketName)
	if err != nil {
		log.Printf("扫描bucket为空%s", err.Error())
		return
	}
	//因为返回的Err是Nil
	if bucketItem == nil {
		return
	}

	if bucketItem.Timestamp > int(time.Now().Unix()) {
		return
	}
	//获取Job信息
	jobObj, err := getJob(bucketItem.Jobid)
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
	err = removeFromBucket(bucketName, jobObj.Id)
	if err != nil {
		log.Printf("删除bucket失败|%s|", err.Error())
	}

}

//@todo 对于这个id,应该是使用发号器来进行实现
//===========================     test case         =========================
//job.Id = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)
//job.Topic = "TEST_TOPIC"
//job.Delay = int(time.Now().Unix()) + 30*24*60
//job.Body = "hello world"
//job.Callback = "http://www.baidu.com"
//===========================     test case end      ==========================
//push数据到redis中
func Push(job model.Job) (error) {
	if job.Id == 0 || job.Topic == "" || job.Delay == 0 || job.Callback == "" {
		return errors.New("有部分数据为空")
	}
	println(job.Id)
	err := putJob(job.Id, job)
	if err != nil {
		log.Printf("放入job poll error |%s", err.Error())
		return err
	}
	////默认的Bucket,此处建议由多个Bucket来组成
	err = pushBucket(config.DefaultBucketName, job.Delay, job.Id)
	if err != nil {
		log.Printf("放入篮子error|%s", err.Error())
		return err
	}
	return nil
}
