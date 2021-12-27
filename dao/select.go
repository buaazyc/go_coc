package dao

import (
	"fmt"
	"log"

	"go_coc/time"
)

// QueryActiveClanTags 获取数据库current_war表中活跃的clanTags
func QueryActiveClanTags() ([]string, error) {
	// 调用sql语句：查询最近一个月有战绩的部落标签
	sql := `SELECT distinct clan_tag FROM current_war WHERE start_time like ?;`
	rows, err := mysqlProxy.Query(sql, fmt.Sprintf("%v%%", time.CurMonth()))
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

// QueryAllWars 从数据库中获取特定部落的所有战绩
func QueryAllWars(clan string) ([]*CurrentWar, error) {
	sql := `SELECT clan_tag, start_time, war_info FROM current_war WHERE clan_tag = ?;`
	rows, err := mysqlProxy.Query(sql, clan)
	if err != nil {
		return nil, err
	}
	var res []*CurrentWar
	for rows.Next() {
		row := &CurrentWar{}
		if err := rows.Scan(&row.Tag, &row.Time, &row.Info); err != nil {
			log.Printf("rows.Scan err: %v", err)
			continue
		}
		res = append(res, row)
	}
	return res, nil
}

// QueryNewestWar 从数据库中获取特定部落最新的战绩
func QueryNewestWar(clan string) (*CurrentWar, error) {
	sql := `SELECT clan_tag, start_time, war_info FROM current_war 
	WHERE clan_tag = ? ORDER BY create_time DESC LIMIT 1;`
	rows, err := mysqlProxy.Query(sql, clan)
	if err != nil {
		return nil, err
	}
	var res []*CurrentWar
	for rows.Next() {
		row := &CurrentWar{}
		if err := rows.Scan(&row.Tag, &row.Time, &row.Info); err != nil {
			log.Printf("rows.Scan err: %v", err)
			continue
		}
		res = append(res, row)
	}
	if len(res) == 0 {
		return nil, nil
	}
	return res[0], nil
}
