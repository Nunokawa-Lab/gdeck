package tui

import (
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nunokawa/gdeck/cmd/internal/model"
	"github.com/nunokawa/gdeck/cmd/internal/store"
)

/** modelの状態を変更する処理を実装するファイル */

func (m *Model) initSearch() {
	m.mode = ModeSearch
	m.searchInput.SetValue("")
	m.searchInput.Focus()
	m.errorMsg = ""
}

func (m *Model) applySearch(text string) {
	if text == "" {
		m.filteredRequests = m.requests
		return
	}

	var filtered []model.RequestItem

	for _, req := range m.requests {
		if strings.Contains(
			strings.ToLower(req.Name),
			strings.ToLower(text),
		) {
			filtered = append(filtered, req)
		}
	}

	m.filteredRequests = filtered
}

func (m *Model) resetSearch() {
	m.mode = ModeNormal

	m.searchInput.Blur()
	m.searchInput.SetValue("")

	m.filteredRequests = []model.RequestItem{}
}

// カーソルの当たっているリクエストをロードするメソッド
func (m *Model) loadCurrentRequest() {

	requests := m.requests
	if m.mode == ModeSearch {
		if len(m.filteredRequests) > 0 {
			requests = m.filteredRequests
		} else {
			// 検索モード中に0件ヒットの場合は空を格納
			m.currentRequest = nil
			return
		}
	} else {
		// それ以外は1件もない時のガードを用意
		if len(m.requests) < 1 {
			m.currentRequest = nil
			return
		}
	}
	focusedRequest := requests[m.cursor]

	reqs, err := store.Load(focusedRequest.Name)
	if err != nil {
		m.errorMsg = err.Error()
		m.currentRequest = nil
		return
	}

	if len(reqs) == 0 {
		m.currentRequest = nil
		return
	}

	m.currentRequest = reqs[0]
}

// リスト上でカーソルが当たっているリクエスト名を返す
func (m *Model) selectedRequestName() (string, bool) {
	requests := m.requests
	if m.mode == ModeSearch {
		requests = m.filteredRequests
	}
	if len(requests) < 1 || m.cursor < 0 || m.cursor >= len(requests) {
		return "", false
	}
	return requests[m.cursor].Name, true
}

// 渡されたリクエストにカーソルを置くメソッド
func (m *Model) setSelectedRequest(selectedName string) {
	normalize := func(name string) string {
		return strings.TrimSuffix(name, filepath.Ext(name))
	}
	target := normalize(selectedName)

	for i, req := range m.requests {
		// req.Name → 基本は拡張子あり
		// selectedName → 呼び出し元で違う（保存成功時の時は拡張子がない）
		// そのため、拡張子を除いた比較も入れる
		if req.Name == selectedName || normalize(req.Name) == target {
			m.cursor = i
			break
		}
	}
}

func (m *Model) showPreview() {
	m.rightPaneView = RightPanePreview
	m.loading = false
	m.response = nil
}

func (m *Model) startLoading(selected *model.RequestItem) {
	m.rightPaneView = RightPaneLoading
	m.loading = true
	m.response = nil
	m.errorMsg = ""
	m.selected = selected
}

func (m *Model) showResponse(res *model.Response) {
	m.rightPaneView = RightPaneResponse
	m.loading = false
	m.response = res
}

func (m *Model) showErrorResponse() {
	m.rightPaneView = RightPaneResponse
	m.loading = false
	m.response = nil
}

// 端末サイズに合わせてペイン・viewport の寸法を更新する
func (m *Model) applyWindowSize(width, height int) {
	m.leftPaneWidth = int(float64(width) * 0.35)
	m.rightPaneWidth = width - m.leftPaneWidth - 8

	// header / search / footer などペイン以外の高さ（全モードで統一）
	m.paneHeight = height - 13

	// 奇数は正しい高さを割り出せないため偶数にする（改行含め2行でひとかたまりのため2の倍数が正）
	if m.paneHeight%2 != 0 {
		m.paneHeight++
	}
	if m.paneHeight < 0 {
		m.paneHeight = 0
	}

	m.leftViewport.Width = m.leftPaneWidth
	m.leftViewport.Height = m.paneHeight
	m.rightViewport.Width = m.rightPaneWidth
	m.rightViewport.Height = m.paneHeight

	// 表示中のリクエスト数セット（ペイン領域の高さの1/2が表示されている）
	m.displayRequestCnt = m.paneHeight / 2
}

func (m *Model) initSave() {
	m.mode = ModeSave
	m.focus = FocusResponse
	m.saveForm.focus = focusSaveFieldName
	m.errorMsg = ""
}

func (m *Model) resetSave() {
	m.mode = ModeNormal
	m.focus = FocusList
	m.saveForm.AllBlurFormFiled()
	m.saveForm.AllClearFormFiled()
	m.loading = false

	// 保存フォーム表示中に run が完了していた場合は結果を残す
	if m.rightPaneView == RightPaneResponse {
		return
	}

	m.errorMsg = ""
	m.showPreview()
}

// 全てのsaveFormのフォーカスカーソルを消す
func (sf *saveForm) AllBlurFormFiled() {
	sf.name.Blur()
	sf.method.Blur()
	sf.url.Blur()
	sf.header.Blur()
	sf.body.Blur()
}

// 全てのsaveFormの入力値を消す
func (sf *saveForm) AllClearFormFiled() {
	sf.name.SetValue("")
	sf.method.SetValue("")
	sf.url.SetValue("")
	sf.header.SetValue("")
	sf.body.SetValue("")
}

// saveForm.focusに設定
func (sf *saveForm) focusSaveFormFiled(focus SaveFocusFiled) tea.Cmd {
	sf.AllBlurFormFiled()
	switch focus {
	case 0:
		sf.focus = focusSaveFieldName
		return sf.name.Focus()
	case 1:
		sf.focus = focusSaveFieldMethod
		return sf.method.Focus()
	case 2:
		sf.focus = focusSaveFieldURL
		return sf.url.Focus()
	case 3:
		sf.focus = focusSaveFieldHeader
		return sf.header.Focus()
	case 4:
		sf.focus = focusSaveFieldBody
		return sf.body.Focus()
	}
	return nil
}

// saveFormの更新
// saveForm.focusの値でどのフォームを更新するか切り分け
// カーソルのBlink（点滅）も担当
func (sf *saveForm) updateForm(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	switch sf.focus {
	case focusSaveFieldName:
		sf.name, cmd = sf.name.Update(msg)
	case focusSaveFieldMethod:
		sf.method, cmd = sf.method.Update(msg)
	case focusSaveFieldURL:
		sf.url, cmd = sf.url.Update(msg)
	case focusSaveFieldHeader:
		sf.header, cmd = sf.header.Update(msg)
	case focusSaveFieldBody:
		sf.body, cmd = sf.body.Update(msg)
	}
	return cmd
}

func (m *Model) startSaveLoading() {
	m.loading = true
	m.errorMsg = ""
}
