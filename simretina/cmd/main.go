package main

import (
	"os"
	"simretina/common"

	"simretina/core/router"

	_ "github.com/joho/godotenv/autoload"
	"go.pfgit.cn/letsgo/xdev"
)

func main() {
	var Log = xdev.XNewLoggerDefault("./log/" + common.APP_NAME + ".log")
	Log.Info(common.APP_NAME + "/" + common.APP_VERSION)

	if err := common.ReadConfig(); err != nil {
		os.Exit(1)
	}
	router.InitRouter(common.Config.HostPort)
}
