package output

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// JSONを整形する関数
func FormatJSON(data []byte) string {
	// 書き込み先のバッファ作成
	var out bytes.Buffer

	// json.Indent(書き込み先, 元データ, インデント開始, インデント文字)
	err := json.Indent(&out, data, "", "  ")
	if err != nil {
		// JSON以外だった場合は文字列として返す
		return string(data)
	}

	return out.String()
}

// ステータスコード色付け
func ColorStatus(status string, code int) string {
	switch {
	case code >= 200 && code < 300:
		return "\033[42;30m" + status + "\033[0m" // 緑背景＋黒文字
	case code >= 300 && code < 400:
		return "\033[44;37m" + status + "\033[0m" // 青背景＋白文字
	case code >= 400 && code < 500:
		return "\033[43;30m" + status + "\033[0m" // 黄背景＋黒文字
	case code >= 500:
		return "\033[41;37m" + status + "\033[0m" // 赤背景＋白文字
	default:
		return status
	}
}

func SelectStatusIcon(code int) string {
	switch {
	case code >= 200 && code < 300:
		return "✅" // success
	case code >= 300 && code < 500:
		return "⚠️ " // warning（半角空白は必要なため消さない）
	case code >= 500:
		return "❌" // error
	default:
		return ""
	}
}

func MethodIcon(method string) string {
	switch method {
	case "GET":
		return "🔵 GET"
	case "POST":
		return "🟢 POST"
	case "PUT":
		return "🟡 PUT"
	case "PATCH":
		return "🟣 PATCH"
	case "DELETE":
		return "🔴 DELETE"
	default:
		return "⚪ Unknown Method"
	}
}

var (
	styleGET    = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true) // 青
	stylePOST   = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true) // 緑
	stylePUT    = lipgloss.NewStyle().Foreground(lipgloss.Color("11")).Bold(true) // 黄
	stylePATCH  = lipgloss.NewStyle().Foreground(lipgloss.Color("13")).Bold(true) // 紫
	styleDELETE = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)  // 赤
	styleDark   = lipgloss.NewStyle().Foreground(lipgloss.Color("7"))             // 白/グレー
)

func MethodColor(method string) string {
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

func PadRight(str string, width int) string {
	padding := width - lipgloss.Width(str)

	if padding > 0 {
		str += strings.Repeat(" ", padding)
	}

	return str
}
