package output

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	formatter "github.com/nunokawa/gdeck/cmd/internal"
	"github.com/nunokawa/gdeck/cmd/internal/model"
)

// ファイルとしてエクスポート
func WriteFile(res *model.Response, path string, isVerbose bool) {

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
				"Status: %v\n\nHeader: %s\n\nBody: %s\n\nTime: %v",
				res.StatusCode,
				formatter.FormatJSON(hBytes),
				formatter.FormatJSON(res.Body),
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
				"Status: %v\n\nBody: %s\n\nTime: %v",
				res.StatusCode,
				formatter.FormatJSON(res.Body),
				res.Time,
			)
			helperWriteFile([]byte(content))
		}
	}
}
