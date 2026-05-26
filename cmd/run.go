package cmd

import (
	"fmt"

	outputHandler "github.com/nunokawa/gdeck/cmd/internal/output"
	"github.com/nunokawa/gdeck/cmd/internal/runner"

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
	Use:     "run [({path/})name]",
	Short:   "Run saved request",
	Example: "gdeck run SampleCmd --env dev",
	Args:    cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {

		name := args[0]

		results, err := runner.Run(
			name,
			runner.RunOptions{
				Timeout: timeout,
				EnvName: envName,
				Body:    requestData,
				Headers: requestHeaders,
			},
		)

		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		total := len(results)

		for i, result := range results {

			if result.Error != nil {
				fmt.Println("Error:", result.Error)
				continue
			}

			outputHandler.PrintResponse(
				result.Response,
				isVerbose,
				i+1,
				total,
				result.Request.Method,
				result.Request.Name,
				result.Request.URL,
			)
		}
	},
}

func init() {
	// -v
	runCmd.Flags().BoolVarP(&isVerbose, "verbose", "v", false, "Verbose output")
	// -d
	runCmd.Flags().StringVarP(&requestData, "data", "d", "", "Request body")
	// -H
	runCmd.Flags().StringArrayVarP(&requestHeaders, "header", "H", []string{}, "Request header")
	// -t
	runCmd.Flags().IntVarP(&timeout, "timeout", "t", 10, "timeout seconds")
	// --env
	runCmd.Flags().StringVar(&envName, "env", "", "environment name")

	// rootに登録
	rootCmd.AddCommand(runCmd)
}
