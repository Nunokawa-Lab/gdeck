package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/nunokawa/gdeck/cmd/internal/output"
)

func (m Model) renderResponse() string {

	var s string
	h := "📡 HTTP Reponse"

	if m.loading {
		s = fmt.Sprintf(
			"\n%s Sending Request...",
			m.spinner.View(),
		)
	} else if m.response == nil {
		h = "🔍 Preview"
		s = output.RenderTUIPreview(m.currentRequest)
	} else {
		method := m.selected.Method
		s = output.RenderTUIResponse(m.response, method)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		headerStyle.Render(h),
		listStyle.Render(s),
	)

}
