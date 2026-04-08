package httpclient

import (
	"bytes"
	"encoding/json"
)

// JSONを整形する関数
func FormatJSON(data []byte) string {
	// 書き込み先のバッファ作成
	var out bytes.Buffer

	// json.Indent(書き込み先, 元データ, インデント開始, インデント文字)
	err := json.Indent(&out, data, "", "  ")
	if err != nil {
		// JSON以外だった場合は文字列として返す
		return string(data)
	}

	return out.String()
}
