package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/nunokawa/gdeck/cmd/internal/output"
)

func (m Model) renderResponse() string {

	var h, s string

	if m.response == nil {
		h = "🔍 Preview"
		s = output.RenderTUIPreview(m.currentRequest)
	} else {
		h = "📡 HTTP Reponse"
		method := m.selected.Method
		s = output.RenderTUIResponse(m.response, method)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		headerStyle.Render(h),
		listStyle.Render(s),
	)

}
