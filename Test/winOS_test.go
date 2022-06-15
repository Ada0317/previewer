package Test

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

func Test_Executable(t *testing.T) {
	exe, err := os.Executable()
	t.Log(exe)
	t.Log(err)
	getwd, err := os.Getwd()

	t.Log(getwd)
	t.Log(err)
}

func Test_flag(t *testing.T) { //好像没法用test去测试
	showVersion := flag.Bool("v", false, "展示版本信息")
	flag.Parse()
	fmt.Printf("%+v", *showVersion)
}
