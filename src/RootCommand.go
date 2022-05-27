package src

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// cobra  框架使用
var RootCommand = &cobra.Command{
	Use:   "cli",
	Short: "cli is demo for cobra",
}

var BuildSubCommand = &cobra.Command{
	Use: "build",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var PsCommand = &cobra.Command{
	Use:   "ps",
	Short: "这是一个操作ps的命令",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(" args is:", len(args))
		fmt.Println("/bin/bash", "-c", "ps -ef| grep "+args[0])
		if len(args) != 0 {
			cmd := exec.Command("/bin/bash", "-c", "ps -ef| grep "+args[0])
			//创建获取命令输出管道
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
				return
			}

			//执行命令
			if err := cmd.Start(); err != nil {
				fmt.Println("Error:The command is err,", err)
				return
			}

			//读取所有输出
			bytes, err := ioutil.ReadAll(stdout)
			if err != nil {
				fmt.Println("ReadAll Stdout:", err.Error())
				return
			}

			if err := cmd.Wait(); err != nil {
				fmt.Println("wait:", err.Error())
				return
			}
			fmt.Printf("stdout:\n\n %s", bytes)

		}
	},
}

var GitCmd = &cobra.Command{
	Use:   "git",
	Short: "Git is distributed version control system",
	Long: `Git is a free and open source distributed version control system
	designed to handle everything from small to very large projects 
	with speed and efficiency.`,
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unrecognized command"))
	},
}

func Error(cmd *cobra.Command, args []string, err error) {
	fmt.Fprintf(os.Stderr, "execute %s args:%v error:%v\n", cmd.Name(), args, err)
	os.Exit(1)
}

var Cmd = &cobra.Command{
	Use:   "hello",
	Short: "hello test",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return fmt.Errorf("invalid color specified: %s", args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello , World!")
	},
}

func init() {
	RootCommand.AddCommand(BuildSubCommand)
	RootCommand.AddCommand(PsCommand)
	RootCommand.AddCommand(Cmd)
	RootCommand.AddCommand(GitCmd)
	RootCommand.Flags().BoolP("update", "u", false, "")
}
