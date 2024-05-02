package main

import (
	"fmt"
	"os"

	list "github.com/EchoNoReturn/yojavaTools/command/list"
	manager "github.com/EchoNoReturn/yojavaTools/command/manager"
	value "github.com/EchoNoReturn/yojavaTools/common"
	"github.com/mkideal/cli"
)

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(list.ListCommand),
		cli.Tree(&manager.Add),
		cli.Tree(&manager.Remove),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var help = cli.HelpCommand("显示帮助信息")

// root command
type rootT struct {
	cli.Helper
	Name string `cli:"n,name" usage:"your name"`
}

var root = &cli.Command{
	Desc: "yojava version: " + value.Version,
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		ctx.String("Hello, root command, I am %s\n", argv.Name)
		return nil
	},
}
