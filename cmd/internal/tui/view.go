package tui

import (
	"fmt"
	"path/filepath"

	"github.com/nunokawa/gdeck/cmd/internal/output"
)

func (m Model) View() string {

	s := "\ngdeck TUI 😎\n\n"

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
		method := output.MethodColor(req.Method)
		method = output.PadRight(method, 8)
		s += fmt.Sprintf(
			"%s %s %s\n",
			cursor,
			method,
			cmdName,
		)
	}

	if m.selected != "" {
		s += fmt.Sprintf(
			"\nSelected: %s\n\n",
			m.selected,
		)
	}

	s += "\nq: quit\n\n"

	if m.response != nil {

		s += "\n--------------------\n\n"

		s += output.RenderTUIResponse(
			m.response,
			m.requests[m.cursor].Method,
		)
	}

	if m.errorMsg != "" {
		s += "\n❌ " + m.errorMsg + "\n"
	}

	return s
}
