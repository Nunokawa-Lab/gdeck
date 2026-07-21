package httpclient

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Nunokawa-Lab/gdeck/cmd/internal/model"
)

/** GET */
func Get(url string, options model.Options) (*model.Response, error) {
	return Do("GET", url, "", nil, options)
}

/** POST */
func Post(url string, requestBody string, headers []string, options model.Options) (*model.Response, error) {
	return Do("POST", url, requestBody, headers, options)
}

/** 共通処理関数 */
func Do(method string, url string, body string, headers []string, options model.Options) (*model.Response, error) {

	// bodyがあればio.Reader型に変換（NewRequest()第三引数に渡せるようにするため）
	var reader io.Reader
	if body != "" {
		reader = strings.NewReader(body)
	}

	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		return nil, err
	}

	for _, h := range headers {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) == 2 {
			// key:value の関係でAdd()
			req.Header.Add(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
		}
	}
	client := &http.Client{
		Timeout: time.Duration(options.Timeout) * time.Second,
	}

	start := time.Now()
	resp, err := client.Do(req)
	duration := time.Since(start)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &model.Response{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       bodyBytes,
		Time:       duration,
	}, nil
}
