package tui

import (
	"fmt"
	"time"

	"github.com/Nunokawa-Lab/gdeck/cmd/internal/model"
	"github.com/Nunokawa-Lab/gdeck/cmd/internal/runner"
	"github.com/Nunokawa-Lab/gdeck/cmd/internal/store"
	tea "github.com/charmbracelet/bubbletea"
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
		if len(results) == 0 {
			return runFinishedMsg{
				err: fmt.Errorf("no results"),
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

// 非同期でsaveコマンドを実行
// save: Nameが同名なら更新、別名なら新規作成
// edit: Nameが同名でも別名でも更新、ただし別名の場合の更新は「新規作成＋旧ファイル削除」で更新とする
func asyncSaveCmd(request *model.Request, originalName string) tea.Cmd {

	return func() tea.Msg {
		if request.Name == "" || request.Method == "" || request.URL == "" {
			return saveFinishedMsg{err: fmt.Errorf("not enough arguments")}
		}

		saveName := request.Name
		renamed := false

		if originalName != "" {
			if requestNamesEqual(originalName, request.Name) {
				saveName = originalName
			} else {
				renamed = true
			}
		}

		updated, err := store.Save(saveName, request)
		if err != nil {
			return saveFinishedMsg{name: request.Name, err: err}
		}

		if renamed {
			if err := store.Delete(originalName); err != nil {
				return saveFinishedMsg{name: request.Name, err: err}
			}
			updated = true
		} else if originalName != "" {
			updated = true
		}

		return saveFinishedMsg{name: request.Name, updated: updated, err: err}
	}
}

// 指定した時間だけメッセージが表示される
func clearStatusAfter(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(time.Time) tea.Msg {
		return clearStatusMsg{}
	})
}
