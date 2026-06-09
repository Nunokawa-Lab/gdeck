package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nunokawa/gdeck/cmd/internal/store"
)

func (m Model) Init() tea.Cmd {
	// 次の回転イベントを返す（予約）
	return tea.Batch(
		m.spinner.Tick,
		textinput.Blink,
	)
}

func InitialModel() (Model, error) {

	requests, err := store.List()
	if err != nil {
		return Model{}, err
	}

	// spinner new create
	s := spinner.New()
	s.Spinner = spinner.Dot

	// viewport new create
	vp := viewport.New(80, 20)

	// textinput new create
	ti := textinput.New()
	ti.CharLimit = 100
	ti.Width = 30
	ti.Placeholder = "get-comment"

	// bubbleteaに渡すinterfaceは Init() Update() View() をレシーバーに持っている必要あり
	m := Model{
		requests:      requests,
		cursor:        0,
		spinner:       s,
		leftViewport:  vp,
		rightViewport: vp,
		searchInput:   ti,
	}

	m.loadCurrentRequest()

	m.leftViewport.SetContent(m.requestListContent())
	m.rightViewport.SetContent(m.responseContent())

	return m, nil
}
