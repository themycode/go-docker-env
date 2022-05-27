package src

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	platform string
)

var openCmds = map[string]string{
	"windows": "cmd /c start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "cli is demo for cobra",
}

var baikeCmd = &cobra.Command{
	Use:     "baike",
	Aliases: []string{"bk", "wk", "wiki"},
	Short:   "find things in baike site",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := findInBaiKe(args[0], platform)
		if err != nil {
			fmt.Println("err", err)
			os.Exit(1)
		}
	},
}

func init() {
	RootCommand.AddCommand(baikeCmd)
	baikeCmd.Flags().StringVarP(&platform, "platform", "p", "baidu", "platform to find things")
}

func findInBaiKe(keyword, platform string) error {
	var link string
	if platform == "baidu" || platform == "bd" {
		link = fmt.Sprintf("https://baike.baidu.com/item/%s", keyword)
	}
	if platform == "hudong" || platform == "baike" || platform == "hd" {
		link = fmt.Sprintf("http://www.baike.com/wiki/%s", keyword)
	}
	if platform == "wikipedia" || platform == "wiki" || platform == "wp" {
		link = fmt.Sprintf("https://zh.wikipedia.org/wiki/%s", keyword)
	}
	if link == "" {
		return fmt.Errorf("invalid platform")
	}
	return openLink(link)
}

func openLink(link string) error {
	goos := runtime.GOOS
	opencmd := "open"
	opencmd, ok := openCmds[goos]
	if !ok {
		return fmt.Errorf("can not open link in %s", goos)
	}
	if err := exec.Command(opencmd, link).Start(); err != nil {
		return err
	}
	return nil
}

func Execute() {
	if err := baikeCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
