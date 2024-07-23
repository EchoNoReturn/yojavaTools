package manager

import (
	"github.com/mkideal/cli"
)

type UseManager struct {
	cli.Helper
}

var Use = &cli.Command{
	Name: "use",
	Desc: "use [javaVersion] 使用 Java 的某个版本",
	Argv: func() interface{} {
		return new(UseManager)
	},
	Fn: func(ctx *cli.Context) error {
		return nil
	},
}
