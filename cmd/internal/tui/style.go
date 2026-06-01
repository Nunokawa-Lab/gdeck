package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

/**
BubbleTea(TUI) 専用のスタイルファイル
*/

var (
	// lipgloss
	styleGET    = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true) // 青
	stylePOST   = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true) // 緑
	stylePUT    = lipgloss.NewStyle().Foreground(lipgloss.Color("11")).Bold(true) // 黄
	stylePATCH  = lipgloss.NewStyle().Foreground(lipgloss.Color("13")).Bold(true) // 紫
	styleDELETE = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)  // 赤
	styleDark   = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))             // 白/グレー

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("7")).
			PaddingLeft(1).
			MarginBottom(1)

	headerStyle = lipgloss.NewStyle().
			Bold(true).
			BorderForeground(lipgloss.Color("8")).
			Foreground(lipgloss.Color("7")).
			PaddingLeft(1)

	footerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("99")).
			PaddingLeft(1).
			MarginTop(1)

	activePaneStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("14"))

	inactivePaneStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("8"))
)

func methodColor(method string) string {
	switch method {
	case "GET":
		return styleGET.Render("GET")
	case "POST":
		return stylePOST.Render("POST")
	case "PUT":
		return stylePUT.Render("PUT")
	case "PATCH":
		return stylePATCH.Render("PATCH")
	case "DELETE":
		return styleDELETE.Render("DELETE")
	default:
		return styleDark.Render("Unknown Method")
	}
}

func padRight(str string, width int) string {
	padding := width - lipgloss.Width(str)

	if padding > 0 {
		str += strings.Repeat(" ", padding)
	}

	return str
}
