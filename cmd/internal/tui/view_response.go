package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderResponse() string {

	h := "📡 HTTP Response"

	if m.response == nil {
		h = "🔍 Preview"
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		headerStyle.Render(h),
		borderStyle.Render(
			m.viewport.View(),
		),
	)

}
