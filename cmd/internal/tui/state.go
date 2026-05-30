package tui

import (
	"github.com/nunokawa/gdeck/cmd/internal/store"
)

func (m *Model) loadCurrentRequest() {
	forcusedRequest := m.requests[m.cursor]

	reqs, err := store.Load(forcusedRequest.Name)

	if err != nil {
		m.errorMsg = err.Error()
		return
	}

	if len(reqs) == 0 {
		return
	}

	m.currentRequest = reqs[0]
}
