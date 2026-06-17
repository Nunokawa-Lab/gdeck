package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// 検索モード中の挙動
	if m.searchMode {

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				m.resetSearch()
				m.leftViewport.SetContent(m.requestListContent())
				return m, nil

			case "up":
				if m.cursor > 0 {
					m.cursor--

					m.response = nil

					// スクロール
					if m.cursor <= ((len(m.filteredRequests) - 1) - m.displayRequestCnt) {
						m.leftViewport.ScrollUp(2)
					}
				}
			case "down":
				if m.cursor < len(m.filteredRequests)-1 {
					m.cursor++

					m.response = nil

					// スクロール
					if m.cursor >= m.displayRequestCnt {
						m.leftViewport.ScrollDown(2)
					}
				}
			default:
				// 各値を初期値に戻す
				m.response = nil
				m.cursor = 0
			}
		}

		// カーソルを点滅させるためにBlinkMsg型も含めて全msgをUpdate()に渡す必要あり
		m.searchInput, cmd = m.searchInput.Update(msg)

		// 絞り込み
		m.applySearch(m.searchInput.Value())
		m.loadCurrentRequest()
		m.leftViewport.SetContent(m.requestListContent())
		m.rightViewport.SetContent(m.responseContent())

		return m, cmd

	}

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "left":
			m.focus = FocusList
		case "right":
			m.focus = FocusResponse
		case "/":
			// 検索モード初期設定
			m.initSearch()

			// 他の各値も初期値に戻す
			m.focus = FocusList
			m.response = nil
			m.cursor = 0
			m.loadCurrentRequest()
			m.rightViewport.SetContent(m.responseContent())

			return m, textinput.Blink
		}

		// 左pane挙動
		if m.focus == FocusList {
			m.rightViewport.SetContent(m.responseContent())

			switch msg.String() {
			case "up":
				if m.cursor > 0 {
					m.cursor--

					m.response = nil
					m.loadCurrentRequest()
					m.leftViewport.SetContent(m.requestListContent())
					m.rightViewport.SetContent(m.responseContent())

					// スクロール
					if m.cursor <= ((len(m.requests) - 1) - m.displayRequestCnt) {
						m.leftViewport.ScrollUp(2)
					}
				}
			case "down":
				if m.cursor < len(m.requests)-1 {
					m.cursor++

					m.response = nil
					m.loadCurrentRequest()
					m.leftViewport.SetContent(m.requestListContent())
					m.rightViewport.SetContent(m.responseContent())

					// スクロール
					if m.cursor >= m.displayRequestCnt {
						m.leftViewport.ScrollDown(2)
					}
				}
			case "enter":
				selected := m.requests[m.cursor]

				// 実行前に初期値をセット
				m.loading = true
				m.response = nil
				m.errorMsg = ""
				m.selected = &selected
				// ローディングUI表示
				m.rightViewport.SetContent(m.responseContent())

				return m, asyncRunCmd(selected.Name, selected.Method)
			}
		}

		// 右pane挙動
		if m.focus == FocusResponse {
			m.rightViewport.SetContent(m.responseContent())

			// スクロール
			switch msg.String() {

			case "up":
				m.rightViewport.ScrollUp(1)
			case "down":
				m.rightViewport.ScrollDown(1)
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
		m.rightViewport.SetContent(m.responseContent())

		return m, nil
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)

		if m.loading {
			m.rightViewport.SetContent(
				m.responseContent(),
			)
		}

		return m, cmd
	case tea.WindowSizeMsg:
		// サイズセット
		m.leftPaneWidth = int(float64(msg.Width) * 0.35)
		m.rightPaneWidth = msg.Width - m.leftPaneWidth - 8
		m.paneHeight = msg.Height - 11

		// viewportにも高さ・幅をセット
		m.leftViewport.Width = m.leftPaneWidth
		m.leftViewport.Height = m.paneHeight
		m.rightViewport.Width = m.rightPaneWidth
		m.rightViewport.Height = m.paneHeight

		// 表示中のリクエスト数セット（ペイン領域の高さの1/2が表示されている）
		m.displayRequestCnt = m.paneHeight / 2
	}

	return m, cmd
}
