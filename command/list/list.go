package list

import (
	"fmt"

	javaInfoHandle "github.com/EchoNoReturn/yojavaTools/datahandles"
	"github.com/mkideal/cli"
)

type List struct {
	cli.Helper
	Online bool `cli:"o,online" usage:"获取在线支持的Java版本"`
}

var ListCommand = &cli.Command{
	Name: "list",
	Desc: "列出所有Java版本  默认展示本地java版本列表",
	Argv: func() interface{} { return new(List) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*List)
		if argv.Online {
			// 发起请求获取在线支持的Java版本
			// TODO 采用新的策略，给出下载jdk的网址，或者直接帮助用户打开到默认浏览器，
			/*
				让用户自己手动下载tar包，并把包自行放在自己选择的目录下，然后使用add命令添加tar文件。
				注意这里不要让用户自己去解压，统一使用命令进行解压。
			*/
			fmt.Println("支持的在线java版本")
			return nil
		}
		// 列出本机中收录的java版本
		data, err := javaInfoHandle.GetLocalJava()
		if err != nil {
			fmt.Println("读取本地java版本数据失败:", err)
			return nil
		}
		javaInfoHandle.PrintJavaList(data)
		return nil
	},
}
