package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
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

	spinner spinner.Model
}
