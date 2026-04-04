package httpclient

import (
	"io"
	"net/http"
	"strings"
)

/********************
 GET
********************/
// 実ロジック
func Get(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	// 最後に接続閉じる
	defer res.Body.Close()

	// res.Bodyはストリームのためbyte[]として書き出す
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}


/********************
 POST
********************/
// レスポンスの構造体
type Response struct {
	Status string
	Body   []byte
}
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

	b, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, err
	}

	res := &Response{
		Status: httpRes.Status,
		Body:   b,
	}
	return res, nil

}
