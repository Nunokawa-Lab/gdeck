package cmd

import (
	"fmt"
	"path/filepath"

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
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Show Saved-Command List",
	Example: "gdeck list",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		requestItems, err := store.List()
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}

		for _, reqeustItem := range requestItems {
			name := reqeustItem.Name
			ext := filepath.Ext(name)
			cmdName := name[:len(name)-len(ext)]

			// 自動生成ファイルが出来上がっている場合もあるため、特定の拡張子以外は出力対象外
			if ext != ".json" && ext != ".txt" {
				continue
			}
			fmt.Println(cmdName)
		}
	},
}

func init() {
	// 登録
	rootCmd.AddCommand(listCmd)
}
