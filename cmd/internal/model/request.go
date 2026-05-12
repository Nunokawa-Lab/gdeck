package model

type Request struct {
	RequestName string `json:"request_name"`
	Method  string   `json:"method"`
	URL     string   `json:"url"`
	Headers []string `json:"headers"`
	Body    string   `json:"body"`
}
