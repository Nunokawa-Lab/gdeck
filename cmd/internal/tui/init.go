package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nunokawa/gdeck/cmd/internal/store"
)

func (m Model) Init() tea.Cmd {
	// 次の回転イベントを返す（予約）
	return m.spinner.Tick
}

func InitialModel() (Model, error) {
	requests, err := store.List()
	if err != nil {
		return Model{}, err
	}

	s := spinner.New()
	s.Spinner = spinner.Dot

	// bubbleteaに渡すinterfaceは Init() Update() View() をレシーバーに持っている必要あり
	m := Model{
		requests: requests,
		cursor:   0,
		spinner:  s,
	}

	m.loadCurrentRequest()

	return m, nil
}
