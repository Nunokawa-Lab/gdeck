package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderResponse(width int, height int) string {

	h := "📡 Response"

	if m.response == nil {
		h = "🔍 Request Preview"
	}

	paneStyle := inactiveRightPaneStyle
	if m.focus == FocusResponse {
		paneStyle = activeRightPaneStyle
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		painHeaderLine(h, width, m.focus == FocusResponse),
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
