package cmd

import (
	"apictl/cmd/internal"
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
 *   cmd: コマンドの情報（フラグ取得などに使う）
 *   args: CLIで渡された引数（今回はURLが入る）
 */
var deleteCmd = &cobra.Command{
	Use:     "delete [({path/})name]",
	Short:   "Delete Saved command",
	Example: "apictl delete getUser",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		err := store.Delete(name)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Status Code:  " + internal.ColorStatus("204 No-Content", 204))
	},
}

func init() {
	// Deleteコマンドを登録
	rootCmd.AddCommand(deleteCmd)
}
