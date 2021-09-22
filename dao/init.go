package dao

import (
	"database/sql"
	"fmt"

	"go_coc/config"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlProxy *sql.DB

func ConnectDB() (err error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		config.Conf.MysqlUser,
		config.Conf.MysqlPassword,
		config.Conf.MysqlHost,
		config.Conf.MysqlPort,
		config.Conf.MysqlDBName,
	)
	mysqlProxy, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err = mysqlProxy.Ping(); err != nil {
		return err
	}
	return nil
}
