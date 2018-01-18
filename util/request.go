package util

//请求处理

import "net/http"

type Request struct {
	res http.ResponseWriter
	req *http.Request
}
