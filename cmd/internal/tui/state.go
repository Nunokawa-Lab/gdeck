package tui

import (
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
	}
	focusedRequest := requests[m.cursor]

	reqs, err := store.Load(focusedRequest.Name)

	if err != nil {
		m.errorMsg = err.Error()
		return
	}

	if len(reqs) == 0 {
		return
	}

	m.currentRequest = reqs[0]
}

// 渡されたリクエストにカーソルを置くメソッド
func (m *Model) setSelectedRequest(selectedName string) {
	for i, req := range m.requests {
		if req.Name == selectedName {
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
