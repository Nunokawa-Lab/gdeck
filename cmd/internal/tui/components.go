package tui

import (
	"strings"
	// "unicode/utf8"

	"github.com/charmbracelet/lipgloss"
)

// ペインの上部を文字列計算し生成
func painHeaderLine(title string, width int, isActive bool) string {

	if width <= 2 {
		return "┌" + title + "┐"
	}

	content := " " + title + " "

	// スタイル適用後の実表示幅を計算
	styledContent := content
	if isActive {
		styledContent = activeHeaderTitleStyle.Render(content)
	} else {
		styledContent = inactiveHeaderTitleStyle.Render(content)
	}

	contentWidth := lipgloss.Width(styledContent)
	innerWidth := width - 2 // 角（┌, ┐）を除いた幅

	// title が大きすぎて入らない場合は全体を横線で埋める
	if contentWidth > innerWidth {
		line := "┌" + strings.Repeat("─", innerWidth) + "┐"
		if isActive {
			line = activeColorStyle.Render(line)
		} else {
			line = inactiveColorStyle.Render(line)
		}
		return line
	}

	// 左側の横線は固定で4本（足りない場合はinnerまで）
	leftWidth := 4
	if leftWidth > innerWidth {
		leftWidth = innerWidth
	}

	// 右側の横線は残りで埋める
	rightWidth := innerWidth - leftWidth - contentWidth

	leftContent := "┌" + strings.Repeat("─", leftWidth)
	rightContent := strings.Repeat("─", rightWidth+4) + "┐"

	if isActive {
		leftContent = activeColorStyle.Render(leftContent)
		rightContent = activeColorStyle.Render(rightContent)
	} else {
		leftContent = inactiveColorStyle.Render(leftContent)
		rightContent = inactiveColorStyle.Render(rightContent)
	}

	return leftContent + styledContent + rightContent
}

func painFooterLine(text string, width int, isActive bool) string {

	if width <= 2 {
		return "└" + text + "┘"
	}

	// text が空の場合は全体を横線で埋める
	if text == "" {
		line := "└" + strings.Repeat("─", width+2) + "┘"
		if isActive {
			line = activeColorStyle.Render(line)
		} else {
			line = inactiveColorStyle.Render(line)
		}
		return line
	}

	// text がある場合は真ん中に配置
	content := " " + text + " "

	// スタイル適用後の実表示幅を計算
	styledContent := content
	if isActive {
		styledContent = activeLeftPaneFooterStyle.Render(content)
	} else {
		styledContent = inactiveLeftPaneFooterStyle.Render(content)
	}

	contentWidth := lipgloss.Width(styledContent)
	innerWidth := width - 2 // 角（└, ┘）を除いた幅

	// text が大きすぎて入らない場合は全体を横線で埋める
	if contentWidth > innerWidth {
		line := "└" + strings.Repeat("─", innerWidth) + "┘"
		if isActive {
			line = activeColorStyle.Render(line)
		} else {
			line = inactiveColorStyle.Render(line)
		}
		return line
	}

	// 左右の横線を計算して分割
	separatorWidth := (innerWidth - contentWidth) / 2
	rightSeparatorWidth := innerWidth - contentWidth - separatorWidth

	leftContent := "└" + strings.Repeat("─", separatorWidth+2)
	rightContent := strings.Repeat("─", rightSeparatorWidth+2) + "┘"

	if isActive {
		leftContent = activeColorStyle.Render(leftContent)
		rightContent = activeColorStyle.Render(rightContent)
	} else {
		leftContent = inactiveColorStyle.Render(leftContent)
		rightContent = inactiveColorStyle.Render(rightContent)
	}

	return leftContent + styledContent + rightContent
}

// 検索窓
func (m Model) searchBar() string {

	if !m.searchMode {
		return ""
	}

	input := m.searchInput.View()

	box := searchBoxStyle.Render(input)

	return lipgloss.JoinVertical(lipgloss.Left, box)
}

// フッター
func (m Model) footer() string {

	ft := "↑↓ Move&Scroll   ↵ Run   d Delete   ←→ Focus   / SearchMode   q Quit"
	if m.searchMode {
		ft = "↑↓ Select   ↵ Confirm   esc Cancel"
	}
	if m.deleteConfirm {
		ft = "⚠️  Delete " + m.currentRequest.Name + " ?   y Yes   n No"
		return footerDeleteStyle.Render(ft)
	}

	return footerStyle.Render(ft)
}
