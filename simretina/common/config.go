package common

import "go.pfgit.cn/letsgo/xdev"

type XConfig struct {
	ConnStr  string `key:"common.db_conn" commit:"false"`
	LogLevel string `key:"common.log_level" default:"INFO" commit:"false"`

	// Port string `key:"common.port" default:"INFO"`

	//TODO:定义模块自己需要的配置，字段如何定义可查看xdev.ReadConfig说明
	HostPort string `key:"common.host_port" default:"http://172.16.1.101:30082" commit:"false"`
}

var Config *XConfig

func ReadConfig() error {
	Config = new(XConfig)
	return xdev.ReadConfig(APP_CONFIG_PATH, APP_NAME, Config)
}
