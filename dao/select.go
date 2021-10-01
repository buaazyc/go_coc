package dao

import "log"

// QueryAllClanTags 获取数据库current_war表中所有等clanTags
func QueryAllClanTags() ([]string, error) {
	// 调用sql语句
	sql := `SELECT clan_tag FROM current_war;`
	rows, err := mysqlProxy.Query(sql)
	if err != nil {
		return nil, err
	}
	clanTags := make(map[string]bool)
	// 结果逐行读取加入集合中
	for rows.Next() {
		var clanTag string
		if err := rows.Scan(&clanTag); err != nil {
			log.Printf("rows.Scan err: %v", err)
			continue
		}
		if !clanTags[clanTag] {
			clanTags[clanTag] = true
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	// 集合转切片
	var res []string
	for clanTag := range clanTags {
		res = append(res, clanTag)
	}
	return res, nil
}

// QueryLastCurrentWar 从数据库中获取最新的部落战
func QueryLastCurrentWar(clan string) (string, error) {
	sql := `SELECT start_time, war_info FROM current_war WHERE clan_tag = ?;`
	rows, err := mysqlProxy.Query(sql, clan)
	if err != nil {
		return "", err
	}
	var (
		lastTime    string
		lastWarInfo string
	)
	for rows.Next() {
		var (
			startTime string
			warInfo   string
		)
		if err := rows.Scan(&startTime, &warInfo); err != nil {
			log.Printf("rows.Scan err: %v", err)
			continue
		}
		if lastTime == "" || lastTime < startTime {
			lastTime = startTime
			lastWarInfo = warInfo
		}
	}
	return lastWarInfo, nil
}
