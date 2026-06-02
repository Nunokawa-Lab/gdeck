package tui

import (
	"strings"
	"unicode/utf8"
)

// ペインの上部を文字列計算し生成
func headerLine(title string, width int, isActive bool) string {
	// 罫線等特殊文字の影響か、3文字分落ちるため最初に帳尻合わせる
	width += 3

	if width <= 2 {
		return "┌" + title + "┐"
	}

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
		return "┌" + strings.Repeat("─", inner) + "┐"
	}

	// 右側の横線を残りで埋める
	right := inner - left - cl
	if right < 0 {
		right = 0
	}

	leftContent := "┌" + strings.Repeat("─", left)
	rightContent := strings.Repeat("─", right) + "┐"
	if isActive {
		leftContent = activeHeaderStyle.Render(leftContent)
		rightContent = activeHeaderStyle.Render(rightContent)
		content = activeHeaderTitleStyle.Render(content)
	} else {
		leftContent = inactiveHeaderStyle.Render(leftContent)
		rightContent = inactiveHeaderStyle.Render(rightContent)
		content = inactiveHeaderTitleStyle.Render(content)
	}

	return leftContent + content + rightContent
}