package cmd

import (
	"fmt"

	"github.com/nunokawa/gdeck/cmd/internal/env"

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
var envSetCmd = &cobra.Command{
	Use:   "set KEY VALUE",
	Short: "Set environment variable",
	Example: `
		gdeck env set TOKEN abc123
		gdeck env set BASE_URL https://api.example.com
	`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		key := args[0]
		value := args[1]

		err := env.Set(key, value)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Printf("Saved: %s\n", key)
	},
}

func init() {
	envCmd.AddCommand(envSetCmd)
}
