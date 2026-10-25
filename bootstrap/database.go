package bootstrap

import (
	"errors"
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/database"
	"gohub/pkg/logger"
	"net"
	"net/url"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupDB 初始化数据库和 ORM
func SetupDB() {

	var dbConfig gorm.Dialector
	switch config.Get("database.connection") {
	case "mysql":
		// 构建 DSN 信息
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=True&multiStatements=true&loc=Local",
			config.Get("database.mysql.username"),
			config.Get("database.mysql.password"),
			config.Get("database.mysql.host"),
			config.Get("database.mysql.port"),
			config.Get("database.mysql.database"),
			config.Get("database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	case "sqlite":
		// 初始化 sqlite
		database := config.Get("database.sqlite.database")
		dbConfig = sqlite.Open(database)
	case "postgres":
		// 使用 URL 编码，确保密码等包含特殊字符时仍能正确解析。
		postgresURL := &url.URL{
			Scheme: "postgres",
			User:   url.UserPassword(config.Get("database.postgres.username"), config.Get("database.postgres.password")),
			Host:   net.JoinHostPort(config.Get("database.postgres.host"), config.Get("database.postgres.port")),
			Path:   "/" + config.Get("database.postgres.database"),
		}
		query := url.Values{}
		query.Set("sslmode", config.Get("database.postgres.sslmode"))
		query.Set("TimeZone", config.Get("database.postgres.timezone"))
		postgresURL.RawQuery = query.Encode()
		dbConfig = postgres.Open(postgresURL.String())
	default:
		panic(errors.New("database connection not supported"))
	}

	// 连接数据库，并设置 GORM 的日志模式
	database.Connect(dbConfig, logger.NewGormLogger())

	// 设置最大连接数
	database.SQLDB.SetMaxOpenConns(config.GetInt("database.max_open_connections"))
	// 设置最大空闲连接数
	database.SQLDB.SetMaxIdleConns(config.GetInt("database.max_idle_connections"))
	// 设置每个链接的过期时间
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.max_life_seconds")) * time.Second)
}
