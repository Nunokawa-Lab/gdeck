package model

import "strings"

type Request struct {
	Name    string   `json:"name"`
	Method  string   `json:"method"`
	URL     string   `json:"url"`
	Headers []string `json:"headers"`
	Body    string   `json:"body"`
}

func (r *Request) ToStringHeaders() string {
	// strings.Builderを使うことでメモリを無駄にせず文字列を作れる
	// （ループで文字列連結を行うと、都度メモリを確保してしまうため）
	var builder strings.Builder
	for _, header := range r.Headers {
		builder.WriteString(header)
		builder.WriteString("\n")
	}
	hStr := builder.String()

	return hStr
}
