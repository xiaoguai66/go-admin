package conf

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

func InitDB() (*gorm.DB, error) {
	logMode := logger.Info
	if !viper.GetBool("mode.develop") {
		logMode = logger.Error
	}
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dsn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, err
	}

	sqlDB, _ := db.DB()
	//如果 MaxOpenConns 大于 0 但小于新的 MaxIdleConns，则新的 MaxIdleConns 将减少以匹配 MaxOpenConns 限制
	//如果 MaxIdleConns 大于 0 并且新的 MaxOpenConns 小于 MaxIdleConns，则 MaxIdleConns 将减少以匹配新的 MaxOpenConns 限制
	sqlDB.SetMaxIdleConns(viper.GetInt("db.maxIdleConn")) //设置空闲连接池中的最大连接数
	sqlDB.SetMaxOpenConns(viper.GetInt("db.maxOpenConn")) //置到数据库的最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour)                   //设置连接可复用的最大时间

	return db, nil
}
