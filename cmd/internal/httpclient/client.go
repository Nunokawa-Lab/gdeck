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
	// 詳細出力か
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

	// 詳細出力か
	if isVerbose {
		// JSON出力か
		if json.Valid(res.Body) {
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
		// JSON出力か
		if json.Valid(res.Body) {
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
func Get(url string) (*Response, error) {
	return Do("GET", url, "", nil)
}

/********************
 POST
********************/
func Post(url string, requestBody string, headers []string) (*Response, error) {
	return Do("POST", url, requestBody, headers)
}

/********************
 共通処理関数
********************/
func Do(method string, url string, body string, headers []string) (*Response, error) {
	
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

	client := &http.Client{}

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

	return &Response{
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       bodyBytes,
		Time:       duration,
	}, nil
}