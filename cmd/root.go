package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "qr",
	Short: "二维码生成与解析",
	Long:  "将一个生成二维码库与解码二维码库的组合的粗糙命令行",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

// Execute .
func Execute() {
	if err := rootCMD.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
