package model

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Status     string          `json:"status"`
	StatusCode int             `json:"status_code"`
	Header     http.Header     `json:"header"`
	Body       json.RawMessage `json:"body"`
	Time       time.Duration   `json:"time"`
}