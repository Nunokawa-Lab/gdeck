package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nunokawa/gdeck/cmd/internal/store"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func InitialModel() (Model, error) {
	requests, err := store.List()
	if err != nil {
		return Model{}, err
	}

	// bubbleteaに渡すinterfaceは Init() Update() View() をレシーバーに持っている必要あり
	m := Model{
		requests: requests,
		cursor:   0,
	}

	m.loadCurrentRequest()

	return m, nil
}
