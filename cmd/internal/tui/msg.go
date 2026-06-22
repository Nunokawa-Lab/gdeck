package tui

import "github.com/nunokawa/gdeck/cmd/internal/model"

/** 非同期処理の結果を格納する構造体を定義するファイル */

type runFinishedMsg struct {
	response *model.Response
	err      error
}
