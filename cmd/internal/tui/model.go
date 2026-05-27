package tui

import "github.com/nunokawa/gdeck/cmd/internal/model"

type Model struct {
	requests []model.RequestItem
	cursor   int    //現在の選択位置
	selected *model.RequestItem

	response *model.Response
	loading  bool
	errorMsg string
}
