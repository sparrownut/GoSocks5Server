package main

import (
	"GoSocks5Server/Global"
	"GoSocks5Server/utils/systemutils"
	"fmt"
	"github.com/armon/go-socks5"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name:      "GoSocks5Server",
		Usage:     "Socks5透明代理服务器 \n 可配合外部工具实现流量操作", // 这里写协议
		UsageText: "lazy to write...",
		Version:   "0.1.0",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "DBG", Aliases: []string{"D"}, Destination: &Global.DBG, Value: false, Usage: "DBG MOD", Required: false},
			&cli.StringFlag{Name: "checkN", Aliases: []string{"C"}, Destination: &Global.PORT, Value: "8000", Usage: "监听端口", Required: false},
			//&cli.IntFlag{Name: "checkN", Aliases: []string{"C"}, Destination: &Global.CHECKN, Value: 3, Usage: "同一端口检测次数", Required: false},
		},
		HideHelpCommand: true,
		Action: func(c *cli.Context) error {
			err := do()
			if err != nil {

			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		//panic(err)
	}

	//fmt.Printf(os.Args[1])
}
func do() error {
	fmt.Printf("Socks5服务器开启\n")
	fmt.Printf("监听端口%v\n", Global.PORT)
	systemutils.SetCpuWithMax()
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	// Create SOCKS5 proxy on localhost port 8000
	if err := server.ListenAndServe("tcp", "127.0.0.1:8000"); err != nil {
		panic(err)
	}
	return nil

}
