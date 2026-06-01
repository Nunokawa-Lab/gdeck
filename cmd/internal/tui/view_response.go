package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderResponse(width int, height int) string {

	h := "📡 HTTP Response"

	if m.response == nil {
		h = "🔍 Preview"
	}

	borderStyle := inactivePaneStyle
	if m.focus == FocusResponse {
		borderStyle = activePaneStyle
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		headerStyle.Render(h),
		borderStyle.Render(
			lipgloss.NewStyle().
				Width(width).
				Height(height).
				Render(
					m.viewport.View(),
				),
		),
	)

}
