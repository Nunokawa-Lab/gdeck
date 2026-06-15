package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderList(width int, height int) string {

	paneStyle := activeLeftPaneStyle
	if m.focus != FocusList {
		paneStyle = inactiveLeftPaneStyle
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		painHeaderLine("📂 Requests", width, (m.focus == FocusList)),
		paneStyle.Render(
			lipgloss.NewStyle().
				Width(width).
				Height(height).
				Render(
					m.leftViewport.View(),
				),
		),
		painFooterLine(m.requestFooterText(), width, (m.focus == FocusList)),
	)
}

func (m Model) requestFooterText() string {
	requestCount := len(m.requests)
	visibleCount := m.displayRequestCnt

	if visibleCount <= 0 || requestCount <= visibleCount {
		return "No more …"
	}

	// スクロールされ隠れた上の行数を数える
	remaining := requestCount - visibleCount - (m.leftViewport.YOffset / 2)
	if remaining <= 0 {
		return "End."
	}

	return fmt.Sprintf("↓ More %d …", remaining)
}
