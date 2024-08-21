package dao

import (
	"gin-ranking/config"
	"gin-ranking/pkg/logger"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	Db, err = gorm.Open("mysql", config.MysqlDSN)
	if err != nil {
		logger.Error(map[string]interface{}{"mysql connect error:": err.Error()})
	}
	if Db.Error != nil {
		logger.Error(map[string]interface{}{"mysql connect error:": Db.Error.Error()})
	}

	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	Db.DB().SetConnMaxLifetime(time.Hour)
}
