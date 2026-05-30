package tui

import (

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {

	header := titleStyle.Render("gdeck TUI 😎")
	left := m.renderList()
	right := m.renderResponse()
	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		right,
	)
	footer := footerStyle.Render("q: quit")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		body,
		footer,
	)
}
