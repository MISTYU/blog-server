package main

import (
	"blog_srvs/blog_srv/model"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// [username]:[password]@tcp(address)/dbname
	dsn := "root:root@tcp(120.77.200.233:3308)/czblog?charset=utf8mb4&parseTime=True&loc=Local"
	// 设置全局 logger, 这个 logger 在我们执行每个 sql 语句的时候会打印每一行 sql
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			Colorful:      true,        // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 生成的表是单数
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 同步表
	_ = db.AutoMigrate(&model.Blog{})
}
