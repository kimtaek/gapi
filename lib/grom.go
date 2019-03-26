package lib

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var master *gorm.DB
var slave *gorm.DB

func InitDatabase() *gorm.DB {
	var err error
	if config.AppMode == "" {
		return master
	}
	dsn := config.DatabaseConfig.MysqlConfig.Username + ":"
	dsn += config.DatabaseConfig.MysqlConfig.Password + "@tcp("
	dsn += config.DatabaseConfig.MysqlConfig.Host + ":"
	dsn += config.DatabaseConfig.MysqlConfig.Port + ")/"
	dsn += config.DatabaseConfig.MysqlConfig.Database + "?charset=utf8&parseTime=True"

	master, err = gorm.Open("mysql", dsn)
	if err != nil {
		SendSlackMessage(Slack{
			Text: "ORM: " + err.Error(),
		})
	}
	err = master.DB().Ping()
	if err != nil {
		SendSlackMessage(Slack{
			Text: "ORM: " + err.Error(),
		})
	}
	if config.AppMode != "release" {
		master.LogMode(true)
	}
	master.DB().SetMaxIdleConns(config.DatabaseConfig.MysqlConfig.MaxIdleConnections)
	master.DB().SetMaxOpenConns(config.DatabaseConfig.MysqlConfig.MaxOpenConnections)
	return master
}

func Master() *gorm.DB {
	return master
}

func Slave() *gorm.DB {
	return master
}
