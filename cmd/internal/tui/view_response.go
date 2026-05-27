package tui

import (
	"github.com/nunokawa/gdeck/cmd/internal/output"
)

func (m Model) renderResponse() string {
	
	if m.selected == nil {
		return listStyle.Render("")
	}

	method := m.selected.Method
	return listStyle.Render(output.RenderTUIResponse(m.response, method))

}