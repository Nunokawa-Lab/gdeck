package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

// レスポンス構造体
type Response struct {
	Status     string          `json:"status"`
	StatusCode int             `json:"status_code"`
	Header     http.Header     `json:"header"`
	Body       json.RawMessage `json:"body"`
	Time       time.Duration   `json:"time"`
}

// レスポンス共通関数
func (res *Response) PrintResponse(isVerbose bool) {
	if isVerbose {
		hBytes, err := json.Marshal(res.Header)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}

		fmt.Printf(
			"Status Code: %s\n\nHeader: %s\n\nBody: %s\n\nTime: %v",
			colorStatus(res.Status, res.StatusCode),
			FormatJSON(hBytes),
			FormatJSON(res.Body),
			res.Time,
		)

	} else {
		fmt.Printf(
			"Status Code: %s\n\nBody: %s\n\nTime: %v",
			colorStatus(res.Status, res.StatusCode),
			FormatJSON(res.Body),
			res.Time,
		)
	}
}

// ファイルとしてエクスポート
func (res *Response) WriteFile(path string, isVerbose bool) {

	helperWriteFile := func(data []byte) {
		err := os.WriteFile(path, data, 0644)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}
	}

	// 詳細オプションオン（-v）
	if isVerbose {
		if json.Valid(res.Body) {
			// レスポンスボディーがJSON形式の場合
			b, err := json.MarshalIndent(res, "", "  ")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			helperWriteFile(b)

		} else {
			// それ以外はテキストを出力
			hBytes, err := json.Marshal(res.Header)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				return
			}

			content := fmt.Sprintf(
				"Status Code: %v\n\nHeader: %s\n\nBody: %s\n\nTime: %v",
				res.StatusCode,
				FormatJSON(hBytes),
				FormatJSON(res.Body),
				res.Time,
			)
			helperWriteFile([]byte(content))
		}

	} else {
		if json.Valid(res.Body) {
			// レスポンスボディーがJSON形式の場合
			out := struct {
				StatusCode int             `json:"status_code"`
				Body       json.RawMessage `json:"body"`
				Time       time.Duration   `json:"time"`
			}{
				StatusCode: res.StatusCode,
				Body:       res.Body,
				Time:       res.Time,
			}

			b, err := json.MarshalIndent(out, "", "  ")
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			helperWriteFile(b)

		} else {
			// それ以外はテキストを出力
			content := fmt.Sprintf(
				"Status Code: %v\n\nBody: %s\n\nTime: %v",
				res.StatusCode,
				FormatJSON(res.Body),
				res.Time,
			)
			helperWriteFile([]byte(content))
		}
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
	start := time.Now()
	getRes, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	duration := time.Since(start)

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
		Time:       duration,
	}

	return res, nil
}

/********************
 POST
********************/
// 実ロジック
func Post(url string, requestBody string, headers []string) (*Response, error) {
	// 拡張性が乏しいため http.Post() でリクエストは作成しない
	req, err := http.NewRequest("POST", url, strings.NewReader(requestBody))
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
	start := time.Now()
	httpRes, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	duration := time.Since(start)
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
		Time:       duration,
	}
	return res, nil

}
