package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
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

	vp := viewport.New(80, 20)

	// bubbleteaに渡すinterfaceは Init() Update() View() をレシーバーに持っている必要あり
	m := Model{
		requests: requests,
		cursor:   0,
		spinner:  s,
		leftViewport: vp,
		viewport: vp,
	}

	m.loadCurrentRequest()

	m.leftViewport.SetContent(m.requestListContent())
	m.viewport.SetContent(m.responseContent())

	return m, nil
}
