package cmd

import (
	"fmt"
	"strings"

	"github.com/nunokawa/gdeck/cmd/internal/model"
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
 *   cmd: コマンドの情報（フラグ取得などに使う）
 *   args: CLIで渡された引数（今回はURLが入る）
 */
var saveCmd = &cobra.Command{
	Use:     "save [({path/})name] [method] [url]",
	Short:   "Save request command for later use",
	Example: "gdeck save SampleCmd get https://example.com",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
		method := args[1]
		url := args[2]

		req := &model.Request{
			Name:    name,
			Method:  strings.ToUpper(method),
			URL:     url,
			Headers: requestHeaders,
			Body:    requestData,
		}

		err := store.Save(name, req)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("Saved: ", name)
	},
}

func init() {
	// Flagsで引数・オプションを設定
	// -d
	saveCmd.Flags().StringVarP(&requestData, "data", "d", "", "Request body")
	// -H
	saveCmd.Flags().StringArrayVarP(&requestHeaders, "header", "H", []string{}, "Request header")

	// 登録
	rootCmd.AddCommand(saveCmd)
}
