package ostools

import (
	"errors"
	"fmt"
	"os"
	"runtime"
)

/*
获取系统的类型
*/
func getOsType() string {
	return runtime.GOROOT()
}

func UseOsAdept(winfn func(), linuxfn func(), macfn func()) {
	err := errors.New("not support")
	switch os.Getegid() {
	case 0:
		winfn()
	case 1:
		linuxfn()
	case 2:
		macfn()
	default:
		fmt.Println(err)
	}
}