package Test

import (
	"os"
	"testing"
)

func Test_Executable(t *testing.T) {
	exe, err := os.Executable()

	t.Log(exe)
	t.Log(err)
}
