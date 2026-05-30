package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
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

			// 実行前に初期値をセット
			m.loading = true
			m.response = nil
			m.errorMsg = ""
			m.selected = &selected

			return m, asyncRunCmd(selected.Name, selected.Method)
		}
	case runFinishedMsg:
		m.loading = false

		if msg.err != nil {
			m.errorMsg = msg.err.Error()
			return m, nil
		}

		m.response = msg.response

		return m, nil
	case spinner.TickMsg:
		var cmd tea.Cmd

		m.spinner, cmd = m.spinner.Update(msg)

		return m, cmd
	}

	return m, nil
}
