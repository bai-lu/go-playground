package utils

import (
	"encoding/json"
)

// json 格式输出
type response struct {
	Code    int         `json:"code"`           // 状态码
	Message string      `json:"message"`        // 信息
	Data    interface{} `json:"data,omitempty"` // 数据
}

type JsonResponse struct{}

const SuccessCode = 2000
const TagErrorCode = 2001
const InspectErrorCode = 2002
const FailureCode = 5000
const NotFound = 404
const ServerError = 500

const SuccessContent = "操作成功"
const FailureContent = "操作失败"

func (j *JsonResponse) Success(code int, message string, data interface{}) string {
	return j.Response(code, message, data)
}

func (j *JsonResponse) Failure(code int, message string, err ...error) string {
	return j.ResponseErr(code, message, err...)
}

func (j *JsonResponse) Response(code int, message string, data interface{}) string {
	resp := response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	result, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	return string(result)
}

func (j *JsonResponse) ResponseErr(code int, message string, errs ...error) string {
	data := []string{}
	for _, err := range errs {
		data = append(data, err.Error())
	}
	resp := response{
		Code:    code,
		Message: message,
		Data:    data,
	}
	result, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	return string(result)
}
