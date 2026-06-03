package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/nunokawa/gdeck/cmd/internal/model"
)

type Model struct {
	requests       []model.RequestItem //リクエストアイテムリスト
	cursor         int                 //現在のカーソルの位置
	selected       *model.RequestItem  //選択されたリクエストアイテム
	currentRequest *model.Request      //カーソルが当たっているリクエスト情報

	response *model.Response //実行されたリクエストレスポンス
	loading  bool
	errorMsg string

	spinner  spinner.Model
	leftViewport viewport.Model
	rightViewport viewport.Model

	focus FocusPane

	leftPaneWidth  int //左paneの幅
	rightPaneWidth int //右paneの幅
	paneHeight     int

	displayRequestCnt int //現在の左ペインの高さに対して、表示されているリクエストリストの数
}

type FocusPane int

const (
	FocusList     FocusPane = iota //左paneのリストにユニークな連番をあてる
	FocusResponse                  //右pane
)
