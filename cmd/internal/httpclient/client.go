package httpclient

import (
	"io"
	"net/http"
)

// Getコマンドの実ロジック
func Get(url string) (string, error) {
	res, err := http.Get(url)
	if  err != nil {
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