package model

import "time"

type RunResult struct {
	Request  *Request
	Response *Response
	Error    error
	Duration time.Duration
}