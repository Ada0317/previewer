package Test

import (
	"golang.org/x/sys/windows"
	"syscall"
	"testing"
)

func Test_windows_ShellExecute(t *testing.T) {
	cmd := ""
	verbPtr, err := syscall.UTF16PtrFromString(cmd) //操作类型
	if err != nil {
		t.Fatal(verbPtr)
	}
	open, err := syscall.UTF16PtrFromString("open") //操作类型
	if err != nil {
		t.Fatal(verbPtr)
	}
	calc, err := syscall.UTF16PtrFromString("calc.exe") //操作类型
	if err != nil {
		t.Fatal(verbPtr)
	}
	var showCmd int32 = 1
	err = windows.ShellExecute(0, open, calc, verbPtr, verbPtr, showCmd)
	if err != nil {
		t.Fatal(err)
	}

}

func Test_Stdin(t *testing.T) {

}
