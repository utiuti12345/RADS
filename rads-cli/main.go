package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
	"rads-cli/handler"
	"rads-cli/infra/web"
)

func main() {
	app := cli.NewApp()

	app.Name = "RosterAutmaticDeliverySystem-Cli"
	app.Usage = "RosterAutmaticDeliverySystem-Cli"
	app.Version = "0.0.1"

	c, _ := web.NewClient("http://localhost:1323")
	h, _ := handler.NewHandler(c)

	app.Action = func(context *cli.Context) error {
		driveName := context.String("f")
		dateTime := context.String("d")

		if driveName == "" {
			fmt.Printf("-f をつけて実行してください")
			return nil
		}
		if dateTime == "" {
			fmt.Printf("-d をつけて実行してください")
			return nil
		}

		h.CopyFile(driveName, dateTime)

		return nil
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "datetime, d",
			Usage: "コピーしたい年と月を設定する -d 202001",
		},
		cli.StringFlag{
			Name:  "folder, f",
			Usage: "コピーしたい共有ドライブを設定する -f 勤務表",
		},
	}

	app.Run(os.Args)
}
