package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// レスポンス構造体
type Response struct {
	Status     string
	StatusCode int
	Header     http.Header
	Body       []byte
}

// レスポンス共通関数
func (res *Response) PrintResponse(isVerbose bool) {
	if isVerbose {
		headerJSONBytes, err := json.Marshal(res.Header)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}

		fmt.Printf(
			"Status Code: %s\n\nHeader: %s\n\nBody: %s",
			colorStatus(res.Status, res.StatusCode),
			FormatJSON(headerJSONBytes),
			FormatJSON(res.Body),
		)
	} else {
		fmt.Printf("Status Code: %v\n", colorStatus(res.Status, res.StatusCode))
		fmt.Println(FormatJSON(res.Body))
	}
}

// ステータスコード色付け
func colorStatus(status string, code int) string {
	switch {
	case code >= 200 && code < 300:
		return "\033[42;30m " + status + " \033[0m" // 緑背景＋黒文字
	case code >= 300 && code < 400:
		return "\033[44;37m " + status + " \033[0m" // 青背景＋白文字
	case code >= 400 && code < 500:
		return "\033[43;30m " + status + " \033[0m" // 黄背景＋黒文字
	case code >= 500:
		return "\033[41;37m " + status + " \033[0m" // 赤背景＋白文字
	default:
		return status
	}
}

/********************
 GET
********************/
// 実ロジック
func Get(url string) (*Response, error) {
	getRes, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// 最後に接続閉じる
	defer getRes.Body.Close()

	if hLen := len(getRes.Header); hLen < 1 {
		return nil, err
	}

	// res.Bodyはストリームのためbyte[]として書き出す
	body, err := io.ReadAll(getRes.Body)
	if err != nil {
		return nil, err
	}

	res := &Response{
		Status:     getRes.Status,
		StatusCode: getRes.StatusCode,
		Header:     getRes.Header,
		Body:       body,
	}

	return res, nil
}

/********************
 POST
********************/
// 実ロジック
func Post(url string, body string, headers []string) (*Response, error) {
	// 拡張性が乏しいため http.Post() でリクエストは作成しない
	req, err := http.NewRequest("POST", url, strings.NewReader(body))
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

	client := &http.Client{}
	httpRes, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpRes.Body.Close()

	if hLen := len(httpRes.Header); hLen < 1 {
		return nil, err
	}

	b, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	res := &Response{
		Status:     httpRes.Status,
		StatusCode: httpRes.StatusCode,
		Header:     httpRes.Header,
		Body:       b,
	}
	return res, nil

}
