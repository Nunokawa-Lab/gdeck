package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderRightPane(width int, height int) string {
	if m.mode == ModeSave {
		return m.renderNewRequest(width, height)
	}

	return m.renderResponse(width, height)
}

// プレビュー・ローディング・レスポンス
func (m Model) renderResponse(width int, height int) string {

	var h string
	switch m.rightPaneView {
	case RightPaneResponse:
		h = "📡 Response"
	case RightPaneLoading:
		h = "📡 Response" // TODO "⏳ Running... " とかにしてもいいかも
	default:
		h = "🔍 Request Preview"
	}

	paneStyle := inactiveRightPaneStyle
	if m.focus == FocusResponse {
		paneStyle = activeRightPaneStyle
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		painHeaderLine(h, width, m.focus == FocusResponse),
		paneStyle.Render(
			lipgloss.NewStyle().
				Width(width).
				Height(height).
				Render(
					m.rightViewport.View(),
				),
		),
	)

}

// 新規リクエスト作成
func (m Model) renderNewRequest(width int, height int) string {
	// TODO あとで textinput に差し替える
	paneStyle := activeRightPaneStyle // 保存中は常にアクティブ扱い
	return lipgloss.JoinVertical(
		lipgloss.Left,
		painHeaderLine("✏️ New Request", width, true),
		paneStyle.Render(
			lipgloss.NewStyle().
				Width(width).
				Height(height).
				Render("New Request form (TODO)"),
		),
	)
} 