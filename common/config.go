package common

import (
	"encoding/json"
	"github.com/taoshihan1991/imaptool/tools"
	"io/ioutil"
	"os"
)

type Mysql struct {
	Server   string
	Port     string
	Database string
	Username string
	Password string
}

func GetMysqlConf() *Mysql {
	var mysql = &Mysql{}
	isExist, _ := tools.IsFileExist(MysqlConf)
	if !isExist {
		return GetMysqlConfFromEnvironment()
	}
	info, err := ioutil.ReadFile(MysqlConf)
	if err != nil {
		return mysql
	}
	err = json.Unmarshal(info, mysql)
	return mysql
}

func GetMysqlConfFromEnvironment() *Mysql {
	var mysql = &Mysql{}
	mysql.Server = os.Getenv("MYSQL_SERVER")
	mysql.Port = os.Getenv("MYSQL_PORT")
	mysql.Username = os.Getenv("MYSQL_USERNAME")
	mysql.Password = os.Getenv("MYSQL_PASSWORD")
	mysql.Database = os.Getenv("MYSQL_DATABASE")
	return mysql
}