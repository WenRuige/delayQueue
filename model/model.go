package model

//状态码
const (
	Success = 0
	Failure = -1
)

//返回体信息
type ResponseBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

//定时任务Job
type Job struct {
	Topic      string `json:"topic"`
	Id         int    `json:"id"`
	Delay      int    `json:"delay"`
	Body       string `json:"body"`
	Callback   string `json:"callback"`
	RetryTimes int    `json:"retry_times"`
}


//篮子
type BucketItem struct {
	Timestamp int `json:"timestamp"`
	Jobid    int `json:"jobid"`
}