package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderResponse(width int, height int) string {

	h := "📡 HTTP Response"

	if m.response == nil {
		h = "🔍 Preview"
	}

	paneStyle := inactivePaneStyle
	if m.focus == FocusResponse {
		paneStyle = activePaneStyle
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		headerLine(h, width, m.focus == FocusResponse),
		paneStyle.Render(
			lipgloss.NewStyle().
				Width(width).
				Height(height).
				Render(
					m.rightViewport.View(),
				),
		),
	)

}
