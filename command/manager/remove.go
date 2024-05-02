package manager

import (
	"fmt"

	"github.com/mkideal/cli"

	javaInfoHandle "github.com/EchoNoReturn/yojavaTools/datahandles"
)

type RemoveOpts struct {
	cli.Helper
}

var Remove = cli.Command{
	Name: "rm",
	Desc: "rm [Name/Version] 以删除已添加的java版本,支持版本号模糊匹配(从版本号开头开始匹配)，支持同时删除多个",
	Argv: func() interface{} { return new(RemoveOpts) },
	Fn: func(ctx *cli.Context) error {
		rmList := ctx.Args()
		if len(rmList) == 0 {
			println("请输入要删除的java名称或版本号，使用\nyojava remove java名称或版本号1 java名称或版本号2...\n例如：yojava remove java-11 java-8\n查看帮助请使用yojava help remove")
			return nil
		}
		for i := 0; i < len(rmList); i++ {
			isOk, err := javaInfoHandle.RemoveJavaFromJson(rmList[i])
			if err != nil {
				fmt.Println("删除" + rmList[i] + "失败:", err)
				return err
			} else if isOk {
				fmt.Println("删除" + rmList[i] + "成功")
			}
		}
		return nil
	},
}
