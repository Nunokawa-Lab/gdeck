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
	styleGET    = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Background(lipgloss.Color("24")).Bold(true).Padding(0, 1)  // 青
	stylePOST   = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Background(lipgloss.Color("28")).Bold(true).Padding(0, 1)  // 緑
	stylePUT    = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Background(lipgloss.Color("130")).Bold(true).Padding(0, 1) // 黄
	stylePATCH  = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Background(lipgloss.Color("60")).Bold(true).Padding(0, 1)  // 紫
	styleDELETE = lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Background(lipgloss.Color("88")).Bold(true).Padding(0, 1)  // 赤
	styleDark   = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))                                                              // 白/グレー

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("223")).
			PaddingLeft(1).
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
			MarginTop(1)

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
