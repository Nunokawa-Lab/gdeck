package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		// 共通
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "left":
			m.focus = FocusList
		case "right":
			m.focus = FocusResponse
		}

		// 左pane挙動
		if m.focus == FocusList {
			switch msg.String() {
			case "up":
				if m.cursor > 0 {
					m.cursor--

					m.response = nil
					m.loadCurrentRequest()
					m.viewport.SetContent(m.responseContent())
				}
			case "down":
				if m.cursor < len(m.requests)-1 {
					m.cursor++

					m.response = nil
					m.loadCurrentRequest()
					m.viewport.SetContent(m.responseContent())
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
		}

		// 右pane挙動
		if m.focus == FocusResponse {
			// スクロール
			switch msg.String() {

			case "up":
				m.viewport.ScrollUp(1)
			case "down":
				m.viewport.ScrollDown(1)
			}

			return m, nil
		}
	case runFinishedMsg:
		m.loading = false

		if msg.err != nil {
			m.errorMsg = msg.err.Error()
			return m, nil
		}

		m.response = msg.response

		// コンテンツをviewportにセット
		m.viewport.SetContent(m.responseContent())

		return m, nil
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)

		return m, cmd
	case tea.WindowSizeMsg:
		// サイズセット
		m.leftPaneWidth = int(float64(msg.Width) * 0.3)
		m.rightPaneWidth = msg.Width - m.leftPaneWidth - 8
		m.paneHeight = msg.Height - 12

		// viewportにも高さ・幅をセット
		m.viewport.Width = m.rightPaneWidth
		m.viewport.Height = m.paneHeight
	}

	return m, cmd
}
