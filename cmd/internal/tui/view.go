package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {

	header := titleStyle.Render("gdeck TUI 😎")

	search := searchStyle.Render(m.searchBar())

	left := m.renderList(m.leftPaneWidth, m.paneHeight)
	right := m.renderResponse(m.rightPaneWidth, m.paneHeight)
	body := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		right,
	)
	footer := m.footer()

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		search,
		body,
		footer,
	)
}
