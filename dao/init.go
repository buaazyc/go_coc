package dao

import (
	"database/sql"
	"fmt"

	"go_coc/config"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlProxy *sql.DB

type CurrentWar struct {
	Tag  string `db:"clan_tag"`
	Time string `db:"start_time"`
	Info string `db:"war_info"`
}

// ConnectDB 连接数据库
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

/*
建表命令
create table current_war
(clan_tag varchar(32) not null, start_time timestamp not null, info text not null)
engine=InnoDB DEFAULT charset=utf8mb4;
*/
