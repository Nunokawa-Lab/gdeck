package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gdeck",
	Short: "API Tester CLI",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(fmt.Println("コマンド起動エラー：", err.Error()))
	}
}
