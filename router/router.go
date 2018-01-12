package router

import (
	"queue/model"
	"encoding/json"
	"log"
	"net/http"
	//"queue/core"
)

//主动push
func Push(resp http.ResponseWriter, req *http.Request) {
	//若是get方法
	if req.Method == "GET" {
		req.ParseForm()
		res := req.Form["data"][0]
		println(res)
	}

	//core.Push()
	resp.Write(generateSuccessBody("success", ""))
}
func generateResponseBody(errno int, msg string, data interface{}) ([]byte) {
	body := &model.ResponseBody{
	}
	body.Code = errno
	body.Message = msg
	body.Data = data
	bytes, err := json.Marshal(body)
	if (err != nil) {
		log.Printf("生成数据失败 %s", err.Error())
		return []byte(`{"code":"1", "message": "生成响应body异常", "data":[]}`)
	}
	return bytes
}

func generateSuccessBody(msg string, data interface{}) []byte {
	return generateResponseBody(model.Success, msg, data)
}

func generateFailureBody(msg string) []byte {
	return generateResponseBody(model.Failure, msg, nil)
}
