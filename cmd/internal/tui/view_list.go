package tui

import (
	"fmt"
	"path/filepath"
)

func (m Model) renderList() string {

	var s string
	for i, req := range m.requests {

		cursor := " "

		if i == m.cursor {
			// 一覧の位置とカーソルの位置が同じ箇所に矢印”>”を表示する
			cursor = ">"
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

	return listStyle.Render(s)
}