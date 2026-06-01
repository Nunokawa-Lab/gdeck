package tui

import (
	"fmt"

	"github.com/nunokawa/gdeck/cmd/internal/output"
)

/*
*

	出力テキストを返す
	 ┗ロード
	 ┗プレビュー
	 ┗レスポンス
*/
func (m Model) responseContent() string {

	if m.loading {

		return fmt.Sprintf(
			"\n%s Sending Request...",
			m.spinner.View(),
		)
	}

	if m.response == nil {

		return output.RenderTUIPreview(
			m.currentRequest,
		)
	}

	return output.RenderTUIResponse(
		m.response,
		m.selected.Method,
	)
}
