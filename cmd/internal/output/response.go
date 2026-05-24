package output

import (
	"encoding/json"
	"fmt"
	neturl "net/url"
	"time"

	formatter "github.com/nunokawa/gdeck/cmd/internal"
	"github.com/nunokawa/gdeck/cmd/internal/model"
)

// レスポンス共通関数
func PrintResponse(
	res *model.Response,
	isVerbose bool,
	currentNum int,
	len int,
	method string,
	reqeustName string,
	url string,
) {

	u, err := neturl.Parse(url)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	// 詳細出力か
	if isVerbose {
		hBytes, err := json.Marshal(res.Header)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return
		}

		fmt.Printf(
			`┌─────────────────────────────
│ [%v/%v] %s
└─────────────────────────────
%s  %s

%s Status  %s
⏳ Time    %v

📨 Header
%s

📦 Body
%s

`,
			currentNum,
			len,
			reqeustName,
			methodIcon(method),
			u.Path,
			selectStatusIcon(res.StatusCode),
			formatter.ColorStatus(res.Status, res.StatusCode),
			res.Time.Truncate(time.Microsecond),
			formatter.FormatJSON(hBytes),
			formatter.FormatJSON(res.Body),
		)

	} else {
		fmt.Printf(
			`┌─────────────────────────────
│ [%v/%v] %s
└─────────────────────────────
%s  %s

%s Status  %s
⏳ Time    %v

📦 Body
%s

`,
			currentNum,
			len,
			reqeustName,
			methodIcon(method),
			u.Path,
			selectStatusIcon(res.StatusCode),
			formatter.ColorStatus(res.Status, res.StatusCode),
			res.Time.Truncate(time.Microsecond),
			formatter.FormatJSON(res.Body),
		)
	}
}

func selectStatusIcon(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "✅" // success
	case code >= 300 && code < 500:
		return "⚠️ " // warning（半角空白は必要なため消さない）
	case code >= 500:
		return "❌" // error
	default:
		return ""
	}
}

func methodIcon(method string) string {
	switch method {
	case "GET":
		return "🔵 GET"
	case "POST":
		return "🟢 POST"
	case "PUT":
		return "🟡 PUT"
	case "PATCH":
		return "🟣 PATCH"
	case "DELETE":
		return "🔴 DELETE"
	default:
		return "⚪ Unknown Method"
	}
}
