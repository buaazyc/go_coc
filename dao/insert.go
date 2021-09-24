package dao

// InsertCurrentWar 表current_war插入数据
func InsertCurrentWar(clan string, time string, war string) error {
	SQL := `insert into current_war (clan_tag, start_time, war_info) values (?, ?, ?)
	on duplicate key update war_info=?;`
	_, err := mysqlProxy.Exec(SQL, clan, time, war, war)
	return err
}
