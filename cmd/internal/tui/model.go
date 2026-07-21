package tui

import (
	"github.com/Nunokawa-Lab/gdeck/cmd/internal/model"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
)

type Model struct {
	requests       []model.RequestItem //リクエストアイテムリスト
	cursor         int                 //現在のカーソルの位置
	selected       *model.RequestItem  //選択されたリクエストアイテム
	currentRequest *model.Request      //カーソルが当たっているリクエスト情報

	response      *model.Response //実行されたリクエストレスポンス
	rightPaneView RightPaneView   // 右ペインの表示状態（表示判定には使用しないようにする）
	loading       bool            // rightPaneView と同期して更新する
	errorMsg      string
	statusMsg     string // フッター表示用の一時メッセージ（成功など）

	spinner       spinner.Model
	leftViewport  viewport.Model
	rightViewport viewport.Model

	focus FocusPane

	leftPaneWidth  int //左paneの幅
	rightPaneWidth int //右paneの幅
	paneHeight     int

	displayRequestCnt int //現在の左ペインの高さに対して、表示されているリクエストリストの数

	filteredRequests []model.RequestItem

	searchInput textinput.Model

	saveForm           saveForm
	saveFormFieldCount int
	editingRequestName string // 空なら新規保存、非空なら edit 中の元リスト名（相対パス）

	mode Mode
}

type FocusPane int

const (
	FocusList     FocusPane = iota //左paneのリストにユニークな連番をあてる
	FocusResponse                  //右pane
)

type Mode int

const (
	ModeNormal Mode = iota
	ModeSearch
	ModeDeleteConfirm
	ModeSave
)

type RightPaneView int

const (
	RightPanePreview RightPaneView = iota
	RightPaneLoading
	RightPaneResponse
	// リクエストの新規作成も右ペイン領域を使用するが Mode で管理する
)

type SaveFocusFiled int

const (
	focusSaveFieldName SaveFocusFiled = iota
	focusSaveFieldMethod
	focusSaveFieldURL
	focusSaveFieldHeader
	focusSaveFieldBody
)

// saveコマンド用フォーム
type saveForm struct {
	name   textinput.Model
	method textinput.Model
	url    textinput.Model
	header textarea.Model
	body   textarea.Model
	focus  SaveFocusFiled // 縦に並んだtextinput毎のフォーカス
}

// 比較や計算する時に型不整合が起きない用のInt変換
func (sf *saveForm) toIntFocus() int {
	return int(sf.focus)
}
