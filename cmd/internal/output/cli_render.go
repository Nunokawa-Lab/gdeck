package output

import (
	"encoding/json"
	"fmt"
	neturl "net/url"
	"time"

	"github.com/nunokawa/gdeck/cmd/internal/model"
)

// レスポンス共通関数
// TODO 関数名やこの関数が自体がprintしている構成も今後変える必要がありそうか検討する
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
			MethodIcon(method),
			u.Path,
			SelectStatusIcon(res.StatusCode),
			ColorStatus(res.Status, res.StatusCode),
			res.Time.Truncate(time.Microsecond),
			FormatJSON(hBytes),
			FormatJSON(res.Body),
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
			MethodIcon(method),
			u.Path,
			SelectStatusIcon(res.StatusCode),
			ColorStatus(res.Status, res.StatusCode),
			res.Time.Truncate(time.Microsecond),
			FormatJSON(res.Body),
		)
	}
}
