package model

type Request struct {
	Method  string   `json:"method"`
	URL     string   `json:"url"`
	Headers []string `json:"headers"`
	Body    string   `json:"body"`
}
