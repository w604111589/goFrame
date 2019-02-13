package common

import "time"

type ResponseData struct {
	Code int              `json:"code"`
	Data interface{}	  `json:"data"`
	Msg  string		  `json:"msg"`
	Time string		  `json:"time"`
}


func Success(data interface{}) *ResponseData{
	return &ResponseData{
		200,data,"成功",time.Now().Format("2006-01-02 15:04:05"),
	}
}

func SuccessMsg(msg string) *ResponseData{
	return &ResponseData{
		200,nil,msg,time.Now().Format("2006-01-02 15:04:05"),
	}
}


func Fail(code int,msg string) *ResponseData{
	return &ResponseData{
		code,nil,msg,time.Now().Format("2006-01-02 15:04:05"),
	}
}

func Error(code int) *ResponseData{
	return &ResponseData{
		code,nil,"请求异常",time.Now().Format("2006-01-02 15:04:05"),
	}
}


