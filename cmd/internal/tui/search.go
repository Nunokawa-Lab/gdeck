package tui

import (
	"strings"

	"github.com/nunokawa/gdeck/cmd/internal/model"
)

func (m *Model) applySearch(text string) {
	if text == "" {
		m.filteredRequests = m.requests
		return
	}

	var filtered []model.RequestItem

	for _, req := range m.requests {
		if strings.Contains(
			strings.ToLower(req.Name),
			strings.ToLower(text),
		) {
			filtered = append(filtered, req)
		}
	}

	m.filteredRequests = filtered
}

func (m *Model) resetSearch() {
	m.searchMode = false

	m.searchInput.Blur()
	m.searchInput.SetValue("")

	m.filteredRequests = m.requests
}
