package tui

import "github.com/nunokawa/gdeck/cmd/internal/model"

type runFinishedMsg struct {
	response *model.Response
	err      error
}
