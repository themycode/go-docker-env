package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// urfave 框架使用
func HiGreet() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Action: func(c *cli.Context) error {
			name := "Nefertiti"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func readFileConfig() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Load configuration from `FILE`",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("err:", err)
	}
}
func goTool() {
	app := cli.NewApp()
	app.Name = "GoTool"
	app.Usage = "To save the world"
	app.Version = "1.0.0"

	// 预设变量
	var host string
	var debug bool

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "host",
			Value:       "127.0.0.1",
			Usage:       "server Address",
			Destination: &host,
		},
		cli.IntFlag{
			Name:  "port,p",
			Value: 8888,
			Usage: "Server port",
		},
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "debug mode",
			Destination: &debug,
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Printf("host = %v \n", host)
		fmt.Printf("host = %v \n", c.Int("port")) // 不使用变量接收，直接解析
		fmt.Printf("host = %v \n", debug)

	}
}

func main() {
	// HiGreet()
	// readFileConfig()
	goTool()
	fmt.Println("begin")

	for i := 0; i < 10; i++ {
		if i == 7 {
			goto Print
		}
		fmt.Println("i = ", i)

	}
Print:
	fmt.Println("end")

}
