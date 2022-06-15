package Test

import (
	"Coattails/bootstrap"
	"Coattails/helper/windows/process"
	"errors"
	"log"
	"regexp"
	"strconv"
	"testing"
)

var (
	lolCommandlineReg     = regexp.MustCompile(`--remoting-auth-token=(.+?)" "--app-port=(\d+)"`)
	ErrLolProcessNotFound = errors.New("未找到lol进程")
)

func Test_getsummoner(t *testing.T) {
	err := bootstrap.InitApp()
	if err != nil {
		log.Fatalln(err)
		return
	}
	cmdline, err := process.GetProcessCommand("LeagueClientUx.exe")
	if err != nil {
		log.Fatalln(ErrLolProcessNotFound)
		return
	}
	btsChunk := lolCommandlineReg.FindSubmatch([]byte(cmdline))
	if len(btsChunk) < 3 {
		log.Fatalln(ErrLolProcessNotFound)
	}

	token := string(btsChunk[1])
	port, err := strconv.Atoi(string(btsChunk[2]))

	t.Log(token + "\n")
	t.Log(port)
}
