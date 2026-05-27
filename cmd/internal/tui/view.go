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


	// if m.selected != "" {
	// 	s += fmt.Sprintf(
	// 		"\nSelected: %s\n\n",
	// 		m.selected,
	// 	)
	// }

	// s += "\nq: quit\n\n"


	// if m.response != nil {

	// 	s += "\n--------------------\n\n"

	// 	s += output.RenderTUIResponse(
	// 		m.response,
	// 		m.requests[m.cursor].Method,
	// 	)
	// }

	// if m.errorMsg != "" {
	// 	s += "\n❌ " + m.errorMsg + "\n"
	// }

	// return s
}
