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

	app.Name = "RosterAutomaticDeliverySystem-Cli"
	app.Usage = "RosterAutomaticDeliverySystem-Cli"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		year := context.String("y")
		month := context.String("m")
		driveName := context.String("f")
		host := context.String("h")

		if driveName == "" {
			fmt.Printf("-f をつけて実行してください")
			return nil
		}
		if year == "" {
			fmt.Printf("-y をつけて実行してください")
			return nil
		}
		if month == "" {
			fmt.Printf("-m をつけて実行してください")
			return nil
		}
		if host == "" {
			fmt.Printf("-h をつけて実行してください")
			return nil
		}

		c, _ := web.NewClient(host)
		h, _ := handler.NewHandler(c)

		err := h.CopyFile(driveName, year,month)
		if err != nil {
			return err
		}
		return nil
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "year, y",
			Usage: "コピーしたい年を設定する -y 2020",
		},
		cli.StringFlag{
			Name:  "month, m",
			Usage: "コピーしたい月を設定する -m 01",
		},
		cli.StringFlag{
			Name:  "folder, f",
			Usage: "コピーしたい共有ドライブを設定する -f 勤務表",
		},
		cli.StringFlag{
			Name:  "host, h",
			Usage: "RADSのAPIサーバーを指定する -h http://localhost:1323",
		},
	}

	app.Run(os.Args)
}
