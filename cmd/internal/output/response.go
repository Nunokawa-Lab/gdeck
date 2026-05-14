package output

import (
	"encoding/json"
	"fmt"

	formatter "github.com/nunokawa/gdeck/cmd/internal"
	"github.com/nunokawa/gdeck/cmd/internal/model"
)

// レスポンス共通関数
func PrintResponse(res *model.Response, isVerbose bool) {
	// 詳細出力か
	if isVerbose {
		hBytes, err := json.Marshal(res.Header)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}

		fmt.Printf(
			"Status Code: %s\n\nHeader: %s\n\nBody: %s\n\nTime: %v\n\n",
			formatter.ColorStatus(res.Status, res.StatusCode),
			formatter.FormatJSON(hBytes),
			formatter.FormatJSON(res.Body),
			res.Time,
		)

	} else {
		fmt.Printf(
			"Status Code: %s\n\nBody: %s\n\nTime: %v\n\n",
			formatter.ColorStatus(res.Status, res.StatusCode),
			formatter.FormatJSON(res.Body),
			res.Time,
		)
	}
}
