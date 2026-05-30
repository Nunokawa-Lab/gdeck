package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/nunokawa/gdeck/cmd/internal/output"
)

func (m Model) renderResponse() string {
	
	var s string
	
	if m.selected != nil {
		method := m.selected.Method
		s = output.RenderTUIResponse(m.response, method)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		headerStyle.Render("📡 HTTP Reponse"),
		listStyle.Render(s),
	)

}