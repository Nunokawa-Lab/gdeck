package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func Start() error {

	model, err := InitialModel()
	if err != nil {
		return err
	}

	p := tea.NewProgram(model)

	_, err = p.Run()

	return err
}
