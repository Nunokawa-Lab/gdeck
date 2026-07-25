package cmd

import (
	"fmt"
	"log"

	"github.com/Nunokawa-Lab/gdeck/cmd/internal/store"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gdeck",
	Short: "API Tester CLI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {   // 各コマンドのRunが実行する前に実行されるサブコマンド
		if err := store.EnsureDirs(); err != nil {
			log.Fatal(err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(fmt.Println("Command execution failed：", err.Error()))
	}
}
