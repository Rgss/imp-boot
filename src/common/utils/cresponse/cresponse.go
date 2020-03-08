package cresponse

// 响应结果
type Response struct {
	ErrorCode int         `json:"errCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

// 成功
const ResponseErrorCodeSuccess = 0

// 失败
const ResponseErrorCodeCommon = 1

// 默认提示
const SuccessMessage = "ok"

// 公用空响应结果
var EmptyResponseData = make(map[string]interface{})

// 对外使用
var ResponseResult = Response{
	ErrorCode: ResponseErrorCodeSuccess,
	Message:   "success",
	Data:      make(map[string]interface{}),
}

// 失败响应
func Fail(args ...interface{}) Response {
	errCode := ResponseErrorCodeCommon
	message := "系统繁忙！"

	for key, val := range args {
		if key == 0 {
			message = val.(string)
		}

		if key == 1 {
			errCode = val.(int)
		}
	}

	resp := Response{
		ErrorCode: errCode,
		Message:   message,
		Data:      make(map[string]interface{}),
	}
	return resp
}

// 成功响应
func Success(args ...interface{}) Response {
	data := EmptyResponseData
	message := SuccessMessage

	for key, val := range args {
		if key == 0 {
			data = val.(map[string]interface{})
		}

		if key == 1 {
			message = val.(string)
		}
	}

	resp := Response{
		ErrorCode: ResponseErrorCodeSuccess,
		Message:   message,
		Data:      data,
	}
	return resp
}
