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

func RenderTUIPreview(req *model.Request) string {

	body := "{}"
	if req.Body != "" {
		body = req.Body
	}

	return fmt.Sprintf(
		"# Method\n%s \n\n"+
			"# URL\n%v\n\n"+
			"# Header\n%v\n\n"+
			"# Body\n%s",
		AddIconToMethod(req.Method),
		req.URL,
		req.Headers,
		FormatJSON([]byte(body)),
	)
}
