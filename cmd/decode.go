package cmd

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.com/tuotoo/qrcode"
)

var (
	remoteIMG  *bool
	outToLocal *string
)

var decodeCMD = &cobra.Command{
	Use:   "decode",
	Short: "解析二维码图片",
	Long:  "接收一个path/URL, 解析该文件二维码内容",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			return
		}

		path := args[0]

		data := decode(path)

		if *outToLocal != "" {
			fn, err := os.Create(*outToLocal + ".txt")
			if err != nil {
				log.Fatalln(err)
			}
			defer fn.Close()
			fn.WriteString(data)
		} else {
			fmt.Println(data)
		}
	},
}

func exist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func decode(path string) string {
	var (
		qc  *qrcode.Matrix
		err error
	)

	if *remoteIMG {
		if _, err = url.Parse(path); err != nil {
			log.Fatalln(err)
		}

		resp, err := http.Get(path)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		qc, err = qrcode.Decode(resp.Body)
	} else {
		if exist(path) {
			file, err := os.Open(path)
			if err != nil {
				log.Fatalln(err)
			}
			defer file.Close()

			qc, err = qrcode.Decode(file)
		}
	}
	return qc.Content
}

func init() {
	remoteIMG = decodeCMD.PersistentFlags().Bool("n", false, "网络图片")
	outToLocal = decodeCMD.PersistentFlags().String("o", "", "输出到本地的文件名")
	rootCMD.AddCommand(decodeCMD)
}
