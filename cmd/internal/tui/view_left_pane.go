package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderLeftPane(width int, height int) string {

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

	requests := m.requests
	if m.mode == ModeSearch {
		requests = m.filteredRequests
	}

	requestCount := len(requests)
	visibleCount := m.displayRequestCnt

	// 表示中の高さに収まる場合、最後までスクロールしている場合もh何もテキストを出さない
	if visibleCount <= 0 || requestCount <= visibleCount {
		return ""
	}

	// 残り件数が0以下なら末尾に到達している
	remaining := requestCount - visibleCount - (m.leftViewport.YOffset / 2)
	if remaining <= 0 {
		return ""
	}

	return "↓ More …"
}
