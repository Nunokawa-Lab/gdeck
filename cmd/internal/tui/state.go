package tui

import (
	"strings"

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
}

func (m *Model) resetSave() {
	m.mode = ModeNormal
	m.focus = FocusList
	m.showPreview()
}