package core

import (
	"errors"
	"strconv"
	"net/url"
	"io/ioutil"
	"fmt"
	"net/http"
	"queue/model"
	"github.com/gin-gonic/gin/json"
)

//预备队列

//放到预备队列中
func pushToReadyQueue(topic string, id int) error {
	_, err := exec("RPUSH", topic, id)
	return err
}

//从队列中获取一个元素
func GetReadyQueue(topic string) (error) {
	res, err := exec("LPOP", topic)
	if err != nil {
		errors.New("LPOP ERROR")
	}
	if res == nil {
		return nil
	}
	byteValue := res.([]byte)
	tmp := string(byteValue[0])
	result, err := strconv.Atoi(tmp)
	if err != nil {
		errors.New("转换失败")
	}
	job, err := getJob(result)
	if err != nil {

	}
	httpPost(*job)

	//先回调该接口

	return nil
}

//回调方法
func httpPost(job model.Job) {
	println(job.Callback)
	tmp, err := json.Marshal(job)
	if err != nil {
		errors.New("Json 解析失败")
	}
	resp, err := http.PostForm(job.Callback, url.Values{"key": {"Value"}, "data": {string(tmp)}})

	if err != nil {
		// handle error
		errors.New("http post request error")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
