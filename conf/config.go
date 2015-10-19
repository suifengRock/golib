package conf

import (
	"encoding/json"
	"fmt"
	"golib/log"
	"io/ioutil"
)

type MysqlConf struct {
	ConnType string
	Host     string
	Port     string
	UserName string
	PassW    string
	DBName   string
	Debug    bool
}

type AccessTokenConf struct {
	Addr    string
	Prefix  string
	DB      int
	Expires int
}

type Config struct {
	Debug           bool
	ServerAddr      string
	SessionSecret   string
	Domain          string
	AccessToken     *AccessTokenConf
	UploadAuthToken *AccessTokenConf
	Mysql           *MysqlConf
}

func InitConfig(configFilePath string) *Config {
	Conf = new(Config)

	confPath := "conf/conf.json"
	if len(configFilePath) > 0 {
		confPath = configFilePath
	}

	cfgbuf, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file failed: ", confPath, err))
	}

	err = json.Unmarshal(cfgbuf, Conf)
	if err != nil {
		log.Fatal(fmt.Sprintf("Read config file failed: ", confPath, err))
	}
	log.Info("Using default conf: " + confPath)

	if Conf.Debug {
		Conf.Mysql.Debug = true
	}
	return Conf
}
