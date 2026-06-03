package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderList(width int, height int) string {

	paneStyle := activePaneStyle
	if m.focus != FocusList {
		paneStyle = inactivePaneStyle
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		headerLine("📂 Requests", width, (m.focus == FocusList)),
		paneStyle.Render(
			lipgloss.NewStyle().
				Width(width).
				Height(height).
				Render(
					m.leftViewport.View(),
				),
		),
	)
}
