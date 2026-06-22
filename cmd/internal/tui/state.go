package tui

import (
	"strings"

	"github.com/nunokawa/gdeck/cmd/internal/model"
	"github.com/nunokawa/gdeck/cmd/internal/store"
)

/** modelの状態を変更する処理を実装するファイル */

func (m *Model) initSearch() {
	m.searchMode = true
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
	m.searchMode = false

	m.searchInput.Blur()
	m.searchInput.SetValue("")

	m.filteredRequests = []model.RequestItem{}
}

// カーソルの当たっているリクエストをロードするメソッド
func (m *Model) loadCurrentRequest() {

	requests := m.requests
	if m.searchMode {
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
