package cmd

import (
	"apictl/cmd/internal/httpclient"
	outputHandler "apictl/cmd/internal/output"
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
var runCmd = &cobra.Command{
	Use:   "run [name]",
	Short: "Run saved request",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		// 読み込み
		req, err := store.Load(name)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		// 実行
		res, err := httpclient.Do(
			req.Method,
			req.URL,
			req.Body,
			req.Headers,
		)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		outputHandler.PrintResponse(res, isVerbose) // ← verboseは後でフラグ対応
	},
}

func init() {
	// -v
	runCmd.Flags().BoolVarP(&isVerbose, "verbose", "v", false, "Verbose output")

	// rootに登録
	rootCmd.AddCommand(runCmd)
}
