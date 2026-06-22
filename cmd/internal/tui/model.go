package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/nunokawa/gdeck/cmd/internal/model"
	"github.com/nunokawa/gdeck/cmd/internal/store"
)

type Model struct {
	requests       []model.RequestItem //リクエストアイテムリスト
	cursor         int                 //現在のカーソルの位置
	selected       *model.RequestItem  //選択されたリクエストアイテム
	currentRequest *model.Request      //カーソルが当たっているリクエスト情報

	response *model.Response //実行されたリクエストレスポンス
	loading  bool
	errorMsg string

	spinner       spinner.Model
	leftViewport  viewport.Model
	rightViewport viewport.Model

	focus FocusPane

	leftPaneWidth  int //左paneの幅
	rightPaneWidth int //右paneの幅
	paneHeight     int

	displayRequestCnt int //現在の左ペインの高さに対して、表示されているリクエストリストの数

	searchMode       bool
	filteredRequests []model.RequestItem

	searchInput textinput.Model
}

type FocusPane int

const (
	FocusList     FocusPane = iota //左paneのリストにユニークな連番をあてる
	FocusResponse                  //右pane
)

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
	forcusedRequest := requests[m.cursor]

	reqs, err := store.Load(forcusedRequest.Name)

	if err != nil {
		m.errorMsg = err.Error()
		return
	}

	if len(reqs) == 0 {
		return
	}

	m.currentRequest = reqs[0]
}

// 選択したリクエストにカーソルを置くメソッド
func (m *Model) setSelectedRequest(selectedName string) {
	for i, req := range m.requests {
		if req.Name == selectedName {
			m.cursor = i
			break
		}
	}
}
