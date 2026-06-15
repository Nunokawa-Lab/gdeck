package tui

import (
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
)

// ペインの上部を文字列計算し生成
func painHeaderLine(title string, width int, isActive bool) string {

	if width <= 2 {
		return "┌" + title + "┐"
	}

	// 罫線等特殊文字の影響か、3文字分落ちるため最初に帳尻合わせる
	width += 3

	// 角を引いた長さ
	inner := width - 2

	// 左側の横線は固定で8本（足りない場合はinnerまで）
	left := 4
	if left > inner {
		left = inner
	}

	content := " " + title + " "
	cl := utf8.RuneCountInString(content) // runeで長さを数える

	// タイトルが入らない場合はタイトルを消して全体を横線で埋める
	if cl > inner-left {
		line := "┌" + strings.Repeat("─", inner+1) + "┐"
		if isActive {
			line = activeColorStyle.Render(line)
		} else {
			line = inactiveColorStyle.Render(line)
		}
		return line
	}

	// 右側の横線を残りで埋める
	right := inner - left - cl
	if right < 0 {
		right = 0
	}

	leftContent := "┌" + strings.Repeat("─", left)
	rightContent := strings.Repeat("─", right) + "┐"
	if isActive {
		leftContent = activeColorStyle.Render(leftContent)
		rightContent = activeColorStyle.Render(rightContent)
		content = activeHeaderTitleStyle.Render(content)
	} else {
		leftContent = inactiveColorStyle.Render(leftContent)
		rightContent = inactiveColorStyle.Render(rightContent)
		content = inactiveHeaderTitleStyle.Render(content)
	}

	return leftContent + content + rightContent
}

func painFooterLine(text string, width int, isActive bool) string {

	if width <= 2 {
		return "└" + text + "┘"
	}

	// 罫線等特殊文字の影響か、3文字分落ちるため最初に帳尻合わせる
	width += 3
	if remainder := width % 2; remainder == 1 {
		// 3足して2で割ると必ずあまりは0か1
		// 余りが出る時は1落とされるため帳尻合わせのため+1
		width += 1
	}

	// 角を引いた長さ
	inner := width - 2

	content := " " + text + " "
	cl := utf8.RuneCountInString(content) // runeで長さを数える

	// テキストを真ん中に表示させるために半分で計算
	separator := (inner - cl) / 2

	// タイトルが入らない場合はタイトルを消して全体を横線で埋める
	if cl > inner-separator {
		line := "└" + strings.Repeat("─", inner-1) + "┘"
		if isActive {
			line = activeColorStyle.Render(line)
		} else {
			line = inactiveColorStyle.Render(line)
		}
		return line
	}

	leftContent := "└" + strings.Repeat("─", separator)
	rightContent := strings.Repeat("─", separator) + "┘"

	if isActive {
		leftContent = activeColorStyle.Render(leftContent)
		rightContent = activeColorStyle.Render(rightContent)
		content = activeLeftPaneFooterStyle.Render(content)
	} else {
		leftContent = inactiveColorStyle.Render(leftContent)
		rightContent = inactiveColorStyle.Render(rightContent)
		content = inactiveLeftPaneFooterStyle.Render(content)
	}

	return leftContent + content + rightContent
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

	ft := "↑↓ Move&Scroll   ↵ Run   ←→ Focus   / SearchMode   q Quit"
	if m.searchMode {
		ft = "↑↓ Select   ↵ Confirm   esc Cancel"
	}

	return footerStyle.Render(ft)
}
