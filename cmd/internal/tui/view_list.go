package tui

import (
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) renderList(width int, height int) string {

	var s string
	for i, req := range m.requests {

		cursor := " "

		if i == m.cursor {
			// 一覧の位置とカーソルの位置が同じ箇所に専用マークを表示する
			cursor = "▌"
			cursor = styleDark.Render(cursor)
		}

		name := req.Name
		ext := filepath.Ext(name)
		cmdName := name[:len(name)-len(ext)]

		// スタイルが崩れないように先に色付けし、色付けした文字列を考慮して幅を揃える
		method := methodColor(req.Method)
		method = padRight(method, 8)
		s += fmt.Sprintf(
			"%s %s %s\n",
			cursor,
			method,
			cmdName,
		)
	}

	paneStyle := activePaneStyle
	if m.focus != FocusList {
		paneStyle = inactivePaneStyle
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		headerStyle.Render("📂 Requests"),
		paneStyle.Render(
			lipgloss.NewStyle().
				Width(width).
				Height(height).
				Render(s),
		),
	)
}
