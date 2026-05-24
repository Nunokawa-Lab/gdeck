package cmd

import (
	"fmt"

	"github.com/nunokawa/gdeck/cmd/internal/tui"
	"github.com/spf13/cobra"
)

/**
 * Use:
 *   コマンドの使い方（CLI上での呼び出し形式）
 * Short:
 *   コマンドの簡単な説明（1行）
 *   -- help
 * Example:
 *   使用例
 * Args:
 *   引数のバリデーションルールを定義する
 *   ExactArgs: 設定した引数の数(今回は1)より少なかったらエラー
 * Run:
 *   コマンド実行時に呼ばれる処理本体
 *   cmd: コマンドの情報
 *   args: CLIで渡された引数
 */
var tuiCmd = &cobra.Command{
	Use: "tui",
	Short: "Start gdeck TUI",
	Example: "gdeck tui",
	Run: func(cmd *cobra.Command, args []string) {
		
		err := tui.Start()
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
		
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}