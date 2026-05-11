package cmd

import (
	"apictl/cmd/internal/env"
	"apictl/cmd/internal/httpclient"
	outputHandler "apictl/cmd/internal/output"
	"apictl/cmd/internal/request"
	"apictl/cmd/internal/store"
	"context"
	"errors"
	"fmt"
	"os"

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
	Use:   "run [({path/})name]",
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

		// Body上書き
		if requestData != "" {
			req.Body = requestData
		}

		// Header上書き
		if len(requestHeaders) > 0 {
			req.Headers = request.MergeHeaders(req.Headers, requestHeaders)
		}

		// 環境変数置換
		req.URL, err = env.ReplaceEnv(req.URL)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		req.Body, err = env.ReplaceEnv(req.Body)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for i, h := range req.Headers {
			req.Headers[i], err = env.ReplaceEnv(h)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		}

		// オプション設定
		options := store.DefaultOptions()
		if timeout != 0 {
			options.Timeout = timeout
		}

		// 実行
		res, err := httpclient.Do(
			req.Method,
			req.URL,
			req.Body,
			req.Headers,
			options,
		)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) || os.IsTimeout(err) {
				fmt.Println("Error: Request timed out")
				return
			}

			fmt.Println("Error:", err)
			return
		}

		outputHandler.PrintResponse(res, isVerbose) // ← verboseは後でフラグ対応
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

	// rootに登録
	rootCmd.AddCommand(runCmd)
}
