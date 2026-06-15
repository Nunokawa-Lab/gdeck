package tui

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) requestListContent() string {

	var s string
	rows := make([]string, 0)

	requests := m.requests
	if m.searchMode {
		requests = m.filteredRequests
	}

	if len(requests) < 1 {
		s = "Oh!!\nNo-Hits :("
		return s
	} 

	for i, req := range requests {

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
		method = padRight(method, 10)

		row := fmt.Sprintf(
			"%s %s %s",
			cursor,
			method,
			cmdName,
		)

		row = lipgloss.NewStyle().PaddingBottom(1).Render(row)
		rows = append(rows, row)
	}

	s = strings.Join(rows, "\n")

	return s
}
