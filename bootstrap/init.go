package bootstrap

import (
	"Coattails/global"
	"Coattails/helper/windows/admin"
	"Coattails/service/buffApi"
	"Coattails/service/ws"
	"github.com/jinzhu/now"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/sys/windows"
	"os"
)

const (
	defaultTimeZone = "Asia/Shanghai"
)

func InitApp() error {
	//必须用管理员权限运行
	admin.MustRunWithAdmin()
	//
	initConsole()
	//
	//initConf()
	//initLog(&global.Conf.Log)
	initLib()
	initApi()
	//initGlobal()
	return nil
}

func initConsole() { //控制台的输出模式控制
	stdIn := windows.Handle(os.Stdin.Fd())
	var consoleMode uint32
	_ = windows.GetConsoleMode(stdIn, &consoleMode)
	consoleMode = consoleMode&^windows.ENABLE_QUICK_EDIT_MODE | windows.ENABLE_EXTENDED_FLAGS
	_ = windows.SetConsoleMode(stdIn, consoleMode)
}

//func initConf() {
//	//_ = godotenv.Load(".env")
//	//if bdk.IsFile(".env.local") {
//	//	_ = godotenv.Overload(".env.local")
//	//}
//	// confPath := "./config/config.json"
//	// err := configor.Load(global.Conf, confPath)
//	//*global.Conf = global.DefaultAppConf
//	//err := configor.Load(global.Conf)
//	//if err != nil {
//	//	panic(err)
//	//}
//	err = initClientConf()
//	if err != nil {
//		panic(err)
//	}
//}

//func initClientConf() (err error) {
//	dbPath := conf.SqliteDBPath
//	var db *gorm.DB
//	var dbLogger = gormLogger.Discard
//	//if global.IsDevMode() {
//	//	dbLogger = gormLogger.Default
//	//}
//	gormCfg := &gorm.Config{
//		Logger: dbLogger,
//	}
//	if !bdk.IsFile(dbPath) {
//		db, err = gorm.Open(sqlite.Open(dbPath), gormCfg)
//		bts, _ := json.Marshal(global.DefaultClientConf)
//		err = db.Exec(models.InitLocalClientSql, models.LocalClientConfKey, string(bts)).Error
//		if err != nil {
//			return
//		}
//		*global.ClientConf = global.DefaultClientConf
//	} else {
//		db, err = gorm.Open(sqlite.Open(dbPath), gormCfg)
//		confItem := &models.Config{}
//		err = db.Table("config").Where("k = ?", models.LocalClientConfKey).First(confItem).Error
//		if err != nil {
//			return
//		}
//		localClientConf := &conf.Client{}
//		err = json.Unmarshal([]byte(confItem.Val), localClientConf)
//		if err != nil || conf.ValidClientConf(localClientConf) != nil {
//			return errors.New("本地配置错误")
//		}
//		global.ClientConf = localClientConf
//	}
//	global.SqliteDB = db
//	return nil
//}

func initLib() {
	_ = os.Setenv("TZ", defaultTimeZone)
	now.WeekStartDay = time.Monday
	go func() {
		initUserInfo()
		//if global.Conf.Sentry.Enabled { //哨兵模式
		//	_ = initSentry(global.Conf.Sentry.Dsn)
		//}
	}()
	ws.Init() //初始化websocket连接
}

func initUserInfo() {
	resp, err := http.Get("https://api.ip.sb/ip") //获取本地的ip
	if err != nil {
		return
	}
	defer resp.Body.Close()
	bts, _ := io.ReadAll(resp.Body)
	global.SetUserInfo(global.UserInfo{
		IP: strings.Trim(string(bts), "\n"),
		// Mac:   windows.GetMac(),
		// CpuID: windows.GetCpuID(),
	})
}

func initApi() {
	buffApi.Init(global.Conf.BuffApi.Url, global.Conf.BuffApi.Timeout)
}

func initGlobal() {
	go initAutoReloadCalcConf()
}

func initAutoReloadCalcConf() {
	ticker := time.NewTicker(time.Minute)
	for {
		latestScoreConf, err := buffApi.GetCurrConf()
		if err == nil && latestScoreConf != nil && latestScoreConf.Enabled {
			global.SetScoreConf(*latestScoreConf)
		}
		<-ticker.C
	}
}
