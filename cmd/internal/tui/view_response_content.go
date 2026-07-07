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

	switch m.rightPaneView {

	case RightPaneLoading:
		return fmt.Sprintf(
			"\n%s Sending Request...",
			m.spinner.View(),
		)

	case RightPaneResponse:
		if m.response == nil {
			// フォールバック（通常は来ない）
			return fmt.Sprintln("Listening for your signals 📡✨")
		}
		return output.RenderTUIResponse(
			m.response,
			m.selected.Method,
			m.rightViewport.Width,
			m.focus == FocusResponse,
		)

	default: // RightPanePreview
		if m.currentRequest == nil {
			return fmt.Sprintln("Listening for your signals 📡✨")
		}
		return output.RenderTUIPreview(m.currentRequest)
	}
}
