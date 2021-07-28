package utils

import "time"

type Response struct {
	Code      int         `json:"code"`      // 1：成功 | 0：失败
	Message   string      `json:"message"`   // 成功：success | 失败：报错信息
	Timestamp int         `json:"timestamp"` // 当前时间
	Data      interface{} `json:"data"`      // 返回参数内容主体
}

func NewResponse(code int, message string, data interface{}) *Response {
	res := new(Response)
	res.Code = code
	res.Message = message
	res.Timestamp = int(time.Now().Unix())
	res.Data = data
	return res
}
