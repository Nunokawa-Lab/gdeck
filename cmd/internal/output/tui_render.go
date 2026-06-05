package output

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/nunokawa/gdeck/cmd/internal/model"
)

func RenderTUIResponse(
	res *model.Response,
	method string,
	width int,
	active bool,
) string {
	if width <= 0 {
		width = 1
	}

	separatorStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("14"))
	if !active {
		separatorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	}

	return fmt.Sprintf(
		"%s  |  "+
			"%s %s  |  "+
			"⏳ %v\n\n"+
			"%s\n\n"+
			"▼ Body\n\n%s",
		AddIconToMethod(method),
		SelectStatusIcon(res.StatusCode),
		res.Status,
		res.Time.Truncate(time.Millisecond),
		separatorStyle.Render(strings.Repeat("─", width)),
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
