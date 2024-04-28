package manager

/*
添加本地的java文件到yojava管理
*/

import (
	"errors"
	"os/exec"
	"strings"

	javaInfoHandle "github.com/EchoNoReturn/yojavaTools/datahandles"
	"github.com/mkideal/cli"
)

type AddOptions struct {
	cli.Helper
}

var Add = cli.Command{
	Name:  "add",
	Desc:  "添加 java 到 yojava",
	Argv:  func() interface{} { return new(AddOptions) },
	Fn: func(ctx *cli.Context) error {
		pathList := ctx.Args()
		if len(pathList) == 0 {
			println("未能获取到Java路径\n使用方式：yojava add java路径1 java路径2 ...\n例如：yojava add /usr/lib/jvm/java-11-openjdk-amd64/ /usr/lib/jvm/java-8-openjdk-amd64/")
		}
		for i := 0; i < len(pathList); i++ {
			version, err := getJdkVersion(pathList[i])
			if err != nil {
				return err
			}
			javaInfo := createJavaInfo(pathList[i], version)
			javaInfoHandle.AddJava2Json(javaInfo)
		}
		return nil
	},
}

/*
通过路径获取jdk版本信息
*/
func getJdkVersion(path string) (string, error) {
	// 先判断路径是否存在
	if _, err := exec.LookPath(path+"/bin/java"); err!= nil {
		err := errors.New("路径 " + path + " 下不存在Java文件")
		return "", err
	}
	// 获取版本信息
	cmd := exec.Command(path+"/bin/java", "-version")
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	if err != nil {
		err := errors.New("通过 " + path + " 获取jdk信息失败")
		return "", err
	} else {
		versionInfo := strings.TrimSpace(string(output))
		version := strings.ReplaceAll(strings.Split(versionInfo, " ")[2], "\"", "")
		return version, nil
	}
}

/* 创建JavaData对象内容 */
func createJavaInfo(path string, version string) javaInfoHandle.JavaData {
	return javaInfoHandle.JavaData{
		Name: "java" + version,
		Version: version,
		Path: path,
	}
}
