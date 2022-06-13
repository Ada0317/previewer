package admin

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os"
	"strings"
	"syscall"
)

func MustRunWithAdmin() {
	if !IsAdmin() { //查看是否是管理员权限
		RunMeElevated() //如果不是就开启权限
	}
}

func IsAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}

func RunMeElevated() { //重新用管理员身份打开程序
	verb := "runas"                        //runas命令将用户的权限提升为管理员
	exe, _ := os.Executable()              //获取当前执行路径
	cwd, _ := os.Getwd()                   //返回一个对应当前工作目录的根路径
	args := strings.Join(os.Args[1:], " ") //把字符串数组的元素用step作为间隔链接成一个长的string

	verbPtr, _ := syscall.UTF16PtrFromString(verb) //操作类型
	exePtr, _ := syscall.UTF16PtrFromString(exe)   //文件路径
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)   //文件打开的方式
	argPtr, _ := syscall.UTF16PtrFromString(args)  //参数 -v之类的

	var showCmd int32 = 1                                                    // SW_NORMAL
	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd) //运行一个外部程序
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(-2)
}
