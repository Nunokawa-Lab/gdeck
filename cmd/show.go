package cmd

import (
	"encoding/json"
	"fmt"

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
var showCmd = &cobra.Command{
	Use:     "show [({path/})name]",
	Short:   "Show Saved-Command Detail",
	Example: "gdeck show TestCommand",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		requests, err := store.Load(name)
		if err != nil {
			fmt.Println(err.Error())
		}

		reqLen := len(requests)

		for i, req := range requests {
			reqJSON, err := json.MarshalIndent(req, "", "  ")
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Printf("[%v/%v]\n---------------------------\n", i+1, reqLen)
			fmt.Println(string(reqJSON) + "\n")
		}
	},
}

func init() {
	// 登録
	rootCmd.AddCommand(showCmd)
}
