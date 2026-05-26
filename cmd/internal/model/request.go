package model

type Request struct {
	Name    string   `json:"name"`
	Method  string   `json:"method"`
	URL     string   `json:"url"`
	Headers []string `json:"headers"`
	Body    string   `json:"body"`
}
