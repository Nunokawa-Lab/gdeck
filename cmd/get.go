package cmd

import (
	"apictl/cmd/internal/httpclient"
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
var getCmd = &cobra.Command{
	Use: "get [url]",
	Short: "Send get request",
	Example: "apictl get https://example.com",
	Args: cobra.ExactArgs(1),
	Run: func (cmd *cobra.Command, args []string)  {
		url := args[0]
		res, err := httpclient.Get(url)
		if err != nil {
			fmt.Println("Error : ", err.Error())
			return
		}

		fmt.Println(res)
	},
}

func init() {
	// Getコマンドを登録
	rootCmd.AddCommand(getCmd)
}