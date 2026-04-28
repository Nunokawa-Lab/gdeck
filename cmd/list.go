package cmd

import (
	"apictl/cmd/internal/store"
	"fmt"

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
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Show Saved-Command List",
	Example: "apictl list",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		filenames := store.List()

		for _, name := range filenames {
			fmt.Println(name)
		}
	},
}

func init() {
	// 登録
	rootCmd.AddCommand(listCmd)
}
