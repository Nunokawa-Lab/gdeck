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
var envListCmd = &cobra.Command{
	Use:   "list",
	Short: "List environment variables",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		envs, err := env.LoadEnv()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if len(envs) == 0 {
			fmt.Println("No environment variables")
			return
		}

		for key, value := range envs {
			fmt.Printf("%s=%s\n", key, value)
		}
	},
}

func init() {
	envCmd.AddCommand(envListCmd)
}
