package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {

	header := titleStyle.Render("\ngdeck TUI 😎")

	search := searchStyle.Render(m.searchBar())

	left := m.renderList(m.leftPaneWidth, m.paneHeight)
	right := m.renderResponse(m.rightPaneWidth, m.paneHeight)
	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		right,
	)
	footer := footerStyle.Render("↑↓ Move   Enter Run   ←→ Focus   q Quit")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		search,
		body,
		footer,
	)
}
