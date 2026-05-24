package tui

import tea "github.com/charmbracelet/bubbletea"

func Start() error {

	// bubbleteaに渡すinterfaceは Init() Update() View() をレシーバーに持っている必要あり
	p := tea.NewProgram(Model{})

	_, err := p.Run()

	return err
}