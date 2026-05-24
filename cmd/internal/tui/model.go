package tui

type Model struct{
	requests []string //一覧
	cursor int //現在の選択位置
	selected string //選択されたリクエスト
}
