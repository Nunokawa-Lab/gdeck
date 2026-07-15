package tui

import (
	"fmt"

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
	paneStyle := activeRightPaneStyle // 保存中は常にアクティブ扱い

	var form string

	var labelStyle = lipgloss.NewStyle().
		Bold(true).
		Underline(true)

	if m.loading {
		form = fmt.Sprintf(
			"\n%s Saving Request...",
			m.spinner.View(),
		)
	} else {
		formParts := []string{
			labelStyle.Render("Name"),
			m.saveForm.name.View(),
			"",
			labelStyle.Render("Method"),
			m.saveForm.method.View(),
			"",
			labelStyle.Render("URL"),
			m.saveForm.url.View(),
			"",
			labelStyle.Render("Headers"),
			m.saveForm.header.View(),
			"",
			labelStyle.Render("Body"),
			m.saveForm.body.View(),
		}

		if m.errorMsg != "" {
			formParts = append([]string{
				errorMsgStyle.Render("⚠️  " + m.errorMsg),
				"",
			}, formParts...)
		}

		form = lipgloss.JoinVertical(
			lipgloss.Left,
			formParts...,
		)
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		painHeaderLine("🍀 New Request", width, true),
		paneStyle.Render(
			lipgloss.NewStyle().
				Width(width).
				Height(height).
				Render(form),
		),
	)
}
