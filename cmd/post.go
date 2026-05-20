package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/nunokawa/gdeck/cmd/internal/httpclient"
	outputHandler "github.com/nunokawa/gdeck/cmd/internal/output"
	"github.com/nunokawa/gdeck/cmd/internal/store"

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
var postCmd = &cobra.Command{
	Use:     "post [url]",
	Short:   "Send POST request",
	Example: "gdeck post <URL> -d '<JSON>' -H 'Content-Type: application/json'",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		url := args[0]

		// オプション設定
		options := store.DefaultOptions()
		if timeout != 0 {
			options.Timeout = timeout
		}

		res, err := httpclient.Post(url, requestData, requestHeaders, options)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) || os.IsTimeout(err) {
				fmt.Println("Error:, Request timed out")
				return
			}

			fmt.Println("Error: ", err.Error())
			return
		}

		if output != "" {
			outputHandler.WriteFile(res, output, isVerbose)
		} else {
			outputHandler.PrintResponse(res, isVerbose)
		}
	},
}

func init() {
	// Flagsで引数・オプションを設定
	// -d
	postCmd.Flags().StringVarP(&requestData, "data", "d", "", "Request body")
	// -H
	postCmd.Flags().StringArrayVarP(&requestHeaders, "header", "H", []string{}, "Request header")
	// -v
	postCmd.Flags().BoolVarP(&isVerbose, "verbose", "v", false, "Verbose output")
	// -o
	postCmd.Flags().StringVarP(&output, "output", "o", "", "Output file path")
	// -t
	postCmd.Flags().IntVarP(&timeout, "timeout", "t", 10, "timeout seconds")

	// 登録
	rootCmd.AddCommand(postCmd)
}
