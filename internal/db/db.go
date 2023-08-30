package db

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/oskov/megabot/internal/config"
	"time"
)

var db *sqlx.DB

func GetDb() (*sqlx.DB, error) {
	if db != nil {
		return db, nil
	}
	conf, err := config.GetConfig()
	if err != nil {
		return nil, err
	}

	mysqlConfig := mysql.Config{
		User:                 conf.DbUser,
		Passwd:               conf.DbPass,
		Net:                  "tcp",
		Addr:                 conf.DbAddr,
		DBName:               conf.DbName,
		AllowNativePasswords: true,
		InterpolateParams:    true,
		ParseTime:            true,
	}
	conn, err := sqlx.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		return nil, err
	}
	conn.SetMaxIdleConns(conf.DbMaxIdleConnections)
	conn.SetMaxOpenConns(conf.DbMaxOpenConnections)
	conn.SetConnMaxIdleTime(time.Duration(conf.DbConnectionMaxIdleTimeInSeconds) * time.Second)
	conn.SetConnMaxLifetime(time.Duration(conf.DbConnectionMaxLifeTimeInSeconds) * time.Second)

	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	db = conn

	return conn, nil
}
