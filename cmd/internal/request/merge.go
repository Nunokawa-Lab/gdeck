package request

import (
	"fmt"
	"strings"
)

// ヘッダーの追加・上書き処理
func MergeHeaders(baseHeader []string, overrideHeader []string) []string {
	headerMap := make(map[string]string)

	// slice ▶︎ map
	for _, h := range baseHeader {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) < 2 {
			continue
		}

		// 置換しやすいように空白消す
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		// キーは入力値で差異がでないように小文字に統一
		headerMap[strings.ToLower(key)] = fmt.Sprintf("%s: %s", key, val)
	}

	// overrideで上書き or 追加
	for _, h := range overrideHeader {
		parts := strings.SplitN(h, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])

		headerMap[strings.ToLower(key)] = fmt.Sprintf("%s: %s", key, val)
	}

	var results []string
	for _, val := range headerMap {
		results = append(results, val)
	}

	return results
}
