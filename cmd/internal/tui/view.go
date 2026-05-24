package tui

import "fmt"

func (m Model) View() string {

	s := "gdeck TUI 😎\n\n"

	for i, req := range m.requests {

		cursor := " "

		if i == m.cursor {
			// 一覧の位置とカーソルの位置が同じ箇所に矢印”>”を表示する
			cursor = ">"
		}

		s += fmt.Sprintf(
			"%s %s\n",
			cursor,
			req,
		)
	}

	if m.selected != "" {
		s += fmt.Sprintf(
			"\nSelected: %s\n\n",
			m.selected,
		)
	}

	s += "\nq: quit\n\n"

	return s
}
