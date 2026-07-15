package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

/**
BubbleTea(TUI) 専用のスタイルファイル
*/

var (
	// ========== クラシックにするならこれ ==========
	// styleGET    = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Background(lipgloss.Color("24")).Bold(true).Padding(0, 1)  // 青
	// stylePOST   = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Background(lipgloss.Color("28")).Bold(true).Padding(0, 1)  // 緑
	// stylePUT    = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Background(lipgloss.Color("130")).Bold(true).Padding(0, 1) // 黄
	// stylePATCH  = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Background(lipgloss.Color("60")).Bold(true).Padding(0, 1)  // 紫
	// styleDELETE = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Background(lipgloss.Color("88")).Bold(true).Padding(0, 1)  // 赤

	// ========== パステルにするならこれ ==========
	// styleGetPastel    = lipgloss.NewStyle().Foreground(lipgloss.Color("17")).Background(lipgloss.Color("153")).Bold(true).Padding(0, 1) // パステル青
	// stylePostPastel   = lipgloss.NewStyle().Foreground(lipgloss.Color("17")).Background(lipgloss.Color("157")).Bold(true).Padding(0, 1) // パステル緑
	// stylePutPastel    = lipgloss.NewStyle().Foreground(lipgloss.Color("17")).Background(lipgloss.Color("229")).Bold(true).Padding(0, 1) // パステル黄
	// stylePatchPastel  = lipgloss.NewStyle().Foreground(lipgloss.Color("17")).Background(lipgloss.Color("182")).Bold(true).Padding(0, 1) // パステルピンク
	// styleDeletePastel = lipgloss.NewStyle().Foreground(lipgloss.Color("17")).Background(lipgloss.Color("217")).Bold(true).Padding(0, 1) // パステル赤

	// ========== ダーク系ならこれ ==========
	styleGetDark    = lipgloss.NewStyle().Foreground(lipgloss.Color("87")).Background(lipgloss.Color("235")).Bold(true).Padding(0, 1)  // 薄い青
	stylePostDark   = lipgloss.NewStyle().Foreground(lipgloss.Color("114")).Background(lipgloss.Color("235")).Bold(true).Padding(0, 1) // 薄い緑
	stylePutDark    = lipgloss.NewStyle().Foreground(lipgloss.Color("221")).Background(lipgloss.Color("235")).Bold(true).Padding(0, 1) // 薄い黄
	stylePatchDark  = lipgloss.NewStyle().Foreground(lipgloss.Color("183")).Background(lipgloss.Color("235")).Bold(true).Padding(0, 1) // 薄い紫
	styleDeleteDark = lipgloss.NewStyle().Foreground(lipgloss.Color("210")).Background(lipgloss.Color("235")).Bold(true).Padding(0, 1) // 薄い赤

	styleDark = lipgloss.NewStyle().Foreground(lipgloss.Color("7")) // 白/グレー

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("223")).
			PaddingLeft(1).
			MarginTop(1).
			PaddingTop(1).
			MarginBottom(1)

	activeHeaderTitleStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("14")).
				Bold(true)

	inactiveHeaderTitleStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("7")).
					Bold(true)

	activeColorStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("14"))

	inactiveColorStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("8"))

	footerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("99")).
			PaddingLeft(1).
			MarginTop(1).
			PaddingBottom(1)

	footerDeleteStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("5")).
				PaddingLeft(1).
				MarginTop(1).
				PaddingBottom(1)

	errorMsgStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("203"))

	successMsgStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("42"))

	activeLeftPaneStyle = lipgloss.NewStyle().
				Padding(1, 1, 0, 1).
				Border(lipgloss.NormalBorder(), false, true, false, true).
				BorderForeground(lipgloss.Color("14"))

	inactiveLeftPaneStyle = lipgloss.NewStyle().
				Padding(1, 1, 0, 1).
				Border(lipgloss.NormalBorder(), false, true, false, true).
				BorderForeground(lipgloss.Color("8"))

	activeRightPaneStyle = lipgloss.NewStyle().
				Padding(1, 1, 0, 1).
				Border(lipgloss.NormalBorder(), false, true, true, true).
				BorderForeground(lipgloss.Color("14"))

	inactiveRightPaneStyle = lipgloss.NewStyle().
				Padding(1, 1, 0, 1).
				Border(lipgloss.NormalBorder(), false, true, true, true).
				BorderForeground(lipgloss.Color("8"))

	activeLeftPaneFooterStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("14")).
					Bold(true)

	inactiveLeftPaneFooterStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("7")).
					Bold(true)

	searchStyle = lipgloss.NewStyle().Padding(0, 1, 1, 1)

	searchBoxStyle = lipgloss.NewStyle().Padding(0, 1)

	// saveForm: 文字色はターミナルデフォルト（Foreground/Background なし）
	saveFormInputTextStyle = lipgloss.NewStyle()

	saveFormInputPlaceholderStyle = lipgloss.NewStyle().
					Foreground(lipgloss.Color("240"))
)

// charmbracelet/bubbles/textarea が持つスタイル構造体を定義
func saveFormTextareaStyle() textarea.Style {
	text := saveFormInputTextStyle

	return textarea.Style{
		Base:             lipgloss.NewStyle(),
		CursorLine:       text,
		CursorLineNumber: lipgloss.NewStyle(),
		EndOfBuffer:      lipgloss.NewStyle(),
		LineNumber:       saveFormInputPlaceholderStyle,
		Placeholder:      saveFormInputPlaceholderStyle,
		Prompt:           saveFormInputPlaceholderStyle,
		Text:             text,
	}
}

// charmbracelet/bubbles/textinput にスタイルを適用
func applySaveFormTextInputStyle(ti *textinput.Model) {
	ti.TextStyle = saveFormInputTextStyle
	ti.PlaceholderStyle = saveFormInputPlaceholderStyle
	ti.PromptStyle = saveFormInputPlaceholderStyle
	ti.Cursor.Style = saveFormInputTextStyle
	ti.Cursor.TextStyle = saveFormInputTextStyle
}

// charmbracelet/bubbles/textarea にスタイルを適用
func applySaveFormTextareaStyle(ta *textarea.Model) {
	width := ta.Width()
	style := saveFormTextareaStyle()
	ta.FocusedStyle = style
	ta.BlurredStyle = style
	ta.Prompt = "┃"
	ta.SetWidth(width)
}

func methodColor(method string) string {
	switch method {
	case "GET":
		return styleGetDark.Render("GET")
	case "POST":
		return stylePostDark.Render("POST")
	case "PUT":
		return stylePutDark.Render("PUT")
	case "PATCH":
		return stylePatchDark.Render("PATCH")
	case "DELETE":
		return styleDeleteDark.Render("DELETE")
	default:
		return styleDark.Render("Unknown")
	}
}

func padRight(str string, width int) string {
	padding := width - lipgloss.Width(str)

	if padding > 0 {
		str += strings.Repeat(" ", padding)
	}

	return str
}
