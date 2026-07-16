package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/nunokawa/gdeck/cmd/internal/model"
	"github.com/nunokawa/gdeck/cmd/internal/runner"
	"github.com/nunokawa/gdeck/cmd/internal/store"
)

/** 非同期処理を定義するファイル */

// 非同期でrunを実行
// bubbleteaは`tea.Cmd`が返されると内部的にgoroutineを別タスクで実行してくれる
func asyncRunCmd(name string, method string) tea.Cmd {

	return func() tea.Msg {
		results, err := runner.Run(
			name,
			runner.RunOptions{},
		)
		if err != nil {
			return runFinishedMsg{
				err: err,
			}
		}
		if results[0].Error != nil {
			return runFinishedMsg{err: results[0].Error}
		}

		return runFinishedMsg{
			response: results[0].Response,
		}
	}
}

// 非同期でsaveを実行
// 現状は一瞬のため非同期であるメリットは少ないが念の為
func asyncSaveCmd(request *model.Request) tea.Cmd {

	return func() tea.Msg {
		if request.Name == "" || request.Method == "" || request.URL == "" {
			return saveFinishedMsg{err: fmt.Errorf("not enough arguments")}
		}
		err := store.Save(request.Name, request)
		return saveFinishedMsg{name: request.Name, err: err}
	}
}

// 指定した時間だけメッセージが表示される
func clearStatusAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(time.Time) tea.Msg {
		return clearStatusMsg{}
	})
}
