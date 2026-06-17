package tui

import (
	"fmt"

	"github.com/nunokawa/gdeck/cmd/internal/output"
)

// 出力テキストを返す
//
//	┗ロード
//	┗プレビュー
//	┗レスポンス
func (m Model) responseContent() string {

	if m.loading {

		return fmt.Sprintf(
			"\n%s Sending Request...",
			m.spinner.View(),
		)
	}

	if m.currentRequest == nil {
		// 検索後0件ヒットだった場合に入る
		return fmt.Sprintln("Listening for your signals 📡✨")
	}

	if m.response == nil {

		return output.RenderTUIPreview(
			m.currentRequest,
		)
	}

	return output.RenderTUIResponse(
		m.response,
		m.selected.Method,
		m.rightViewport.Width,
		m.focus == FocusResponse,
	)
}
