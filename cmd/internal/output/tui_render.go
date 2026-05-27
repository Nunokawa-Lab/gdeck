package output

import (
	"fmt"
	"time"

	"github.com/nunokawa/gdeck/cmd/internal/model"
)

func RenderTUIResponse(
	res *model.Response,
	method string,
) string {

	return fmt.Sprintf(
		"%s \n"+
			"%s %s\n"+
			"⏳ %v\n\n"+
			"%s",
		AddIconToMethod(method),
		SelectStatusIcon(res.StatusCode),
		res.Status,
		res.Time.Truncate(time.Millisecond),
		FormatJSON(res.Body),
	)
}
