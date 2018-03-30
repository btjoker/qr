package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

var (
	outFile  *string
	size     *int
	textArt  *bool
	negative *bool
)

var encodeCMD = &cobra.Command{
	Use:   "encode",
	Short: "生成二维码",
	Long:  "更具输入生成二维码图片",
	Run: func(cmd *cobra.Command, args []string) {
		var q *qrcode.QRCode
		var err error
		content := strings.Join(args, " ")
		q, err = qrcode.New(content, qrcode.Highest)
		if err != nil {
			log.Fatalln(err)
		}

		if *textArt {
			art := q.ToString(*negative)
			fmt.Println(art)
			return
		}

		if *negative {
			q.ForegroundColor, q.BackgroundColor = q.BackgroundColor, q.ForegroundColor
		}

		var png []byte
		png, err = q.PNG(*size)
		if err != nil {
			log.Fatalln(err)
		}

		if *outFile == "" {
			os.Stdout.Write(png)
		} else {
			var fn *os.File
			fn, err = os.Create(*outFile + ".png")
			if err != nil {
				log.Fatalln(err)
			}
			defer fn.Close()
			fn.Write(png)
		}

	},
}

func init() {
	outFile = encodeCMD.PersistentFlags().String("o", "", "输出文件名")
	size = encodeCMD.PersistentFlags().Int("s", 256, "图片大小 像素点")
	textArt = encodeCMD.PersistentFlags().Bool("t", false, "输出到stdout")
	negative = encodeCMD.PersistentFlags().Bool("i", false, "反转黑白")

	rootCMD.AddCommand(encodeCMD)
}
