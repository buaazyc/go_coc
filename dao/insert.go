package dao

// InsertCurrentWar 插入数据到表current_war
func InsertCurrentWar(clan string, time string, war string) error {
	if clan == "" {
		return nil
	}
	sql := `insert into current_war (clan_tag, start_time, war_info) values (?, ?, ?)
	on duplicate key update war_info=?;`
	_, err := mysqlProxy.Exec(sql, clan, time, war, war)
	return err
}
