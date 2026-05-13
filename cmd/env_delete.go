package cmd

import (
	"apictl/cmd/internal/env"
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
 *   cmd: コマンドの情報（フラグ取得などに使う）
 *   args: CLIで渡された引数（今回はURLが入る）
 */
var envDeleteCmd = &cobra.Command{
	Use:   "delete KEY",
	Short: "Delete environment variable",
	Example: `
		apictl env delete TOKEN
		apictl env delete BASE_URL
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		key := args[0]

		err := env.Delete(key)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Deleted: %s\n", key)
	},
}

func init() {
	envCmd.AddCommand(envDeleteCmd)
}
