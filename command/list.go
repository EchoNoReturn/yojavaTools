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
			fmt.Println("支持的在线java版本")
			return nil
		}
		// 列出本机中收录的java版本
		data, err := javaInfoHandle.GetLocalJava()
		if err!= nil {
			fmt.Println("读取本地java版本数据失败:", err)
			return nil
		}
		javaInfoHandle.PrintJavaList(data)
		return nil
	},
}



