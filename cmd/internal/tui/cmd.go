package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nunokawa/gdeck/cmd/internal/runner"
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

		return runFinishedMsg{
			response: results[0].Response,
		}
	}
}
