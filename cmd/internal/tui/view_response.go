package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderResponse(width int, height int) string {

	h := "📡 HTTP Response"

	if m.response == nil {
		h = "🔍 Preview"
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		headerStyle.Render(h),
		lipgloss.NewStyle().
			Width(width).
			Height(height).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("8")).
			Render(
				m.viewport.View(),
			),
	)

}
