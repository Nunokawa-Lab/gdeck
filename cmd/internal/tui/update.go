package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nunokawa/gdeck/cmd/internal/runner"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			if m.cursor > 0 {
				m.cursor--

				m.response = nil
				m.loadCurrentRequest()
			}
		case "down":
			if m.cursor < len(m.requests)-1 {
				m.cursor++

				m.response = nil
				m.loadCurrentRequest()
			}
		case "enter":
			selected := m.requests[m.cursor]

			m.selected = &selected

			results, err := runner.Run(
				selected.Name,
				runner.RunOptions{},
			)
			if err != nil {
				m.errorMsg = err.Error()
				return m, nil
			}
			if len(results) > 0 {
				m.response = results[0].Response
			}

			// 成功時はエラー消す
			m.errorMsg = ""
		}
	}

	return m, nil
}
