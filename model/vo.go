package model

import "time"

type Result struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type UrlApply struct {
	LangUrl  string `json:"lang_url"`
	Duration string `json:"duration"`
	Token    string
}

func NewErrorResult(err error) Result {
	return Result{
		Code:      500,
		Msg:       err.Error(),
		Data:      nil,
		Timestamp: time.Now().Unix(),
	}
}
