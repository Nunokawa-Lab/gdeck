package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	// "github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nunokawa/gdeck/cmd/internal/model"
	"github.com/nunokawa/gdeck/cmd/internal/store"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case clearStatusMsg:
		m.statusMsg = ""
		return m, nil
	case runFinishedMsg:
		if msg.err != nil {
			m.errorMsg = msg.err.Error()
			m.showErrorResponse()
		} else {
			m.showResponse(msg.response)
		}

		// 検索モード中に完了した場合は検索を解除し、実行したリクエストにカーソルを戻す
		if m.mode == ModeSearch {
			if len(m.filteredRequests) > 0 {
				selected := m.filteredRequests[m.cursor]
				m.resetSearch()
				m.setSelectedRequest(selected.Name)

				if m.displayRequestCnt < m.cursor+1 {
					m.leftViewport.YOffset = ((m.cursor + 1) - m.displayRequestCnt) * 2
				}
			} else {
				m.resetSearch()
				m.cursor = 0
			}
			m.leftViewport.SetContent(m.requestListContent())
		}

		m.rightViewport.SetContent(m.responseContent())
		return m, nil
	case tea.WindowSizeMsg:
		m.applyWindowSize(msg.Width, msg.Height)
		m.leftViewport.SetContent(m.requestListContent())
		m.rightViewport.SetContent(m.responseContent())
		return m, nil
	}

	/** 検索モード中の挙動 */
	if m.mode == ModeSearch {

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				if len(m.filteredRequests) < 1 {
					m.resetSearch()
					m.cursor = 0
				} else {
					// 検索モードを解除しつつカーソルは当てていたリクエストの位置をキープする
					selected := m.filteredRequests[m.cursor] //リセットする前にカーソルの当たっているリクエスト取得
					m.resetSearch()
					m.setSelectedRequest(selected.Name)

					// もしキープした位置がスクロールしないと見えない位置なら、見える位置までオフセットを調整
					if m.displayRequestCnt < m.cursor+1 {
						m.leftViewport.YOffset = ((m.cursor + 1) - m.displayRequestCnt) * 2
					}
				}

				m.leftViewport.SetContent(m.requestListContent())
				m.rightViewport.SetContent(m.responseContent())
				return m, nil
			case "up":
				if m.cursor > 0 {
					m.cursor--

					m.errorMsg = ""
					m.showPreview()

					// スクロール判定：カーソルが見える範囲の上を超えたら
					firstVisibleIndex := m.leftViewport.YOffset / 2
					if m.cursor < firstVisibleIndex {
						m.leftViewport.ScrollUp(2)
					}
				}
			case "down":
				if m.cursor < len(m.filteredRequests)-1 {
					m.cursor++

					m.errorMsg = ""
					m.showPreview()

					// スクロール判定：カーソルが見える範囲の下を超えたら
					firstVisibleIndex := m.leftViewport.YOffset / 2
					lastVisibleIndex := firstVisibleIndex + m.displayRequestCnt - 1
					if m.cursor > lastVisibleIndex {
						m.leftViewport.ScrollDown(2)
					}
				}
			case "enter":
				if len(m.filteredRequests) < 1 {
					return m, nil
				}

				selected := m.filteredRequests[m.cursor]

				m.startLoading(&selected)
				m.rightViewport.SetContent(m.responseContent())

				return m, asyncRunCmd(selected.Name, selected.Method)
			default:
				// 各値を初期値に戻す
				m.errorMsg = ""
				m.showPreview()
				m.cursor = 0
			}
		case spinner.TickMsg:
			m.spinner, cmd = m.spinner.Update(msg)

			if m.rightPaneView == RightPaneLoading {
				m.rightViewport.SetContent(
					m.responseContent(),
				)
			}

			return m, cmd
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

	/** 削除確認モード中の挙動 */
	if m.mode == ModeDeleteConfirm {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "y":
				name, ok := m.selectedRequestName()
				if !ok {
					m.errorMsg = "no request selected"
					m.rightViewport.SetContent(m.responseContent())
					return m, nil
				}

				// 削除
				err := store.Delete(name)
				if err != nil {
					m.errorMsg = err.Error()
					m.rightViewport.SetContent(m.responseContent())
					return m, nil
				}

				// リクエストを再読み込み
				m.requests, err = store.List()
				if err != nil {
					m.errorMsg = err.Error()
					return m, nil
				}

				// 現在のカーソルの当たっているリクエストを再取得
				if len(m.requests) == 0 {
					m.cursor = 0
				} else if m.cursor >= len(m.requests) {
					m.cursor = len(m.requests) - 1
				}
				m.loadCurrentRequest()

				m.rightPaneView = RightPanePreview

				m.leftViewport.SetContent(m.requestListContent())
				m.rightViewport.SetContent(m.responseContent())

				m.mode = ModeNormal

			case "n":
				m.mode = ModeNormal
				m.errorMsg = ""
				m.rightViewport.SetContent(m.responseContent())
			case "esc":
				// エラーで失敗した時にフッターの説明に「esc」がでるため、errorMsgに値がある時だけ動作させる
				if m.errorMsg != "" {
					m.mode = ModeNormal
					m.errorMsg = ""
					m.rightViewport.SetContent(m.responseContent())
				}
			}
		}

		return m, nil
	}

	/** 保存モード時の挙動 */
	if m.mode == ModeSave {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc":
				m.resetSave()
				m.leftViewport.SetContent(m.requestListContent())
				m.rightViewport.SetContent(m.responseContent())
				return m, nil
			case "ctrl+s":
				// 必須項目
				name := strings.TrimSpace(m.saveForm.name.Value())
				method := strings.ToUpper(strings.TrimSpace(m.saveForm.method.Value()))
				url := strings.TrimSpace(m.saveForm.url.Value())
				if name == "" || method == "" || url == "" {
					m.errorMsg = "Name, Method, and URL are required"
					return m, nil
				}

				// 任意項目
				var headers []string
				for _, line := range strings.Split(m.saveForm.header.Value(), "\n") {
					line = strings.TrimSpace(line)
					if line != "" {
						headers = append(headers, line)
					}
				}
				body := m.saveForm.body.Value()

				m.errorMsg = ""
				req := &model.Request{
					Name:    name,
					Method:  method,
					URL:     url,
					Headers: headers,
					Body:    body,
				}

				m.startSaveLoading()

				// tea.Batch は「複数の非同期処理を並行で予約する」ための API
				// 保存開始と同時に次の spinner.TickMsg が予約されるため、待たずに「Saving Request...」のドットが回り始めやすくなる
				// 同時に case spinner.TickMsg 内で spinner.TickMsg を処理する必要あり
				return m, tea.Batch(asyncSaveCmd(req, m.editingRequestName), m.spinner.Tick)
			case "shift+tab":
				m.errorMsg = ""
				if m.saveForm.focus > 0 {
					m.saveForm.focus--
					return m, m.saveForm.focusSaveFormFiled(m.saveForm.focus)
				}
			case "tab":
				m.errorMsg = ""
				focus := m.saveForm.toIntFocus()
				if (focus + 1) < m.saveFormFieldCount {
					m.saveForm.focus++
					return m, m.saveForm.focusSaveFormFiled(m.saveForm.focus)
				}
			default:
				m.errorMsg = ""
			}
		case saveFinishedMsg:
			if msg.err != nil {
				m.loading = false
				m.errorMsg = msg.err.Error()
				return m, nil
			}

			// リクエストを再読み込み
			requests, err := store.List()
			if err != nil {
				m.loading = false
				m.errorMsg = err.Error()
				return m, nil
			}
			m.requests = requests

			if msg.updated {
				m.statusMsg = fmt.Sprintf("✓ Updated %s", msg.name)
			} else {
				m.statusMsg = fmt.Sprintf("✓ Saved %s", msg.name)
			}
			m.resetSave()
			m.showPreview()
			m.setSelectedRequest(msg.name)
			if m.displayRequestCnt < m.cursor+1 {
				m.leftViewport.YOffset = ((m.cursor + 1) - m.displayRequestCnt) * 2
			} else {
				m.leftViewport.YOffset = 0
			}
			m.loadCurrentRequest()

			m.leftViewport.SetContent(m.requestListContent())
			m.rightViewport.SetContent(m.responseContent())

			// 2秒間だけ成功メッセージ表示
			return m, clearStatusAfter(2 * time.Second)
		case spinner.TickMsg:
			m.spinner, cmd = m.spinner.Update(msg)
			return m, cmd
		}

		cmd = m.saveForm.updateForm(msg)

		return m, cmd
	}

	/** 通常時の挙動 */
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
			if len(m.requests) < 1 {
				return m, nil
			}

			// 検索モード初期設定
			m.initSearch()

			// 他の各値も初期値に戻す
			m.focus = FocusList
			m.showPreview()
			m.cursor = 0
			m.loadCurrentRequest()
			m.leftViewport.SetContent(m.requestListContent())
			m.rightViewport.SetContent(m.responseContent())

			return m, textinput.Blink
		case "s":
			m.initSave()
			return m, m.saveForm.focusSaveFormFiled(focusSaveFieldName)
		case "e":
			if !m.initEdit() {
				return m, nil
			}
			return m, m.saveForm.focusSaveFormFiled(focusSaveFieldName)
		}

		// 左pane挙動
		if m.focus == FocusList {
			m.rightViewport.SetContent(m.responseContent())

			switch msg.String() {
			case "up":
				if m.cursor > 0 {
					m.cursor--

					// スクロール判定：カーソルが見える範囲の上を超えたら
					firstVisibleIndex := m.leftViewport.YOffset / 2
					if m.cursor < firstVisibleIndex {
						m.leftViewport.ScrollUp(2)
					}

					m.errorMsg = ""
					m.showPreview()
					m.loadCurrentRequest()
					m.leftViewport.SetContent(m.requestListContent())
					m.rightViewport.SetContent(m.responseContent())

				}
			case "down":
				if m.cursor < len(m.requests)-1 {
					m.cursor++

					// スクロール判定：カーソルが見える範囲の下を超えたら
					firstVisibleIndex := m.leftViewport.YOffset / 2
					lastVisibleIndex := firstVisibleIndex + m.displayRequestCnt - 1
					if m.cursor > lastVisibleIndex {
						m.leftViewport.ScrollDown(2)
					}

					m.errorMsg = ""
					m.showPreview()
					m.loadCurrentRequest()
					m.leftViewport.SetContent(m.requestListContent())
					m.rightViewport.SetContent(m.responseContent())

				}
			case "enter":
				if len(m.requests) < 1 {
					return m, nil
				}

				selected := m.requests[m.cursor]

				m.startLoading(&selected)
				m.rightViewport.SetContent(m.responseContent())

				return m, asyncRunCmd(selected.Name, selected.Method)
			case "d":
				if _, ok := m.selectedRequestName(); !ok {
					return m, nil
				}
				if m.currentRequest == nil {
					return m, nil
				}

				// 削除確認モードオン
				m.mode = ModeDeleteConfirm
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
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)

		if m.rightPaneView == RightPaneLoading {
			m.rightViewport.SetContent(
				m.responseContent(),
			)
		}

		return m, cmd
	}

	return m, cmd
}
