package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nunokawa/gdeck/cmd/internal/store"
)

func Start() error {

	requests, err := store.List()
	if err != nil {
		return err
	}

	// bubbleteaに渡すinterfaceは Init() Update() View() をレシーバーに持っている必要あり
	p := tea.NewProgram(
		Model{
			requests: requests,
		},
	)

	_, err = p.Run()

	return err
}
