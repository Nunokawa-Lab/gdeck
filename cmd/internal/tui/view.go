package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {

	header := titleStyle.Render("\ngdeck TUI 😎")

	left := m.renderList(m.leftPaneWidth, m.paneHeight)
	right := m.renderResponse(m.rightPaneWidth, m.paneHeight)
	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		right,
	)
	footer := footerStyle.Render("q: quit; ↑: up; ↓: down;")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		footer,
	)
}
